FROM golang:alpine AS builder

RUN apk update && apk add --no-cache tzdata && apk add --no-cache git

WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux go build -o bin/music-library-api ./cmd

FROM alpine:3.18

RUN apk update

COPY --from=builder /app/bin/dashboard-service /usr/local/bin/music-library-api
COPY --from=builder /app/.env /usr/local/bin/.env

CMD ["/usr/local/bin/music-library-api"]

