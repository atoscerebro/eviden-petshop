FROM golang:1.22-alpine as builder

WORKDIR /app

# Avoid downloading every time
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/server/main.go

# Final stage only with binaries
FROM alpine:3.20

COPY --from=builder /app/main /app/main

EXPOSE 4000

CMD ["/app/main"]
