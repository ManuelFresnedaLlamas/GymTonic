FROM golang:1.16-stretch

ENV DEBIAN_FRONTED noninteractive

ENV PGPASSWORD="postgres"

RUN go get github.com/go-delve/delve/cmd/dlv \
 && mkdir /src && mkdir /dst && mkdir /dst/gym && mkdir /dst/template \
 && apt-get update && apt-get -y -q install libreoffice libreoffice-writer libreoffice-core python3 && apt-get -q -y remove libreoffice-gnome && apt-get -q -y install postgresql postgresql-contrib
EXPOSE 8000 40000


WORKDIR /dst
