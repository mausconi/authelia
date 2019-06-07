FROM alpine:3.9.4

WORKDIR /usr/app

RUN apk --no-cache add ca-certificates

ADD dist/authelia authelia
ADD dist/public_html public_html

EXPOSE 9091

VOLUME /etc/authelia
VOLUME /var/lib/authelia

CMD ["./authelia", "-config", "/etc/authelia/config.yml"]
