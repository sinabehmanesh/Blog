FROM nginx:alpine

USER root

WORKDIR /app

COPY . .

COPY ./conf/default.conf /etc/nginx/conf.d

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]