FROM golang:1.15.7-alpine3.12 as builder_stage
RUN apk add --no-cache git
WORKDIR /tmp/nasa-apod-telegram-bot
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/nasa-apod-telegram-bot .

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=builder_stage /tmp/nasa-apod-telegram-bot/out/nasa-apod-telegram-bot /app/nasa-apod-telegram-bot
CMD ["/app/nasa-apod-telegram-bot"]