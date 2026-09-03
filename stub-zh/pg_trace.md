## 用法

来源：

- [官方文档](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension/blob/1cbd9d46a401b22070549543e493f19845eb3ad2/README.md)
- [扩展控制文件](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension/blob/1cbd9d46a401b22070549543e493f19845eb3ad2/pg_trace.control)
- [官方仓库](https://github.com/huraira213/pg_trace---PostgreSQL-Query-Tracing-Extension)

`pg_trace` 提供共享内存缓冲、持久化与保留控制的查询执行追踪。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_trace`：

```sql
CREATE EXTENSION pg_trace;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- All commands in one session
SELECT pg_trace_start();
SELECT 1+1;
SELECT pg_trace_stop();
SELECT * FROM pg_trace_queries;
SELECT pg_trace_flush();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_trace_start` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_trace_stop` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_trace_flush` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_trace_log` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pg_trace_queries` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_trace_cleanup_old` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_trace_clear` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_trace_hourly_stats` | VIEW | 扩展创建的检查或查询视图。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 16, 17, 18；不要推断未列出的主版本。
