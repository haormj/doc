## 安装kafka

kafka版本 2.11-1.1.0

查看使用java版本

https://kafka.apache.org/documentation/#java

下载安装jre,版本: 8u171

部署kafka

https://kafka.apache.org/quickstart

修改配置server.properties中的zookeeper配置

通过system运行kafka

```
[Unit]
Description=Kafka
After=storage.service
After=network.target

[Service]
Type=simple
LimitNOFILE=40960
LimitNPROC=40960
User=root
Group=root
Environment=JRE_HOME=/path/to/jre
WorkingDirectory=/path/to/kafka
ExecStart=/path/to/kafka/bin/kafka-server-start.sh config/server.properties
ExecStop=/path/to/kafka/bin/kafka-server-stop.sh

Restart=on-failure
RestartSec=6s

[Install]
WantedBy=multi-user.target
```

健康检查

```
kafka-topics.sh --list --zookeeper 127.0.0.1:2181
kafka-topics.sh --describe --zookeeper 127.0.0.1:2181 --topic xx
```

kafka集群

```
# server.properties中的broker.id不同即可
```



