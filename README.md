# grpc with go pratice

Pratice grpc microservices with go lang

## Setup env

**Go**

```sh
brew install go
```

**Docker**

```sh
brew install docker --cask
```

**Goose**

```sh
brew install goose
```

## Development

**Starting database**

```sh
docker-compose up
```

**Migration**

```sh
goose -dir database/migrations/${db_name} postgres ${DB_CONNECTION_STRING} up -dir "database/migrations/${db_name}"
```

**Proto gen**

```sh
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```
