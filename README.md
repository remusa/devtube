# README

```bash
export DB_URL=postgres://postgres:postgres@localhost:5432/devtube?sslmode=disable

goose -dir db/schema sqlite database.db up

goose -dir db/schema sqlite database.db down
```
