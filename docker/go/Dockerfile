FROM golang:1.18.1-alpine

#本番環境ではproxy通さない
ENV http_proxy "http://cproxy.okinawa-ct.ac.jp:8080"
ENV https_proxy ${http_proxy} 
ENV no_proxy localhost, 127.0.0.1, /var/run/docker.sock 
RUN  apk update && apk add git

EXPOSE 8080

COPY  ./ /tmp/
ENV ROOT=/tmp/backend
WORKDIR ${ROOT}

CMD [ "go", "run", "main.go" ]

