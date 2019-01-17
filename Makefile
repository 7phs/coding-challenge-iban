IMAGE = github.com/7phs/coding-challenge-iban
VERSION = latest

build: export GO111MODULE=on
build:
	go mod vendor
	go build -o iban-validator -ldflags "-X github.com/7phs/coding-challenge-iban/cmd.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/7phs/coding-challenge-iban/cmd.GitHash=`git rev-parse --short HEAD`"

testing:
	LOG_LEVEL=error DB_PATH=$(shell pwd)/data/countries-iban.yaml go test ./...

testing-short:
	LOG_LEVEL=error DB_PATH=$(shell pwd)/data/countries-iban.yaml go test -short ./...

image:
	docker build -t $(IMAGE):$(VERSION)  .

push:
	docker push $(IMAGE):$(VERSION)

run:
	docker run --rm -it -p 8080:8080 $(IMAGE):$(VERSION)

all: build
