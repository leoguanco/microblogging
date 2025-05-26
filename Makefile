.PHONY: build run test clean docker-build docker-run docker-compose-up docker-compose-down update-deps

build:
	go build -o microblog ./cmd/server

run: build
	./microblog

test:
	go test -v ./...

clean:
	rm -f microblog

docker-build:
	docker build -t microblog:latest .

docker-run: docker-build
	docker run -p 8080:8080 microblog:latest

docker-compose-up:
	docker compose up -d
	@echo "Services are starting up..."
	@echo "Grafana UI will be available at http://localhost:3000 (admin/admin)"
	@echo "Run 'make tracing-dashboard' to get the tracing dashboard URL"

docker-compose-down:
	docker compose down

update-deps:
	go mod tidy
	go mod vendor