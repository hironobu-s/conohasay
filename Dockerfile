FROM alpine:latest
RUN apk --update add curl \
  && curl -sL https://github.com/hironobu-s/conohasay/releases/download/current/conohasay-linux.amd64.gz | zcat > /conohasay && chmod +x /conohasay \
  && apk del curl \
  && rm -rf /var/cache/apk/*
ENTRYPOINT ["/conohasay"]
