FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/snippetbox ./cmd/web

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/snippetbox .
COPY --from=builder /app/ui ./ui

RUN adduser -D appuser
USER appuser

EXPOSE 4000

CMD ["./snippetbox"]