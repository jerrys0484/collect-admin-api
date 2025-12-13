FROM 192.168.0.102:5000/alpine:7c0b08d6-6b90-4fa7-82cb-ceabed56c5f6

WORKDIR /app/

COPY . .

EXPOSE 8879

ENTRYPOINT ["./collect-admin-api", "--gf.gcfg.file=./collect-admin-api.yaml"]
