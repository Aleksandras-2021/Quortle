# Stage 1: Build Go binary
FROM golang:1.25 AS builder

WORKDIR /app

# Copy go modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o quortle .

# Stage 2: Minimal runtime
FROM alpine:latest
WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/quortle .
COPY frontend ./frontend
COPY words.txt .

EXPOSE 80 443

# Run the app
CMD ["./quortle"]