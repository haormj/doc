## nginx

### 安装

https://nginx.org/en/linux_packages.html

centos 7.2

```shell
touch /etc/yum.repos.d/nginx.repo
vi /etc/yum.repos.d/nginx.repo

[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/centos/7/$basearch/
gpgcheck=0
enabled=1

wget http://nginx.org/keys/nginx_signing.key
rpm --import nginx_signing.key

yum install -y nginx.x86_64
```

### 配置

#### 配置健康检查

```
server {
    listen       80;
    location / {
        return 200;
    }
}
```

#### nginx打印http request和response头

http://xyduan.com/logging-request-reponse-header-to-nginx/

```
resonse header
$upstream_http_resonseHeaderName
例如
$upstream_http_x_rqid

request header
$http_requestHeaderName
例如
$http_x_forwarded_for
```

#### nginx 配置静态页面

```
server {
    listen       80;
    server_name  localhost;
    root /path/to/root;
    index index.html;
}
```

#### 增加证书http server配置

```
server {
    listen 443;
    server_name localhost;
    ssl on;
    ssl_certificate   /path/to/crt.pem;
    ssl_certificate_key  /path/to/key.pem;
    ssl_session_timeout 5m;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers on;
    location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $host;
        proxy_ssl_verify off;
        proxy_pass https://localhost:443;
    }
}
```

#### nginx配置tcp转发

```
stream {
    server {
        listen 8080;
        proxy_connect_timeout 3s;
        proxy_timeout 24h;
        proxy_pass localhost:8080;
    }
}
```
#### nginx配置tcp域名动态解析(commercial subscription)

http://nginx.org/en/docs/stream/ngx_stream_upstream_module.html#server

```
stream {
    resolver 127.0.0.1 valid=30s ipv6=off;

    upstream backend {
        server haormj.com:3306 resolve;
    }

    server {
        listen 3306;
        proxy_connect_timeout 3s;
        proxy_timeout 86400s;
        proxy_pass backend;
    }
}
```



