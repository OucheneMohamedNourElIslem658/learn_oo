FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .

COPY .env .

EXPOSE 8080
EXPOSE 50051

CMD ["./backend"]