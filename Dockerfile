FROM golang:latest

COPY . /app
WORKDIR /app/cmd/news

RUN go mod download

RUN go build -o ./server

EXPOSE 5000

CMD ["./server"]