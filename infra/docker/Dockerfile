FROM golang:1.18.0-buster AS build

WORKDIR /app

RUN ls
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /matching-timestamps ./cmd/main.go

