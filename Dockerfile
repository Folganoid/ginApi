FROM golang:latest

LABEL maintainer=""sdfsdf

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

RUN cd ./cmd/api; go build -o api

EXPOSE $PORT

CMD ["./cmd/api/api"]

