## 安装kong

https://konghq.com/install/

安装kong-dashboard

```
docker run \
  --net host \
  --log-opt max-size=10m \
  --log-opt max-file=10 \
  --cpu-shares 512 \
  --memory 512m \
  --detach \
  --name kong-dashboard \
  pgbi/kong-dashboard \
  start \
  --port 8080 \
  --kong-url http://127.0.0.1:8001 \
  --basic-auth username=password 
```

