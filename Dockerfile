FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build multiple binaries
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/price_indexer ./cmd/price_indexer
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/indexer ./cmd/indexer
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./api

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binaries from the builder stage
COPY --from=builder /app/bin/indexer .
COPY --from=builder /app/bin/price_indexer .
COPY --from=builder /app/bin/api .