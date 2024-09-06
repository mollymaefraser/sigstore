FROM alpine:3.14

RUN mkdir /app

ADD ./cmd/sigstore /app

CMD ["/app/sogstore]