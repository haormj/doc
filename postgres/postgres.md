# Postgres

## 安装

通过docker部署postgres 9.6.9
https://hub.docker.com/r/_/postgres/

```shell
docker run \
  --restart always \
  --network host \
  --env POSTGRES_USER="postgres" \
  --env POSTGRES_PASSWORD="postgres" \
  --volume /opt/postgres/data:/var/lib/postgresql/data \
  --detach \
  --name postgres_v9.6.9 \
  postgres:9.6.9
```

## 使用

导出数据库

```shell
./pg_dump \
  --host 127.0.0.1 \
  --port 5432 \
  --username postgres \
  --password \
  --dbname db \
  > db.sql
```

导入数据库

```shell
./psql \
  --host 127.0.0.1 \
  --port 5432 \
  --username postgres \
  --password \
  --dbname db \
  --file /tmp/kong.sql
```

