# ECHO BLOG

## ER Diagram
![ER Diagram](docs/er-diagram.png)

## Useful commands

### Create Database and User
```sql
CREATE DATABASE echoblog;

CREATE USER echoblog WITH SUPERUSER PASSWORD 'secret';

GRANT ALL PRIVILEGES ON DATABASE echoblog TO echoblog;
```

### Create migration files
```bash
migrate create -ext sql -dir db/migrations -seq init_schema
```

### Set PostgreSQL URL
```bash
export POSTGRESQL_URL='postgres://echoblog:secret@localhost:5432/echoblog?sslmode=disable'
```

### Execute Migrations
```bash
migrate -database ${POSTGRESQL_URL} -path db/migrations up
migrate -database ${POSTGRESQL_URL} -path db/migrations down
```