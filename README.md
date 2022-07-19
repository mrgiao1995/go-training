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

**Proto gen**

```sh
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```

## Testing

**1.Running the grpc service**

***Customer***
```sh
go run grpc/customer/main.go
```

***Flight***
```sh
go run grpc/flight/main.go
```

***Booking***
```sh
go run grpc/booking/main.go
```

**2.Running the api**

***Customer***
```sh
go run clients/rest/customer/main.go
```
