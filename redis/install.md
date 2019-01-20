## 安装Redis

https://redis.io/topics/quickstart

### centos7安装

```shell
yum install -y redis.x86_64
systemctl enable redis
systemctl start redis
# Redis server v=3.2.10 sha=00000000:0 malloc=jemalloc-3.6.0 bits=64 build=c8b45a0ec7dc67c6
```

### 自行安装

```
bind 127.0.0.1
protected-mode yes
port 6379
tcp-backlog 511
timeout 0
tcp-keepalive 300
daemonize no
supervised no
pidfile /var/run/redis_6379.pid
loglevel notice
logfile ""
databases 16
save 900 1
save 300 10
save 60 10000
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes
dbfilename dump.rdb
dir ./
slave-serve-stale-data yes
slave-read-only yes
repl-diskless-sync no
repl-diskless-sync-delay 5
repl-disable-tcp-nodelay no
slave-priority 100
appendonly no
appendfilename "appendonly.aof"
appendfsync everysec
no-appendfsync-on-rewrite no
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
aof-load-truncated yes
lua-time-limit 5000
slowlog-log-slower-than 10000
slowlog-max-len 128
latency-monitor-threshold 0
notify-keyspace-events ""
hash-max-ziplist-entries 512
hash-max-ziplist-value 64
list-max-ziplist-size -2
list-compress-depth 0
set-max-intset-entries 512
zset-max-ziplist-entries 128
zset-max-ziplist-value 64
hll-sparse-max-bytes 3000
activerehashing yes
client-output-buffer-limit normal 0 0 0
client-output-buffer-limit slave 256mb 64mb 60
client-output-buffer-limit pubsub 32mb 8mb 60
hz 10
aof-rewrite-incremental-fsync yes
```

redis.service

```
[Unit]
Description=Advanced key-value store
After=network.target
Documentation=http://redis.io/documentation, man:redis-server(1)

[Service]
Type=forking
ExecStart=/usr/bin/redis-server /etc/redis/redis.conf
PIDFile=/var/run/redis/redis-server.pid
TimeoutStopSec=0
Restart=always
User=redis
Group=redis

ExecStartPre=-/bin/run-parts --verbose /etc/redis/redis-server.pre-up.d
ExecStartPost=-/bin/run-parts --verbose /etc/redis/redis-server.post-up.d
ExecStop=-/bin/run-parts --verbose /etc/redis/redis-server.pre-down.d
ExecStop=/bin/kill -s TERM $MAINPID
ExecStopPost=-/bin/run-parts --verbose /etc/redis/redis-server.post-down.d

PrivateTmp=yes
PrivateDevices=yes
ProtectHome=yes
ReadOnlyDirectories=/
ReadWriteDirectories=-/var/lib/redis
ReadWriteDirectories=-/var/log/redis
ReadWriteDirectories=-/var/run/redis
CapabilityBoundingSet=~CAP_SYS_PTRACE

# redis-server writes its own config file when in cluster mode so we allow
# writing there (NB. ProtectSystem=true over ProtectSystem=full)
ProtectSystem=true
ReadWriteDirectories=-/etc/redis

[Install]
WantedBy=multi-user.target
Alias=redis.service
```

