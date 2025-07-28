FROM golang:latest

WORKDIR /app

COPY load_balancer.go /app/

EXPOSE 8000

CMD [ "go", "run", "load_balancer.go" ]