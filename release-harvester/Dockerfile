FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o release-harvester ./main/main.go

FROM alpine
COPY --from=builder /app/release-harvester /release-harvester
ENTRYPOINT ["/release-harvester"]