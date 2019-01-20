## Kong使用说明

### 介绍

版本0.10.0,[文档链接](https://getkong.org/docs/0.10.x/).

kong管理侧默认端口为8001

目前配置api可以通过dashboard配置,如果对url做负载均衡,那么就需要使用upstream和target

upstream 上游服务器,也就是我们的业务服务器

target 为upstream组中的具体机器

### List upstreams

```
curl \
  --include \
  --request GET \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/"
```

### Add upstreams

```
curl \
  --include \
  --request POST \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/" \
  --data '{"name":"hello.world"}'
```

### Delete upstreams

```
curl \
  --include \
  --request DELETE \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/hello.world" 
```

### List targets(列出所有)

```
curl \
  --include \
  --request GET \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/hello.world/targets"
```

### List Active targets(列出激活)

```
curl \
  --include \
  --request GET \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/hello.world/targets/active/"
```

### Add target

```
curl \
  --include \
  --request POST \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/hello.world/targets" \
  --data '{
  	"target":"127.0.0.1:8080",
  	"weight":100
  }'
```

### Delete target(不是真正的删除,只是将权重修改为零,还可以在all targets中看到)

```
curl \
  --include \
  --request DELETE \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/upstreams/hello.world/targets/127.0.0.1:8080"
```

### Cluster

```
curl \
  --include \
  --request GET \
  --header "Content-Type: application/json" \
  --url "http://127.0.0.1:8001/cluster"
```

