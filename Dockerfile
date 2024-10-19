FROM golang:1.19 AS builder

WORKDIR /mnt/engine

COPY . .

RUN make build

FROM golang:1.19

WORKDIR /mnt/engine

COPY --from=builder /mnt/engine/build .

CMD ["./bin/boilerplate"]
