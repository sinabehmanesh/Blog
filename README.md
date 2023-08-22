# Blog
This repository hosts the source code of the Sina behmanesh personal website.
it is currently under construction.
all the API and Frontend source codes are hosted here.

> Reach the website at [sinabehmanesh.com](https://sinabehmanesh.com)

## Tech stack

# API
API is written in **Golang**.
Using **MUX** for serving and server-side routings. it's simple and easy.
Mysql connection and CRUD are handled by Golang native **Mysql-driver** and **GORM** will take its place soon.

This Backend has integrations with:
* MySQL

> More integrations will be added soon.



# Frontend
Frontend is just **HTML**, you can find more on index.html.
I do not have any Frontend knowledge and this ancient theme(it's like 90s) for the website suits it. I like it!


# Database
**Mysql-server** version **8** is operational as database. its connection is exposed over PORT and not Socket!
an automated backup will be activated soon, I am planning to implement it using **SHELL** script.
Mysql deployment is bare-metal and not managed on containerized. it is a single instance of MySQL serving data to API.

# Server
on the server we are exposing this website using **Nginx**, I have defined an upstream and some proxy_pass so the external request finds its safe Path to the website Edge. I also customized the Nginx logging in case I wanted to aggregate logs in the future.

# CICD
deployment pipeline is using **Github Action** and selfhosted runners to **BUILD - DEPLOY - TEST** application.
API and Frontend are being deployed using docker, checkout Dockerfile and workflow to find more about the details.

> All secrets such as database credentials and URLs are included inside the action secrets.


---

that's it, I'm working on this website in my free time. It's cool and I have more items to add, like the latest news, Syslog and etc.
thank you!