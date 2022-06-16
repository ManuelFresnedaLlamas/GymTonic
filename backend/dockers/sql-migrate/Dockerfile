FROM golang:1.16

ENV PATH="/go/bin:$PATH"

RUN go get github.com/rubenv/sql-migrate/... \
  && mkdir /migrations

WORKDIR /migrations

ENTRYPOINT ["tail", "-f", "/dev/null"]
