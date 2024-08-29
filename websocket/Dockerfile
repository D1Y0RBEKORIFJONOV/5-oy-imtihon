FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY . .

WORKDIR /src/cmd

RUN go mod download

RUN go build -o /bin/ws

FROM scratch

COPY --from=build /bin/ws /bin/ws

CMD [ "/bin/ws" ]
