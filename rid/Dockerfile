FROM golang:1.7.5-alpine


#  Glide
#-----------------------------------------------
ENV GLIDE_VERSION 0.11.1

RUN apk add --no-cache -U --virtual .build-deps \
    curl \
  && curl -fL "https://github.com/Masterminds/glide/releases/download/v$GLIDE_VERSION/glide-v$GLIDE_VERSION-linux-amd64.zip" -o glide.zip \
  && unzip glide.zip \
  && mv ./linux-amd64/glide /usr/local/bin/glide \
  && rm -fr ./linux-amd64 \
  && rm ./glide.zip \
  && apk del .build-deps \
  \
  && apk add --no-cache -U --virtual .glide-deps \
    git


#  glibc (https://github.com/sgerrand/alpine-pkg-glibc)
#-----------------------------------------------
# @see http://qiita.com/tkusumi/items/1668df2fc4cf61a68282
RUN apk add --no-cache -U --virtual .build-deps \
    ca-certificates \
    wget \
  && wget -q -O /etc/apk/keys/sgerrand.rsa.pub \
    https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub \
  && wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-2.25-r0.apk \
  && apk del .build-deps \
  && apk add glibc-2.25-r0.apk


#  Library
#-----------------------------------------------
RUN apk add --no-cache -U --virtual .build-deps \
    build-base \
  && go get -u \
    github.com/golang/lint/golint \
    github.com/mattes/migrate \
  && apk del .build-deps

RUN apk add --no-cache -U \
    bash \
    build-base \
    coreutils \
    make


#  App
#-----------------------------------------------
ENV APP_DIR /go/src/github.com/creasty/go-server-boilerplate
WORKDIR $APP_DIR
RUN ln -sf $APP_DIR /app
