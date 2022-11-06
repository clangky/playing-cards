FROM golang:1.19.3-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8080