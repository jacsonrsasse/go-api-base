# GO API BASE

## How to work with it

You'll need a few softwares do work with this project:

### SQLC

Generates fully type-safe idiomatic code from SQL.

[Link](https://sqlc.dev/)

Write your SQL's insde the `queries.sql` file and run `sqlc generate` to get the files.

### Go Goose

Migration handler

[Link](https://github.com/pressly/goose)

```bash
# to create a sequencial migration
goose -s create add_some_column sql
```
