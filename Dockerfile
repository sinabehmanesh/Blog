FROM golang:alpine as build

USER root

WORKDIR /app

COPY . .

RUN go build -o blog

FROM alpine:latest 

WORKDIR /app

COPY --from=build /app/blog .

EXPOSE 3000

CMD [ "/app/blog" ]

# CMD ["nginx", "-g", "daemon off;"]