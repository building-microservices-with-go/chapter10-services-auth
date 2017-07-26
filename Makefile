unit:
	go test -v --race $(shell go list ./... | grep -v /vendor/)

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o auth .

build_docker:
	docker build -t buildingmicroserviceswithgo/auth:latest .

run:
	go run main.go

staticcheck:
	staticcheck $(shell go list ./... | grep -v /vendor/)

test: unit
