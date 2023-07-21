FROM nginx:alpine

USER daemon

WORKDIR /app

COPY . .

COPY ./conf/blog.conf /etc/nginx/conf.d

CMD [ "nginx", "-S", "reload" ]