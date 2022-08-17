### Toimi test task



#### Run project locally 


1. Copy environment file
```
    cp .env.example .env
```

2. Install dependencies 
```
    go mod install
```
3. Run migrations (used go-migrate: [Go migrate](github.com/golang-migrate/migrate))
```
    migrate -database "postgres://..." -path internal/migrations
```
4. Run code 
```
   make api   (go run cmd/api/main.go) Makefile
```


##Or



1. Copy environment file
```
    cp .env.example .env
```


2. Run docker-compose
```
    docker-compose up -d 
```

3. Dont forget about migration
