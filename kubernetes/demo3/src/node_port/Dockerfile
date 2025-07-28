FROM golang:latest

WORKDIR /app

COPY node_port.go /app/

EXPOSE 8080

CMD ["go", "run", "node_port.go"]