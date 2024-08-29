FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd/user

RUN go mod download

RUN go build -o /bin/hello

FROM scratch

COPY --from=build /bin/hello /bin/hello

CMD [ "/bin/hello" ]
