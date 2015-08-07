FROM miek/alpine-armv6l

MAINTAINER Mark J. Becker <mjb@marb.ec>

RUN apk --update add openssl ca-certificates

COPY rpi-hello /
CMD ["/rpi-hello"]

EXPOSE 80
