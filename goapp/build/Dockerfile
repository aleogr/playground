FROM cgr.dev/chainguard/static:latest

ARG TARGETPLATFORM

COPY $TARGETPLATFORM/app1 /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]