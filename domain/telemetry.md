## 可观测领域相关知识

- opentelemetry 规范
  - log
  - metric
  - trace
- 主要知识点
  - 时序存储
  - 流式计算
    - flink
  - 数据采集
  - 监控告警
  - 智能诊断
  - 微服务 微服务框架
  - 容器 docker kubernetes istio/envoy
  - prometheus grafana zipkin opentelemetry zabbix
  - k8s/docker 云可观测系统 service mesh
  - ebpf

### 时序存储

- 无论是log，metric，trace 都是时序数据，而且数据量都非常大
- 明显是写多，读少的情况，那么对于这类的数据，如何高效存储与读取呢？这其实就是问题
- 目前可选型的方案有哪些，基本的原理是什么，对应的优势和劣势分别是什么
- log
  - elasticsearch
    - 优点：有全文检索，且有配套的 kibana 做展示
    - 缺点：集群规模大时，es的索引和存储成本较高，需要有较多的资源来维护集群的稳定性
  - loki
    - 优点：loki 是转为日志数据设计的时序数据库，类似 prometheus，但是用于日志场景，优化了资源使用，可以与 grafana 集成
    - 缺点：查询功能较弱，对于复杂查询的支持较弱
- metric
  - prometheus
    - tsdb, 默认是单机版本，需要了解具体实现
  - VictoriaMetrics
    - 需要对比这个与prometheus的区别，以及对应存储使用的是什么
- trace
  - Jaeger
    - Elasticsearch、Cassandra、kafka
  - Tempo

### 存储

- elasticsearch
  - 基本概念 cluster node shard index doc mapping setting
  - node 有主节点，协调节点，数据节点，其中主节点主要维护整个集群的元数据信息，协调节点负责提供 http 接口给外部使用，数据节点负责数据存储
  - 一个 index 在创建可以指定 主分片 与 副分片的数量，其中写入操作都是在主分片进行，数据写入完成后，可以在 副节点 进行读取，如果有多个主分片，那么会通过 doc id 哈希取余的方式选择主分片，所以index创建后无法新增主分片，否则涉及到的所有数据都需要调整与迁移
  - 倒排索引 是内容到文档 id 之间的映射关系
    - 单词词典 term dictionary，记录所有文档的单词，记录所有单词到倒排列表的关系
      - 单词词典比较大，一般可以通过 b+ 树 或者是 哈希拉链法 实现
    - 倒排列表 posting list，记录了单词到对应文档的集合，由倒排索引项组成
      - 文档 id
      - 词频 该单词再文档中出现的次数
      - 位置 position
      - 偏移 offset
- prometheus tsdb
- loki
- VictoriaMetrics 存储选的是什么
- influxdb
- hbase
