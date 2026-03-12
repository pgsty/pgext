

## 用法

> [pg_net: 在 SQL 中发起异步 HTTP/HTTPS 请求](https://github.com/supabase/pg_net)

该扩展需要在 `postgresql.conf` 中设置 `shared_preload_libraries = 'pg_net'`。

### GET 请求

```sql
SELECT net.http_get('https://postman-echo.com/get?foo=bar') AS request_id;
```

使用 URL 编码参数和请求头：

```sql
SELECT net.http_get(
  'https://postman-echo.com/get',
  params := '{"foo": "bar"}'::JSONB,
  headers := '{"API-KEY": "<key>"}'::JSONB
) AS request_id;
```

### POST 请求

```sql
SELECT net.http_post(
    'https://postman-echo.com/post',
    '{"key": "value"}'::JSONB,
    headers := '{"Content-Type": "application/json"}'::JSONB
) AS request_id;
```

将表行作为请求体发送：

```sql
WITH row AS (SELECT * FROM my_table LIMIT 1)
SELECT net.http_post(
    'https://api.example.com/data',
    to_jsonb(row.*)
) AS request_id
FROM row;
```

### DELETE 请求

```sql
SELECT net.http_delete('https://api.example.com/resource/42') AS request_id;
```

### 查看响应

```sql
SELECT * FROM net._http_response;
```

### 配置

```sql
SHOW pg_net.batch_size;       -- 默认值：200，每个处理周期的最大行数
SHOW pg_net.ttl;              -- 默认值：6 小时，响应保留时间
SHOW pg_net.database_name;    -- 默认值：'postgres'
```

修改设置：

```sql
ALTER SYSTEM SET pg_net.ttl TO '1 hour';
ALTER SYSTEM SET pg_net.batch_size TO 500;
SELECT pg_reload_conf();
```
