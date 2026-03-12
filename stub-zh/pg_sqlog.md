


## 用法

> [pg_sqlog: 通过 SQL 访问 PostgreSQL CSV 日志](https://github.com/kouber/pg_sqlog)

pg_sqlog 提供一个 SQL 接口，使用 `file_fdw` 支持的外部表来查询 PostgreSQL CSV 日志文件。

### 查询日志

```sql
-- 当天的日志
SELECT * FROM sqlog.log();

-- 指定日期的日志
SELECT * FROM sqlog.log('yesterday');
SELECT * FROM sqlog.log('2024-01-15');

-- 错误摘要
SELECT error_severity, COUNT(*) FROM sqlog.log() GROUP BY 1;
```

### 慢查询分析

```sql
-- 最慢的 3 种查询模式
SELECT
  AVG(sqlog.duration(message)) AS avg_duration,
  COUNT(*) AS count,
  sqlog.preparable_query(message) AS query_pattern
FROM sqlog.log()
WHERE message ~ '^duration'
GROUP BY 3
ORDER BY 1 DESC
LIMIT 3;

-- 带时间信息的查询摘要
SELECT
  log_time::time,
  sqlog.duration(message),
  sqlog.summary(message)
FROM sqlog.log('yesterday')
WHERE message ~ '^duration';
```

### 函数

| 函数 | 描述 |
|----------|-------------|
| `sqlog.log([timestamp])` | 返回指定日期的日志内容 |
| `sqlog.set_date([timestamp])` | 设置 `sqlog.log` 表查询的日期 |
| `sqlog.duration(text)` | 从消息中提取查询持续时间（毫秒） |
| `sqlog.preparable_query(text)` | 将参数替换为 `?` 以便分组 |
| `sqlog.summary(text, int, int)` | 去除元数据，显示前/后 N 个字符 |
| `sqlog.temporary_file_size(text)` | 从消息中提取临时文件大小 |
| `sqlog.autovacuum([timestamp])` | 指定日期的自动清理报告 |
| `sqlog.autoanalyze([timestamp])` | 指定日期的自动分析报告 |

### 自动清理报告

```sql
SELECT * FROM sqlog.autovacuum() LIMIT 5;
SELECT * FROM sqlog.autoanalyze() LIMIT 5;
```

### 前置条件

需要在 `postgresql.conf` 中设置：

```
log_destination = 'syslog,csvlog'
log_filename = 'postgresql.%F'
logging_collector = 'on'
log_rotation_age = '1d'
log_rotation_size = 0
log_truncate_on_rotation = 'on'
log_min_duration_statement = 1000
```
