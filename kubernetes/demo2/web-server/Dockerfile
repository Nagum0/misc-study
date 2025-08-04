FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY src/ ./

EXPOSE 8080

RUN mkdir -p bin
RUN go build -o bin/server .

CMD ["bin/server"]