


## 用法

> [logerrors: 收集日志消息统计信息](https://github.com/munakoiso/logerrors)

`logerrors` 收集 PostgreSQL 日志文件中 WARNING、ERROR 和 FATAL 消息的统计信息，无需解析日志即可轻松监控错误率。

```sql
CREATE EXTENSION logerrors;
```

### 配置参数

| 参数 | 默认值 | 描述 |
|------|--------|------|
| `logerrors.interval` | `5000`（5秒） | 写入统计到缓冲区的间隔（毫秒，最大 60 秒） |
| `logerrors.intervals_count` | `120` | 缓冲区中保留的间隔数（最大 360） |
| `logerrors.excluded_errcodes` | （空） | 要排除的错误代码，逗号分隔 |

### 查询错误统计

```sql
SELECT * FROM pg_log_errors_stats();
 time_interval |  type   |       message        | count | username | database | sqlstate
---------------+---------+----------------------+-------+----------+----------+----------
               | WARNING | TOTAL                |     0 |          |          |
               | ERROR   | TOTAL                |     1 |          |          |
               | FATAL   | TOTAL                |     0 |          |          |
             5 | ERROR   | ERRCODE_SYNTAX_ERROR |     1 | postgres | postgres | 42601
           600 | ERROR   | ERRCODE_SYNTAX_ERROR |     1 | postgres | postgres | 42601
```

### 慢日志统计

```sql
SELECT * FROM pg_slow_log_stats();
 slow_count |         reset_time
------------+----------------------------
          1 | 2020-06-13 00:19:31.084923
```

### 重置统计

```sql
SELECT pg_log_errors_reset();
```
