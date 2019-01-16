FROM golang:1.11-stretch

ADD . /go/src/github.com/7phs/coding-challenge-iban
WORKDIR /go/src/github.com/7phs/coding-challenge-iban

ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/ && \
    echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 ~/.ssh/id_rsa && \
    ssh-keyscan -H gitlab.teamc.io >> ~/.ssh/known_hosts

ARG GIT_HASH
RUN dep ensure -v && \
    go generate ./... && \
    go build -o profile-files-service -ldflags "-X github.com/7phs/coding-challenge-iban/cmd.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/7phs/coding-challenge-iban/cmd.GitHash=`${GIT_HASH}`"

FROM debian:stretch

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && apt-get clean

EXPOSE 8080
WORKDIR /root/
COPY --from=0 /go/src/github.com/7phs/coding-challenge-iban/iban-validator .
COPY data/ ./data

CMD ["./iban-validator"]
