


## 用法

来源：[README](https://github.com/pramsey/pgsql-http/blob/v1.7.1/README.md)、[v1.7.1 release](https://github.com/pramsey/pgsql-http/releases/tag/v1.7.1)

`http` 允许 SQL 代码通过 libcurl 发起 HTTP 请求。它适合受控的集成点，例如触发器通知外部服务、SQL job 拉取小型远程 payload，或数据库侧 webhook 调用。

```sql
CREATE EXTENSION http;
```

### 请求与响应类型

每个请求使用 `http_request`，并返回 `http_response`：

```text
http_request(method http_method, uri varchar, headers http_header[], content_type varchar, content varchar)
http_response(status integer, content_type varchar, headers http_header[], content varchar)
```

便捷包装函数最终调用同一个底层 `http(http_request)` 函数：

- `http_get(uri varchar)`
- `http_get(uri varchar, data jsonb)`
- `http_post(uri varchar, content varchar, content_type varchar)`
- `http_post(uri varchar, data jsonb)`
- `http_put(uri varchar, content varchar, content_type varchar)`
- `http_patch(uri varchar, content varchar, content_type varchar)`
- `http_delete(uri varchar)`
- `http_head(uri varchar)`

### 示例

```sql
SELECT status, content_type, content
FROM http_get('https://httpbun.com/ip');

SELECT content::json->'headers'->>'Authorization'
FROM http((
  'GET',
  'https://httpbun.com/headers',
  http_headers('Authorization', 'Bearer token'),
  NULL,
  NULL
)::http_request);

SELECT status, content::json->'form' AS form
FROM http_post(
  'https://httpbun.com/post',
  jsonb_build_object('myvar', 'myval', 'foo', 'bar')
);

SELECT status, content_type, content::json->>'data' AS data
FROM http_put('https://httpbun.com/put', 'some text', 'text/plain');
```

通过展开 `headers` 数组检查响应头：

```sql
SELECT (unnest(headers)).*
FROM http_get('https://httpbun.com/');
```

### 二进制内容

README 提醒，`varchar::bytea` 对二进制响应体不安全，因为它会在零值字节处停止。读取响应内容时使用 `text_to_bytea(content)`，发送二进制请求体时使用 `bytea_to_text(bytea)`。

```sql
WITH http AS (
  SELECT * FROM http_get('https://httpbingo.org/image/png')
)
SELECT content_type, length(text_to_bytea(content)) AS bytes
FROM http;
```

### 超时与版本说明

`pg_http` 1.7.1 是兼容性和文档版本：增加 timeout examples，增加 PostgreSQL 17 wait-event hooks，并包含 PostgreSQL 19 support fixes。用户可见 SQL API 仍是上面的 README surface。
