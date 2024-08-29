FROM golang:1.23.0 AS builder

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go build -o main cmd/app/main.go

FROM golang:1.23.0

WORKDIR /app

COPY --from=builder /app .

COPY internal/config/certf/localhost.pem /localhost.pem
COPY internal/config/certf/localhost-key.pem  /localhost-key.pem

RUN chmod 600 /app/internal/config/certf/localhost.pem /app/internal/config/certf/localhost-key.pem


CMD ["/app/main"]

