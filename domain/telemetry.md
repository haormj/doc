## 可观测领域相关知识

- opentelemetry 规范
  - log
  - metric
  - trace
- 主要知识点
  - 时序存储
  - 流式计算
  - 数据采集
  - 监控告警
  - 智能诊断
  - 微服务 微服务框架
  - 容器 docker kubernetes istio/envoy
  - prometheus grafana zipkin opentelemetry zabbix
  - k8s/docker 云可观测系统 service mesh

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
- loki
- hbase
