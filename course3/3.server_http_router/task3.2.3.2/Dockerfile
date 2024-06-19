FROM golang:1.19 AS builder
WORKDIR /app
COPY /cmd/main.go .
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 go build -o main main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]