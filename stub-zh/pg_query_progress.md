## 用法

来源：

- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/README.md)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress.control)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress--1.0.sql)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pg_query_progress.c)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/pgqp_compat.h)
- [Official source](https://github.com/ken19990709/pg_query_progress/blob/b0009ff4de03de70ac83ffd23ed1e696e9b81bc7/LICENSE)

`pg_query_progress` 通过视图、函数及可选的 NOTICE 进度条，报告正在运行的 SELECT 语句的成本加权进度估计。control 与安装 SQL 版本为 1.0；仓库中的 4.0.0 标签描述源码分发版本。

### 启用与监控

按服务器主版本选择 PG15、PG16、PG17 或 PG18 分支。此次审阅的主分支面向 PostgreSQL 20 开发版，并非 PostgreSQL 19。测试平台为 Linux。使用监控对象前必须共享预加载并重启服务器。

```ini
shared_preload_libraries = 'pg_query_progress'
```

```sql
CREATE EXTENSION pg_query_progress;
GRANT EXECUTE ON FUNCTION pg_query_progress() TO pg_monitor;
GRANT SELECT ON pg_query_progress TO pg_monitor;
SET pg_query_progress.bar = on;
SELECT pid, state, progress, tuples_processed, age_ms
FROM pg_query_progress WHERE state = 'running';
```

应保留其他预加载项，并从另一会话轮询长查询。查询正文可能包含敏感信息，函数与视图默认撤销 PUBLIC 访问；仅向指定监控角色授予函数执行和视图读取权限。单独执行 `LOAD` 不会初始化共享内存，视图会返回空集并发出警告。

### 配置与限制

`pg_query_progress.track` 由超级用户控制。会话设置包括 `pg_query_progress.bar`、`pg_query_progress.bar_ansi`、`pg_query_progress.bar_interval_ms`、`pg_query_progress.bar_min_query_ms` 与 `pg_query_progress.max_query_len`。进度条使用 NOTICE，因此可能被客户端消息过滤隐藏。

`progress` 是 0 到 1 的比例，并非实测完成百分比。行数估计失准和阻塞式计划会影响进度与剩余时间估计。DML、DDL 和工具语句不被跟踪。每个会话只跟踪一个 portal，执行另一查询可能将未关闭游标标为终止。崩溃或重启后，共享内存记录消失。并行工作进程参与估算，但阻塞的并行阶段可能很少更新中间进度。
