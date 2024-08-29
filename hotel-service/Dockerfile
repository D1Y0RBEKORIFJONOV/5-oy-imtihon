# Stage 1: Build
FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd/hotel

RUN go mod download
RUN go build -o /bin/hello

# Stage 2: Create the final image
FROM alpine:latest

# Устанавливаем CA сертификаты для поддержки SSL/TLS
RUN apk add --no-cache ca-certificates

COPY --from=build /bin/hello /bin/hello

CMD ["/bin/hello"]
