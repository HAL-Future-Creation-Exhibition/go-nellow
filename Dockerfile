FROM golang:latest

ADD ./server .

EXPOSE 8080

ENTRYPOINT ["./server"]
