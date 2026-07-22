## 用法

来源：

- [官方 README](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/README.md)
- [官方 FDW 处理器](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/ch_fdw.c)
- [官方扩展控制文件](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/ch_fdw.control)

`ch_fdw` 是实验性的只读外部数据包装器，允许 PostgreSQL 通过 ClickHouse 原生 TCP 协议进行查询。实现通过 CGo 组合 PostgreSQL C 内部接口和 Go，并支持选定的过滤与聚合下推。

### 核心流程

创建扩展和服务器，然后把 PostgreSQL 外部表映射到已有的 ClickHouse 表，两端列名与类型必须兼容：

```sql
CREATE EXTENSION ch_fdw;

CREATE SERVER clickhouse_server
  FOREIGN DATA WRAPPER ch_fdw
  OPTIONS (host '127.0.0.1', port '9000');

CREATE FOREIGN TABLE click_events (
  event_name text,
  event_count bigint,
  observed_at timestamptz
)
SERVER clickhouse_server
OPTIONS (table 'events');

SELECT event_name, sum(event_count)
FROM click_events
WHERE observed_at >= now() - interval '1 day'
GROUP BY event_name;
```

应使用 `EXPLAIN (VERBOSE)` 确认哪些过滤与聚合确实发送给 ClickHouse；不支持的表达式会留在本地，或在反解析与转换边界失败。

### 已实现边界

- 处理器提供扫描、重新扫描、分析、解释和上层路径回调，但没有外部修改回调，因此不支持 `INSERT`、`UPDATE` 和 `DELETE`。
- 上游记录了 `WHERE` 和聚合下推。源码注释排除了连接、参数、排序下推、分组集、有序集聚合、聚合 `FILTER` 以及若干表达式形式。
- 当前代码使用的服务器选项是 `host` 和 `port`；验证器是空操作，因此 DDL 阶段不会拒绝拼写错误或不安全的选项。

### 运维注意事项

- 上游把项目标为实验性，说明 MessageBird 已不再主动开发或使用，并且只记录了 PostgreSQL 13、ClickHouse 20.3.19.4 与 Ubuntu 20.04 的测试。应把仓库视为参考实现，而不是受维护的兼容承诺。
- 当前连接代码强制打开 ClickHouse 驱动调试模式，并在未提供用户名时使用 ClickHouse 默认用户；其选项解析没有形成已记录的生产凭据流程。连接敏感端点前应审计日志和代码。
- 查询下推会跨越 PostgreSQL 与 ClickHouse 的类型、时区、排序规则、空值、数值和函数语义。应使用代表性边界值比较下推与不下推结果。
- CGo、C 内存、Go 运行时和 PostgreSQL 内存上下文共享同一个后端进程。应在准确的服务器构建上压力测试取消、重连、错误、大结果集、重复扫描和后端关闭。
