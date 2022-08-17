db:
	docker-compose -f docker-compose.dev.yaml up -d db adminer

api:
	go run ./cmd/api/main.go


