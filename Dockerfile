FROM golang:1.15.5-alpine3.12
ADD main.go /main.go
ENTRYPOINT ["go", "run", "main.go"]