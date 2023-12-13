FROM golang:1.19 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o semver ./main/main.go

FROM alpine:3.14
COPY --from=builder /app/semver /semver
ENTRYPOINT ["/semver"]
