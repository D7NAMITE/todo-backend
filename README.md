# Todo Backend
This service is responsible for **todo**-related backend operations 
for example - getting all **todo** for the user, changing **todo** status, and more.

## Prerequisites
* Go 1.18+
* PostgreSQL
* Goose CLI (go install github.com/pressly/goose/v3/cmd/goose@latest)
* sqlc CLI (go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest)
* Air (for hot reloading) (go install github.com/cosmtrek/air@latest)

## Setup Guide
1. Navigate to `todo-service`
```bash
cd todo-service
```
2. Install dependencies:
```bash
go mod tidy
```
3. Configure environment variables: \
You can change file name `.env.sample` to `.env` and edit the following:
   * `DATABASE_URL`: The URL of the database
   * `GOOSE_DRIVER`: Database provider (the `.env.sample` setted as `postgres`)
   * `GOOSE_DBSTRING`: Database connection string for migrations. In this case we can use the same URL with `DATABASE_URL`

4. Database migrations:
```bash
# Apply migrations
goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} "${GOOSE_DBSTRING}" up

# Check status
goose -dir ${GOOSE_MIGRATION_DIR} ${GOOSE_DRIVER} "${GOOSE_DBSTRING}" status
```

5. Generate database code:
```bash
sqlc generate 
```

6. Run the application:
```bash
# Standard run
go run main.go

# Development with hot reload
air 
```
