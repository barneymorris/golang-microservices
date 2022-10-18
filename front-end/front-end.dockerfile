FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

CMD ["go", "run", "/app/cmd/web"]