FROM golang:1.15.7-alpine3.12 as builder_stage
RUN apk add --no-cache git
WORKDIR /app/nasa-apod-telegram-bot
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENTRYPOINT ["go", "run", "main.go", "db.go"]