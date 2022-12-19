# Golang Migrate Demo

## Useful commands

### Create Database and User
```
CREATE DATABASE golang_migrate_demo;

CREATE USER golang_migrate_demo WITH SUPERUSER PASSWORD 'secret';

GRANT ALL PRIVILEGES ON DATABASE golang_migrate_demo TO golang_migrate_demo;
```

### Create migration files
```bash
migrate create -ext sql -dir db/migrations -seq init_schema
```

### Set PostgreSQL URL
```bash
export POSTGRESQL_URL='postgres://golang_migrate_demo:secret@localhost:5432/golang_migrate_demo?sslmode=disable'
```

### Execute Migrations
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations up
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```
