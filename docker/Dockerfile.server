### Stage 1
FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o coineus-server ./cmd/server

### Stage 2
FROM scratch
COPY --from=builder /app /server
CMD ["/server/coineus-server"]