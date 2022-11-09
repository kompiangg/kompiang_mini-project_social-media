FROM golang:1.17-alpine as build

RUN apk update && apk upgrade
RUN apk add --no-cache  git

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN go build -o /app/main

# --------

FROM alpine:3.16.0

WORKDIR /app

# Web service
EXPOSE 8080

RUN apk update

COPY --from=build /app/main /app/main

CMD ["./main"]