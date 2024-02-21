FROM node:18-alpine as build
LABEL org.opencontainers.image.authors="d.a.golovachev <d.a.golovachev@gmail.com>"

RUN apk update && apk add --no-cache git tzdata

ARG API_BASE_URL={{URL}}
ARG API_UPLOADS_URL={{UPLOADS}}

# Create app directory 
RUN mkdir -p /usr/src/app

WORKDIR /usr/src
COPY common ./common
WORKDIR /usr/src/app

# Bundle app source  
COPY fe/web ./web

#TimeZone
ENV TZ=Asia/Bishkek
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone  

# Install app dependencies
WORKDIR /usr/src/app/web
RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["/bin/sh", "-c", "node .output/server/index.mjs"]