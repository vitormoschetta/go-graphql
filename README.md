# Go GraphQL

https://gqlgen.com/

## Getting Started

```bash
go mod init github.com/username/project
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

Run:

```bash
go run server.go
```


## Alterar schema.graphql

Após alterar o arquivo `schema.graphql`, é necessário gerar os arquivos go com o comando:

```bash
go run github.com/99designs/gqlgen generate
```

## Criar tabelas no banco de dados

```bash
cd internal/database
sqlite3 database.db
```

```sql
create table categories (id string, name string, description string);
create table products (id string, name string, description string, price float, category_id string);
```