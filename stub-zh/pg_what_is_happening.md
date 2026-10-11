## 用法

来源：

- [README.md](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/README.md)
- [pg_what_is_happening--1.1.sql](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/pg_what_is_happening--1.1.sql)
- [src/pg_what_is_happening.c](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/src/pg_what_is_happening.c)
- [pg_what_is_happening.control](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/pg_what_is_happening.control)

`pg_what_is_happening` 提供活动查询的实时执行与缓冲区指标。上游声称支持 PostgreSQL 9.4 至 18 及更新版本，但明确不保证动态库 ABI 稳定。

### 核心用法

```ini
shared_preload_libraries = 'pg_what_is_happening'
```

```sql
CREATE EXTENSION pg_what_is_happening;
SELECT * FROM what_is_happening.v1_status;
```

### 运行边界

必须预加载并重启，创建扩展要求超级用户。`what_is_happening.v1_status` 为每个活动计划节点返回一行。发布 SQL 将视图及其底层函数授权给 PUBLIC，需核对查询文本访问权限。可选后台工作进程提供 `/metrics`，默认监听本机 9187 端口，应保护该端点。`what_is_happening.min_cost_to_track` 过滤低成本查询。追踪查询数、节点数与查询文本长度上限的调整需要重启，启用状态与信号超时可重载。查询结束后指标消失，它不是历史统计库。
