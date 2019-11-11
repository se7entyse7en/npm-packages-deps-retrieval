FROM golang:1.13-alpine AS builder

WORKDIR /api
COPY . .
RUN go build -o npm-packages-deps-retrieval main.go

FROM alpine:3.10

WORKDIR /api
COPY --from=builder /api/npm-packages-deps-retrieval .

ENTRYPOINT ["./npm-packages-deps-retrieval"]
