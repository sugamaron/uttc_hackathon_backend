FROM golang:latest as build

ENV GOPATH=

COPY . .

RUN go build main.go

CMD ["./main"]