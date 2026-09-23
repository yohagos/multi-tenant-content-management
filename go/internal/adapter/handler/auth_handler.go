package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/internal/core/domain"
	"github.com/yohagos/multi-content-management/internal/core/service"
	"github.com/yohagos/multi-content-management/pkg/metrics"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		metrics.AuthRegisterTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		if err == service.ErrUserExists {
			metrics.AuthRegisterTotal.WithLabelValues("failure").Inc()
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}
		metrics.AuthRegisterTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.AuthRegisterTotal.WithLabelValues("success").Inc()
	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		metrics.AuthLoginTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	response, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			metrics.AuthLoginTotal.WithLabelValues("failure").Inc()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		metrics.AuthLoginTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.AuthLoginTotal.WithLabelValues("success").Inc()
	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req domain.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		metrics.AuthRefreshTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		if err == service.ErrInvalidToken || err == service.ErrTokenExpired {
			metrics.AuthRefreshTotal.WithLabelValues("failure").Inc()
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		metrics.AuthRefreshTotal.WithLabelValues("failure").Inc()
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	metrics.AuthRefreshTotal.WithLabelValues("success").Inc()
	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization header required / missing"})
		return
	}

	if err := h.authService.Logout(c.Request.Context(), token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
