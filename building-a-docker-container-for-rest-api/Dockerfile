FROM golang:latest

LABEL maintaner="kuma <daniel099d@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=:8000

RUN go build

CMD [ "./building-a-docker-container-for-rest-api" ]