### systemd

#### 查看systemd文档

```shell
info systemd
```

#### service unit example

```c
[Unit]
Description=Test Service
After=storage.service
After=network.target

[Service]
LimitNOFILE=40960
LimitNPROC=40960
User=test
Group=test
Type=simple
WorkingDirectory=/path/to/work
ExecStart=/path/to/exec
MemoryLimit=2G

Restart=on-failure
RestartSec=2s

[Install]
WantedBy=multi-user.target
```

#### systemd.resource-control

```shell
# CPUQuota/MemoryLimit
info systemd.resource-control
```

#### systemd.exec 

```shell
# LimitNOFILE/LimitNPROC
# LimitNOFILE Number of File Descriptors
# LimitNPROC Number of Processes
info systemd.exec 
```

