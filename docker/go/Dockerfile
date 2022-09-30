FROM golang:1.18.1-alpine

#本番環境ではproxy通さない

ENV no_proxy localhost, 127.0.0.1, /var/run/docker.sock 
RUN  apk update && apk add git

EXPOSE 8080

ENV ROOT=/go/src
WORKDIR ${ROOT}
