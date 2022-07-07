FROM ubuntu:rolling

ARG ARCH
COPY dist/${ARCH}/hello-uname-server /usr/local/bin

ENTRYPOINT /usr/local/bin/hello-uname-server
EXPOSE 8080/tcp
