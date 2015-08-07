FROM miek/alpine-armv6l

RUN apk --update add openssl ca-certificates

COPY rpi-hello /
CMD ["/rpi-hello"]

EXPOSE 80
