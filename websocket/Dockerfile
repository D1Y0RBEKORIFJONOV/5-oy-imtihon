# Stage 1: Build
FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd

RUN go mod download
RUN go build -o /bin/ws

# Stage 2: Create the final image
FROM alpine:latest

# Устанавливаем CA сертификаты, если требуется
RUN apk add --no-cache ca-certificates

COPY --from=build /bin/ws /bin/ws

CMD ["/bin/ws"]
