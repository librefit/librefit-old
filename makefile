build:
	go mod download
	go build --tags "json1" -o bin/librefit ./main.go

run_backend:
	go mod download
	go build --tags "json1" -o bin/librefit ./main.go
	./bin/librefit start

run_frontend:
	cd frontend/ ; yarn dev

test:
	go test ./...

fmt:
	go fmt ./...

