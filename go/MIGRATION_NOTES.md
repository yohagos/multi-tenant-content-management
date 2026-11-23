# Multi Tenant Content Management

## Installing migration tool

Running the following command to get and install:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Create migration files

Move to the migration folder in which the migration files have to be stored and run the command:

```bash
migrate create -ext sql -dir ./migrations -seq create_tenants_table
```

| Parameter | Description |
| - | - |
| "-ext sql" | Generate a SQL file |
| "-dir ./migrations" | Path to save file in directory |
| "-seq" | Generate files in sequence |

## Execute migration through console

```bash
migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/multitenant?sslmode=disable" up
```

### More Commands

| Command | Description |
| - | - |
| "down" | Goes down one Migration |
| "version" | Display currently used Version |

