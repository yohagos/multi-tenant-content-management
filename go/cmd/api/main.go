package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/yohagos/multi-content-management/internal/adapter/handler"
	"github.com/yohagos/multi-content-management/internal/adapter/messaging"
	"github.com/yohagos/multi-content-management/internal/adapter/repository"
	"github.com/yohagos/multi-content-management/internal/config"
	"github.com/yohagos/multi-content-management/internal/core/service"
	"github.com/yohagos/multi-content-management/pkg/jwt"
	"github.com/yohagos/multi-content-management/pkg/logger"
	"github.com/yohagos/multi-content-management/pkg/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	log := logger.New(cfg.LogLevel)
	defer log.Sync()

	log.Info("Started Server....")
	log.Info("Loaded Configs, started logger and initializing Database....")

	db := initDB(cfg, log)
	defer db.Close()

	log.Info("Postgres => connected!")

	if err := repository.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations", zap.Error(err))
	}
	log.Info("Migrations completed successfully!")

	if err := repository.SeedAdminUser(db, cfg.AdminEmail, cfg.AdminPassword); err != nil {
		log.Error("Failed to seed admin user", zap.Error(err))
	} else {
		log.Info("Admin user seeded successfully")
	}

	redis := initRedis(cfg, log)
	defer redis.Close()

	log.Info("Redis => connected!")

	rabbit := initRabbitMQ(cfg, log)
	defer rabbit.Close()

	log.Info("RabbitMQ => connected!")

	jwtConfig := jwt.Config{
		Secret: cfg.JWTSecret,
		Expiry: cfg.JWTExpiry,
	}

	tenantRepo := repository.NewTenantRepository(db)
	contentRepo := repository.NewContentRepository(db)
	userRepo := repository.NewUserRepository(db, *log)
	tokenRepo := repository.NewTokenRepository(db)

	if err := tenantRepo.UpdateTenantMetrics(context.Background()); err != nil {
		log.Error("Failed to update tenant metrics")
	}

	if err := contentRepo.UpdateContentMetrics(context.Background()); err != nil {
		log.Error("Failed to update content metrics")
	}

	if err := userRepo.UpdateUserMetrics(context.Background()); err != nil {
		log.Error("Failed to update user metrics")
	}

	cacheRepo := repository.NewCacheRepository(redis)
	publisher := messaging.NewEventPublisher(rabbit)

	tenantService := service.NewTenantService(tenantRepo, cacheRepo, publisher)
	contentService := service.NewContentService(contentRepo, cacheRepo, publisher)

	authService := service.NewAuthService(userRepo, tokenRepo, &jwtConfig)

	tenantHandler := handler.NewTenantHandler(tenantService)
	contentHandler := handler.NewContentHandler(contentService)
	authHandler := handler.NewAuthHandler(authService)

	router := setupRouter(tenantHandler, contentHandler, authHandler, authService, log)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("Server starting", zap.String("Port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced shutdown", zap.Error(err))
	}

	log.Info("Server exited")
}

func initDB(cfg *config.Config, log *logger.Logger) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	log.Info("Database connected successfully.")

	return db
}

func initRedis(cfg *config.Config, log *logger.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	log.Info("Redis connected successfully.")

	return client
}

func initRabbitMQ(cfg *config.Config, log *logger.Logger) *amqp.Connection {
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort,
		),
	)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}

	log.Info("RabbitMQ connected successfully.")
	return conn
}

func setupRouter(
	tenantHandler *handler.TenantHandler,
	contentHandler *handler.ContentHandler,
	authHandler *handler.AuthHandler,
	authService *service.AuthService,
	log *logger.Logger,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.Recovery(log))
	router.Use(middleware.CORS())
	router.Use(middleware.Logging(log))
	router.Use(middleware.TenantContext())
	router.Use(middleware.RateLimiter())
	router.Use(middleware.Metrics())

	router.GET("/metrics", handler.MetricsHandler())

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", middleware.Auth(authService), authHandler.Logout)
	}

	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/api/v1")
	api.Use(middleware.Auth(authService))
	{
		tenants := api.Group("/tenants")
		tenants.Use(middleware.RequiredRole("admin", "tenants"))
		{
			tenants.POST("", tenantHandler.Create)
			tenants.GET("", tenantHandler.List)
			tenants.GET("/:id", tenantHandler.GetByID)
			tenants.PUT("/:id", tenantHandler.Update)
			tenants.DELETE("/:id", tenantHandler.Delete)
		}

		content := api.Group("/content")
		{
			content.POST("", contentHandler.Create)
			content.GET("", contentHandler.List)
			content.GET("/:id", contentHandler.GetByID)
			content.PUT("/:id", contentHandler.Update)
			content.DELETE("/:id", contentHandler.Delete)
			content.POST("/:id/publish", contentHandler.Publish)
		}
	}

	return router
}
