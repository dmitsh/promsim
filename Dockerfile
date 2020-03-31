FROM alpine:3.11

RUN apk update && \
    apk upgrade && \
    rm -rf /var/cache/apk/*

COPY promsim /bin/promsim

USER nobody

ENTRYPOINT ["/bin/promsim", "target"]
