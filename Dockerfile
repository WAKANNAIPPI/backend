FROM golang:1.18.1-alpine

ENV http_proxy "http://cproxy.okinawa-ct.ac.jp:8080"
ENV https_proxy ${http_proxy}
ENV no_proxy localhost, 127.0.0.1, /var/run/docker.sock 
RUN  apk update && apk add git
RUN apk add openssh-cliant

EXPOSE 8080

COPY  ./docker ./src/
ENV ROOT=/tmp/backend
WORKDIR ${ROOT}

CMD [ "/bin/bash" ]
