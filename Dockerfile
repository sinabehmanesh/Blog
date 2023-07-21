FROM nginx:alpine

USER daemon

WORKDIR /app

COPY . .

COPY ./conf/blog.conf /etc/nginx/conf.d

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"