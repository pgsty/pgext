## 用法

来源：

- [pgsql-http v1.7.2 README](https://github.com/pramsey/pgsql-http/blob/v1.7.2/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-http/blob/v1.7.2/http.control)
- [v1.7.1 to v1.7.2 comparison](https://github.com/pramsey/pgsql-http/compare/v1.7.1...v1.7.2)

`http` 允许 PostgreSQL 通过 libcurl 发出同步 HTTP 请求。它适用于受控集成和管理调用，但后端会在 SQL 语句和事务内部等待远程服务的响应。限制可以调用它的用户、设置较短的超时时间，并确保不可信输入不会选择任意 URL。

### 核心工作流程

```sql
CREATE EXTENSION http;

SELECT status, content_type, content
FROM http_get('https://httpbingo.org/get');
```

发送 JSON 并检查响应：

```sql
SELECT status, content::jsonb
FROM http_post(
  'https://httpbingo.org/post',
  '{"event":"invoice.paid"}',
  'application/json'
);
```

通用入口点接受完整的请求：

```sql
SELECT (http((
  'GET',
  'https://httpbingo.org/headers',
  http_headers('Authorization', 'Bearer example'),
  NULL,
  NULL
)::http_request)).status;
```

### 重要对象

- `http_request` 包含 `method`, `uri`, `headers`, `content_type`, 和 `content`。
- `http_response` 包含 `status`, `content_type`, `headers`, 和 `content`。
- `http_header`, `http_header(...)`, 和 `http_headers(...)` 构建请求头；`unnest(response.headers)` 将响应头暴露为行。
- `http(...)` 执行完整的 `http_request`。
- `http_get`, `http_post`, `http_put`, `http_patch`, `http_delete`, 和 `http_head` 是方便的包装器。
- `urlencode(text)` 和 `urlencode(jsonb)` 编码查询数据。
- `http_set_curlopt`, `http_list_curlopt`, 和 `http_reset_curlopt` 管理支持的会话级 libcurl 设置。

### 超时、连接和安全性

默认情况下，每个请求使用一个新的连接。在测量后端生命周期和远程服务器行为之后再启用持久连接：

```sql
SET http.curlopt_timeout_ms = 1000;
SET http.curlopt_connecttimeout_ms = 250;
SET http.curlopt_tcp_keepalive = 1;
```

默认请求超时为五秒。超时将引发 SQL 错误，因此调用者必须处理事务回滚。触发器或长时间事务中的网络延迟可能会持有锁并耗尽数据库连接；优先使用外部工作进程进行持久异步传递。

保持 TLS 验证启用、保护包含凭据的 curl 设置、在使用前验证响应状态和内容，并在网络层和 SQL 权限层限制出站目的地。版本 1.7.2 包含相对于 1.7.1 的构建、测试和 curl 选项常量维护；它没有引入实质性 SQL API 变更。控制文件仍然声明 SQL 扩展版本为 1.7。
