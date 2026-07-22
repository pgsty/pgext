## 用法

来源：

- [Official README](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/README.md)
- [Extension SQL](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/pg_track_slow_queries--1.0.sql)
- [Hook and GUC implementation](https://github.com/julmon/pg_track_slow_queries/blob/eeea6f7c65d2e59cfe01e71329016e0bc364bdca/pg_track_slow_queries.c)

pg_track_slow_queries 是一个预加载扩展，会把超过时长阈值的语句、执行指标以及可选 JSON 计划写入专用服务端文件。SQL 函数可以读取和重置该文件。上游把项目标记为仍在开发，所核验源码来自 2019 年。

### 核心流程

预加载该库，配置正阈值，重启 PostgreSQL，并在用于检查的数据库中创建扩展：

```ini
shared_preload_libraries = 'pg_track_slow_queries'
pg_track_slow_queries.log_min_duration = 1000
pg_track_slow_queries.log_plan = on
pg_track_slow_queries.compression = on
pg_track_slow_queries.max_file_size = 1048576
```

```sql
CREATE EXTENSION pg_track_slow_queries;

SELECT datetime, duration, username, dbname, query, plan
FROM pg_track_slow_queries()
ORDER BY datetime DESC;

SELECT pg_track_slow_queries_reset();
```

reset 调用会截断共享历史；如果需要保留证据，应先导出。

### 重要对象

- `pg_track_slow_queries` 返回时间戳、时长、用户、应用、数据库、临时块、缓存命中率、影响行数、查询文本和 JSON 计划。
- `pg_track_slow_queries_reset` 清空存储文件。
- `pg_track_slow_queries.log_min_duration` 设置毫秒阈值；-1 禁用捕获。
- `pg_track_slow_queries.compression` 控制行压缩。
- `pg_track_slow_queries.max_file_size` 以 KB 限制文件大小；-1 表示无限制。
- `pg_track_slow_queries.log_plan` 控制计划捕获。
- `pg_track_slow_queries.cost_analyze` 在超过成本阈值时启用类似 EXPLAIN ANALYZE 的工作；-1 禁用。

### 安全与性能说明

捕获的 SQL 和计划可能含有字面值、标识符及敏感负载细节，而 reset 函数会销毁共享证据。必须限制函数执行权限和服务端文件访问。扩展不跟踪 utility 语句，也不记录预备语句参数值。计划捕获、压缩、文件大小强制限制，尤其是 analyze 风格执行，都可能带来可观开销；大范围启用前要按实际配置做基准测试并监控磁盘。
