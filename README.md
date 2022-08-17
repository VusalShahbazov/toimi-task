### Toimi test task
Announcements crud example



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


## Or



1. Copy environment file
```
    cp .env.example .env
```


2. Run docker-compose
```
    docker-compose up -d 
```

3. Dont forget about migration






Links to repos what i used
1. [Go clean arch](https://github.com/bxcodec/go-clean-arch)
2. [Folder struct](https://github.com/golang-standards/project-layout)
3. [Gin framework](https://github.com/gin-gonic/gin)
4. [Go pg orm](https://github.com/go-pg/pg)