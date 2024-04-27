FROM golang:1.22-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

RUN go mod download && go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8787

CMD ["./main"]
