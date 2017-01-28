FROM ubuntu:xenial

RUN apt-get update && \
apt-get -y upgrade && \
rm -rf /var/cache/apt/*

COPY squidbot-brain /usr/bin/squidbot-brain

ENTRYPOINT ["/usr/bin/squidbot-brain"]
