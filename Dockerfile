FROM ubuntu:18.04

RUN apt-get update \
    && apt-get -y install curl vim

COPY promsim /bin/promsim
