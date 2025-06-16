FROM golang:1.24.2-alpine

WORKDIR /src
COPY . .

RUN go build -o server

CMD ["./server"]