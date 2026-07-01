


## 用法

来源：[official README](https://github.com/supabase/pg_net)、[v0.20.3 release notes](https://github.com/supabase/pg_net/releases/tag/v0.20.3)、[local package metadata](../db/extension.csv)。

`pg_net` 可以从 SQL 中排队发起异步 HTTP 和 HTTPS 请求。它会创建 `net` 模式，将待处理工作存入 `net.http_request_queue`，并将响应存入 `net._http_response`。后台工作进程使用 `libcurl` 处理队列中的请求。

该扩展要求配置 `shared_preload_libraries = 'pg_net'`，并需要 `libcurl >= 7.83`。

### GET 请求

```sql
CREATE EXTENSION pg_net;

SELECT net.http_get(
  'https://postman-echo.com/get',
  params := '{"foo": "bar"}'::jsonb,
  headers := '{"API-KEY": "<key>"}'::jsonb,
  timeout_milliseconds := 1000
) AS request_id;
```

`net.http_get(url, params, headers, timeout_milliseconds)` 返回一个 `bigint` 请求 ID。

### POST 请求

```sql
SELECT net.http_post(
  'https://postman-echo.com/post',
  body := '{"key": "value"}'::jsonb,
  headers := '{"Content-Type": "application/json"}'::jsonb,
  timeout_milliseconds := 1000
) AS request_id;
```

将一行表数据作为 JSON 发送：

```sql
WITH selected_row AS (
  SELECT * FROM my_table LIMIT 1
)
SELECT net.http_post(
  'https://api.example.com/data',
  to_jsonb(selected_row.*)
) AS request_id
FROM selected_row;
```

### DELETE 请求

```sql
SELECT net.http_delete(
  'https://api.example.com/resource/42',
  timeout_milliseconds := 2000
) AS request_id;
```

`net.http_delete(url, params, headers, timeout_milliseconds)` 是 `SECURITY DEFINER`，并返回一个 `bigint` 请求 ID。

### 查看响应

```sql
SELECT id, status_code, content, error_msg, created
FROM net._http_response
ORDER BY created DESC;
```

可以通过 `status_code` 和 `error_msg` 识别失败请求。响应表不会保留所有原始请求参数；如果需要实现重试工作流，请单独保存请求元数据。

### 配置

```sql
SHOW pg_net.batch_size;
SHOW pg_net.ttl;
SHOW pg_net.database_name;
SHOW pg_net.username;
```

- `pg_net.batch_size` 默认值为 `200`，限制工作进程每轮处理的队列请求数量。
- `pg_net.ttl` 默认值为 `6 hours`，控制响应保留时间。
- `pg_net.database_name` 默认值为 `postgres`，选择工作进程运行所在数据库。
- `pg_net.username` 默认值为 NULL，选择工作进程连接用户；NULL 表示使用引导用户。

可以在 `postgresql.conf` 中修改设置，也可以使用 `ALTER SYSTEM`：

```sql
ALTER SYSTEM SET pg_net.ttl TO '1 hour';
ALTER SYSTEM SET pg_net.batch_size TO 500;
SELECT pg_reload_conf();
```

修改 `pg_net.database_name` 或 `pg_net.username` 后，需要重启工作进程：

```sql
SELECT net.worker_restart();
```

### 注意事项

- Pigsty 元数据记录 PostgreSQL 14-18 的 `pg_net` 版本为 0.20.3，但本地包备注说明 0.20.3 只在 `d12`、`d13`、`el10`、`u24` 和 `u26` 上可用；由于旧版 `libcurl`，`el8`、`el9` 和 `u22` 仍停留在 0.9.2。
- 上游文档声明兼容 PostgreSQL 12+，但此目录行面向 PostgreSQL 14-18 打包。
- `pg_net` 通过 `pg_net.database_name` 在每个集群中只支持一个数据库。
- v0.20.3 是工作进程/维护发布：它会刷新 pgstat 计数器以便 autovacuum 可见，并向 `pg_stat_activity` 报告工作进程活动；没有记录新的 SQL 请求 API。
