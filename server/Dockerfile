FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /app/yata-server

FROM alpine:latest

COPY --from=builder /app/yata-server /app/yata-server

EXPOSE 8080

CMD ["/app/yata-server"]