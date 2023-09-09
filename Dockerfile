FROM golang:1.21-rc-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/bin/ ./cmd/server/...

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/ .
#COPY --from=builder /app/configs/ ./configs/
EXPOSE 3333 3334
CMD ["./cmd"]
