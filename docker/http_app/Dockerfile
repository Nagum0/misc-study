FROM golang:latest

ENV makima=yes \
    reze=yes

WORKDIR /app

COPY . /app/

EXPOSE 8080

RUN pwd

CMD [ "go", "run", "main.go" ]