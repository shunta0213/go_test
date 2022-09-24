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

## To add API End Point

1. Make a go structure in `/entities`  
2. Make port and interactor in `/usecases`  
3. Make controllers, gateways and presenters in `/adaptors`  
4. Make API endpoint in `/infrastructures`

## To add api tables and columns

1. Run  

```terminal
migrate create -ext sql -dir database/migrations -seq [sql file name]
```

2. Edit `database/[version]-migrations/[sql_file_name].up.sql` and `~.down.sql`

3. Run  

`${POSTGRESQL_URL} = 'postgres://user:pass@url:port/tablename'`  

```terminal
migrate -path database/migrations -database ${POSTGRESQL_URL} up 
```

When the migration is failed, your schema is recoded as "dirty".
Anyway, fix your migration and Run

To change your schema clear.

```terminal
migrate -path database/migrations -database ${POSTGRESQL_URL} force 1
```

And downgrade your tables

```terminal
migrate -path database/migrations -database ${POSTGRESQL_URL} down 1
```

And upgrade 1

```terminal
migrate -path database/migrations -database ${POSTGRESQL_URL} up 1
```

4.Run
To get a go structure from database schema.  
See `/models/~`

```terminal
sqlboiler psql
```