FROM golang:1.18.6-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main

ENV CGO_ENABLED 0

EXPOSE 8000 8001 5432

# CMD ["go","run","main.go"]

CMD ./main

