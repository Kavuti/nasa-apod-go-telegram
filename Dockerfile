FROM golang:1.15.7-alpine3.12
RUN mkdir /app
ADD main.go /app
WORKDIR /app
CMD go install -o nasa-apod-telegram-bot
WORKDIR $GOPATH/bin
ENTRYPOINT ["./nasa-apod-telegram-bot"]