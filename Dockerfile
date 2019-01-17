FROM golang:1.12-rc-stretch

ENV SRC=/go/src/github.com/7phs/coding-challenge-iban

ADD . ${SRC}
WORKDIR ${SRC}

RUN make build

FROM debian:stretch

ENV SRC=/go/src/github.com/7phs/coding-challenge-iban

RUN apt-get update \
    && apt-get install -y ca-certificates \
    && apt-get clean

EXPOSE 8080

WORKDIR /root/
COPY --from=0 ${SRC}/iban-validator .
COPY --from=0 ${SRC}/data/ ./data

CMD ["./iban-validator", "run"]
