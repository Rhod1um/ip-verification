# Dockerfile for app2
FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o app2 ./server
CMD ["/app/app2"]
EXPOSE 8080