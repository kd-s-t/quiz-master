FROM ubuntu:20.04
LABEL maintainer="kendantinio@gmail.com"
LABEL version="1.0"
LABEL description="Quiz Master for Quipper"

WORKDIR /home
COPY . /home
RUN apt-get update \
    apt-get -y install curl \
    apt-get -y install sqlite3 \
    apt-get -y install vim \
    apt-get --no-cache git
RUN cd ~
RUN curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
RUN sha256sum go1.16.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
RUN echo PATH=$PATH:/usr/local/go/bin >> ~/.profile
RUN source ~/.profile
RUN chmod u+x /home/bin/*

ENTRYPOINT ["/bin/project"]

FROM golang:1.18-bullseye as builder
RUN go version

RUN chmod u+x /home
COPY . /home

ENTRYPOINT [ "/home" ]