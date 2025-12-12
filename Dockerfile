FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather ./cmd/server

FROM alpine:3.20

# Create a non-root user and group
RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

# Copy binary and set ownership to the non-root user
COPY --from=builder /app/weather /app/weather
RUN chown -R app:app /app

# Switch to non-root user
USER app

EXPOSE 8080
CMD ["./weather"]

