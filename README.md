# Greenlight

> Retrieve and manage movies through a JSON API

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Setup

### 1. Create a Docker container for PostgreSQL
```bash
docker run -d --name greenlight-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
```

### 2. Connect to the PostgreSQL container
```bash
docker exec -it greenlight-postgres psql -U postgres
```

### 3. Create a migration using golang-migrate
```bash
migrate create -seq -ext=.sql -dir=./migrations name_of_the_migration_here
```

### 4. Execute the migration
```bash
migrate -path=./migrations -database="postgres://greenlight:&P455word\$@localhost/greenlight?sslmode=disable" up
```
