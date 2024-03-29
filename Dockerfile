# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# Set execute permissions for scripts
RUN chmod +x wait-for.sh

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY  app.env .
COPY  start.sh .
COPY  wait-for.sh .



EXPOSE 8082
CMD ["/app/main"]
ENTRYPOINT ["sh", "/app/start.sh"]
