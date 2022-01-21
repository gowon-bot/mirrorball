FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build

EXPOSE 8080
CMD ["./gowon-indexer"]
