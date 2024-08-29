FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd/booking

RUN go mod download
RUN go build -o /bin/hello

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /bin/hello /bin/hello

CMD [ "/bin/hello" ]
