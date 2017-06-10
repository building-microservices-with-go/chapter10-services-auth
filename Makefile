unit:
	go test -v --race $(shell go list ./... | grep -v /vendor/)

build_auth:
	CGO_ENABLED=0 GOOS=linux go build -o ./services/auth/auth ./services/auth

build_docker: build_auth
	docker build -t building-microservices-with-go/auth .

run:
	go run main.go

test: unit
