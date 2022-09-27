FROM golang:1.18.1-alpine

ENV http_proxy "http://cproxy.okinawa-ct.ac.jp:8080"
ENV https_proxy ${http_proxy}
ENV no_proxy localhost, 127.0.0.1
RUN  apk update && apk add git

EXPOSE 8080

COPY ./aiueo.txt /tmp
ENV ROOT=backend
WORKDIR ${ROOT}

CMD [ "/bin/bash" ]
