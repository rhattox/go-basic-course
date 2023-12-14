FROM golang:1.21.5

WORKDIR /home/go/src

COPY ./src ./

RUN go build main.go