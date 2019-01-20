### Kafka

Kafka的Log存储解析

https://www.zybuluo.com/jewes/note/64450



server.properties中配置了 log.dirs

其中一个topic可以有多个partition,其中每一个都是会存储到磁盘上,目录名称为

<topic_name>-<partition_id>



00000000009100249402.index
00000000009100249402.log
00000000009100249402.timeindex



kafka日志保留

空间

当设置的大小为 1G, 日志文件最大为1G,那么当日志总大小-阈值 > 1G的时候,就会删除最老的文件

时间 

设置日志的保留时间为7天,用当前时间戳-日志最后的时间戳 > 7天,就删除该日志文件



其中任意一个条件满足都会删除数据日志

