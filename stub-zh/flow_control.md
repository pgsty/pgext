## 用法

来源：

- [Amazon RDS Multi-AZ DB 集群概念](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
- [Amazon RDS for PostgreSQL 扩展版本](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

flow_control 是 AWS 为 RDS for PostgreSQL Multi-AZ DB 集群托管的扩展。reader 上的后台工作进程把 apply lag 报告给 writer；只要任一 reader 超过配置阈值，writer 工作进程就会在事务结束时注入延迟，以降低写入压力。

### 运行与观测

AWS 文档列出的版本为 1.0。在受支持的引擎版本和区域中，把该扩展加入 DB 参数组的预加载设置，按 RDS 要求重启实例，在数据库中创建扩展并查看统计信息：

```sql
CREATE EXTENSION flow_control;

SHOW flow_control.target_standby_apply_lag;

SELECT * FROM get_flow_control_stats();

SELECT pid, wait_event_type, wait_event, query
FROM pg_stat_activity
WHERE wait_event_type = 'Extension';
```

文档给出的默认延迟阈值是两分钟。受到节流的会话会在活动视图和 Performance Insights 中显示 Extension wait event。参数是否可用以及修改流程由 RDS 引擎版本和 DB 参数组控制。

### 适用范围与代价

它不是面向自建 PostgreSQL、Aurora PostgreSQL 或普通主备环境的可移植扩展，而是 RDS Multi-AZ DB 集群服务的一部分，并依赖 AWS 托管的实例间通信。

flow control 最适合短事务、高并发的 OLTP 工作负载。若延迟来自长时间批处理事务，其效果较弱；启用节流时还会降低 writer 吞吐量，也不能保证延迟始终低于目标。AWS 说明可从预加载列表移除 flow_control 并重启 DB 实例来关闭它；修改生产设置前应测试写入延迟和故障转移影响。
