IMAGE = github.com/7phs/coding-challenge-iban
VERSION = latest

build: export GO111MODULE=on
build:

	go build -o iban-validator -ldflags "-X github.com/7phs/coding-challenge-iban/cmd.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/7phs/coding-challenge-iban/cmd.GitHash=`${GIT_HASH}`"

testing:
	LOG_LEVEL=error DATA_PATH=$(shell pwd)/data go test ./...

testing-short:
	LOG_LEVEL=error DATA_PATH=$(shell pwd)/data go test -short ./...

image:
	docker build -t $(IMAGE):$(VERSION) --build-arg SSH_PRIVATE_KEY="$$(cat ~/.ssh/id_rsa)" --build-arg GIT_HASH="$$(git rev-parse --short HEAD)" .

push:
	docker push $(IMAGE):$(VERSION)

run:
	docker run --rm -it -p 8080:8080 $(IMAGE):$(VERSION)

all: build
