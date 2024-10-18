FROM golang:1.19 AS builder

WORKDIR /mnt/engine

COPY . .

RUN ./dist/bin/build.sh

WORKDIR /mnt/engine/build

CMD ["./bin/docker-start.sh"]
