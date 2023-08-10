FROM golang:alpine as build

USER root

WORKDIR /app

COPY . .
COPY .env .

RUN go build -o blog main.go

EXPOSE 3000

CMD [ "/app/blog" ]