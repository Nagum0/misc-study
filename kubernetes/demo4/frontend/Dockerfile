FROM golang:1.24

WORKDIR /app

# Add .dockerignore file later for the helm chart files
COPY . /app

EXPOSE 8082

CMD ["go", "run", "src/main.go"]
