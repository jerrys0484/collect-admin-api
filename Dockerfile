FROM alpine:latest

WORKDIR /app/

COPY bin .

RUN chmod +x bin/collect-admin-api

EXPOSE 8879

ENTRYPOINT ["./collect-admin-api",'--gf.gcfg.file="./collect-admin-api.yaml"']
