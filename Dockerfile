FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o counter-service ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/counter-service .
COPY --from=builder /app/.env .

EXPOSE 3000

CMD ["./counter-service"]