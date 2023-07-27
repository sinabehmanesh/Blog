FROM golang:alpine as build

USER root

WORKDIR /app

COPY . .

RUN go build -o blog

FROM alpine:latest 

COPY --from=build /app/blog .

EXPOSE 3000

CMD [ "/blog" ]

# CMD ["nginx", "-g", "daemon off;"]