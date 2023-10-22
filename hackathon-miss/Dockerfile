FROM golang:latest as build

COPY * .

RUN go build main.go

CMD ["./main"]