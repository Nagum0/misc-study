FROM golang:1.24

WORKDIR /app

COPY . /app/

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "server.go"]