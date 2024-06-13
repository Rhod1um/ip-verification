# Dockerfile for app1
FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o app1 .
CMD ["/app/app1"]
EXPOSE 7777