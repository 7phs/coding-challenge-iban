IMAGE = github.com/7phs/coding-challenge-iban
VERSION = latest

build: export GO111MODULE=on
build:
	go mod vendor
	go build -o iban-validator -ldflags "-X github.com/7phs/coding-challenge-iban/cmd.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/7phs/coding-challenge-iban/cmd.GitHash=`git rev-parse --short HEAD`"

testing:
	LOG_LEVEL=error go test ./...

run: export GO111MODULE=on
run:
	go mod vendor
	LOG_LEVEL=info ADDR=:8080 STAGE=prod go run main.go run

image:
	docker build -t $(IMAGE):$(VERSION)  .

push:
	docker push $(IMAGE):$(VERSION)

image-run:
	docker run --rm -it -p 8080:8080 $(IMAGE):$(VERSION)

all: build
