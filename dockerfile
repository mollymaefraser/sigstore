FROM alpine:3.14

RUN mkdir /app

ADD ./sigstore /app

CMD ["/app/sigstore"]