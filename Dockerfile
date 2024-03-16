FROM golang:1.21-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

RUN go mod download && go build -o main .

FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

EXPOSE 9632

CMD ["./main"]