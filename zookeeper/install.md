## 安装zookeeper

通过docker部署

第一台

```shell
docker run \
  --hostname zoo1 \
  --network host \
  --restart unless-stopped \
  --cpu-shares 512 \
  --memory 512m \
  --memory-swap 512m \
  --log-driver json-file \
  --log-opt max-size=10m \
  --log-opt max-file=3 \
  --publish 2181:2181 \
  --publish 2888:2888 \
  --publish 3888:3888 \
  --env ZOO_MY_ID=1 \
  --env ZOO_SERVERS="server.1=192.168.1.1:2888:3888 server.2=192.168.1.2:2888:3888 server.3=192.168.1.3:2888:3888" \
  --detach \
  --volume /path/to/data:/data \
  --volume /path/to/datalog:/datalog \
  --name zoo1 \
  zookeeper:latest
```

第二台

```shell
docker run \
  --hostname zoo2 \
  --network host \
  --restart unless-stopped \
  --cpu-shares 512 \
  --memory 512m \
  --memory-swap 512m \
  --log-driver json-file \
  --log-opt max-size=10m \
  --log-opt max-file=3 \
  --publish 2181:2181 \
  --publish 2888:2888 \
  --publish 3888:3888 \
  --env ZOO_MY_ID=2 \
  --env ZOO_SERVERS="server.1=192.168.1.1:2888:3888 server.2=192.168.1.2:2888:3888 server.3=192.168.1.3:2888:3888" \
  --detach \
  --volume /path/to/zookeeper/data:/data \
  --volume /path/to/zookeeper/datalog:/datalog \
  --name zoo2 \
  zookeeper:latest
```

第三台

```shell
docker run \
  --hostname zoo3 \
  --network host \
  --restart unless-stopped \
  --cpu-shares 512 \
  --memory 512m \
  --memory-swap 512m \
  --log-driver json-file \
  --log-opt max-size=10m \
  --log-opt max-file=3 \
  --publish 2181:2181 \
  --publish 2888:2888 \
  --publish 3888:3888 \
  --env ZOO_MY_ID=3 \
  --env ZOO_SERVERS="server.1=192.168.1.1:2888:3888 server.2=192.168.1.2:2888:3888 server.3=192.168.1.3:2888:3888" \
  --detach \
  --volume /path/to/zookeeper/data:/data \
  --volume /path/to/zookeeper/datalog:/datalog \
  --name zoo3 \
  zookeeper:latest
```

健康检查

```shell
docker exec -it zoo1 /bin/bash
cd /zookeeper-3.4.12/bin
./zkServer.sh status
```

