# This is the test and practice RESTful API

In this version, gateways includes database information.  
And this breaks DIP (Dependency Inversion Principal)

Not a perfect Clean Architecture.

## Getting Started

### for Mac (or Windows installed make)

To Start DB

```terminal
make up
```

To Start API Server

```terminal
make start
```

### for Windows

To start DB

```terminal
docker-compose up
```

To start api server

```terminal
go run .
```

## API Spec

You can see details in `/infrastructures/rounter.go` and `/adapters/gateways/user_gateways.go`

| End point                            | Method |
| ------------------------------------ | ------ |
| ~/users/all                          | GET    |
| ~/users/range?start=[id]&end=[id]    | GET    |
| ~/users?ids[0]=[id]&ids[1]=[id]&.... | GET    |
| ~/user/:id                           | GET    |

- ~/users/all  
  To get all users.  
- ~/users/range?start=[id]&end=[id]  
  To get users in a certain ids range.  
- ~/users?ids[0]=[id]&ids[1]=[id]&....  
  To get users according to ids  
  Ex) ~/users?ids[0]=1&ids[1]=3&ids[2]=8
- ~/user/:id
  To get a user  