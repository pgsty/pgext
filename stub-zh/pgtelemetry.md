## 用法

来源：

- [上游 README](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/README.md)
- [扩展 control 文件](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/pgtelemetry.control)
- [1.7 版 SQL](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/extension/pgtelemetry--1.7.sql)
- [生成的对象文档](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/doc/pgtelemetry.html)

`pgtelemetry` `1.7` 版是纯 SQL 监控集合。视图和函数汇总关系/数据库/表空间大小、连接和等待、锁、表访问和自动清理、`pg_stat_statements` 计时与缓冲区、WAL 速率、复制槽及备库延迟。

### 示例

先预加载并安装 `pg_stat_statements`，再把该扩展安装到固定的 `pgtelemetry` 模式：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgtelemetry;
SELECT * FROM pgtelemetry.database_size;
SELECT * FROM pgtelemetry.connections_by_state;
SELECT * FROM pgtelemetry.longest_running_active_queries LIMIT 10;
```

安装仅限超级用户。若干对象会暴露集群范围的查询文本、客户端地址、大小、锁、复制状态和其他运维元数据；`wal_telemetry()` 还会维护历史表。只应向监控角色授予所需的具体视图/函数，保护导出的指标，为 WAL 样本设置保留策略，并考虑轮询成本和 PostgreSQL 版本间统计目录变化。
