## 安装

https://hub.docker.com/r/dockage/phppgadmin/

使用docker安装

```shell
docker run \
  --publish 905:80 \
  --env PHP_PG_ADMIN_SERVER_HOST=127.0.0.1 \
  --env PHP_PG_ADMIN_SERVER_PORT=5432 \
  --name phppgadmin \
  --restart always \
  --detach \
  dockage/phppgadmin:latest
```

