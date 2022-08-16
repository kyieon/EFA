FROM golang:1.17.2-alpine

WORKDIR /app

COPY cli/go.* ./
RUN go mod download

COPY cli/*.go ./
COPY cli/cmd ./cmd

RUN go build -o efa
RUN mv efa /usr/bin

EXPOSE 8080