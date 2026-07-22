## 用法

来源：

- [官方 README](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/README.md)
- [官方 SQL API](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/sql/statsd.sql)
- [官方 UDP 客户端实现](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/src/statsd.c)
- [官方 PGXN 版本](https://pgxn.org/dist/statsd/0.2.0/)

`statsd` 版本 `0.2.0` 从 PostgreSQL 后端通过 UDP 直接向 StatsD 端点发送计数器、仪表与计时指标。它是一个归档于 2014 年、采用即发即弃网络行为的客户端；只适用于可以接受指标丢失，而且数据库角色无法滥用出站访问的场景。

### 核心流程

安装扩展，并向可达的 StatsD 监听器发送指标：

```sql
CREATE EXTENSION statsd;

SELECT statsd.add_timing('127.0.0.1', 8125, 'checkout.duration_ms', 70);
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.completed');
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.items', 5);
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.sampled', 5, 0.25);
SELECT statsd.set_gauge('127.0.0.1', 8125, 'workers.busy', 3);
```

所有函数都返回 void。每次调用都传入主机与端口，而不是使用扩展配置。

### 函数形式

- `statsd.add_timing(host, port, metric, milliseconds)` 发送 `ms` 指标。
- `statsd.increment_counter(host, port, metric)` 递增一。
- `statsd.increment_counter(host, port, metric, value)` 发送显式计数器增量。
- `statsd.increment_counter(host, port, metric, value, sample)` 增加采样率。
- `statsd.set_gauge(host, port, metric, value)` 提供整数与双精度重载。

### 运维注意事项

每次调用都会在查询后端中解析 DNS、打开 UDP 套接字并发送一个数据报。UDP 没有送达确认；实现会把解析或发送故障记录为服务器警告，但仍返回 void，因此 SQL 成功不代表指标已送达。SQL 错误地将这些有网络副作用的函数声明为 `IMMUTABLE`；绝不能在索引、生成字段、约束或可能被规划器折叠的表达式中使用。应撤销过宽的 EXECUTE 权限，防止任意 DNS 查询、网络外发或指标注入。指标名也不会被清理。生产工作负载更适合使用仍在维护的异步 exporter 或日志/遥测流水线。
