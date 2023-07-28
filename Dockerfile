FROM golang:alpine as build

USER root

WORKDIR /app

COPY . .

RUN go build -o blog

EXPOSE 3000

CMD [ "/app/blog" ]