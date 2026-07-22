## 用法

来源：

- [1.1 版本 README](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/readme.md)
- [扩展 control 文件](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/pg_https.control)
- [1.1 版本 SQL 接口](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/sql/pg_https--1.1.sql)
- [SQL 入口与 GUC 定义](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/src/pg_https.c)
- [libcurl 请求、重试与响应实现](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/src/https_core.c)

`pg_https` 1.1 是 PostgreSQL 的同步原生 HTTP/HTTPS 客户端。它在 `requests` 模式中封装 libcurl，返回状态码、响应头、响应体、耗时与 curl 错误详情；对外部产生的效果不属于数据库事务。

### 预加载与配置

上游指示管理员将 `pg_https` 加入 `shared_preload_libraries`，重启 PostgreSQL 后再创建扩展：

```conf
shared_preload_libraries = 'pg_https'
```

```sql
CREATE EXTENSION pg_https;

SET pg_https.timeout = 10;
SET pg_https.connect_timeout = 5;
SET pg_https.tls_version = 13;
SET pg_https.verify_peer = true;
SET pg_https.max_response_size = 10485760;
SET pg_https.default_headers = '{"User-Agent":"pg_https/1.1"}';
```

其他可由用户设置的 GUC 包括 `pg_https.ca_file`、`pg_https.tcp_keepalive`、`pg_https.http_version` 和 `pg_https.connection_reuse`。应保持证书与主机名验证开启；自定义 CA 文件必须能被 PostgreSQL 操作系统账户读取。

### 核心工作流

```sql
SELECT status, headers, body, curl_error_code, error_message
FROM requests.get_request(
  'https://api.example.invalid/health',
  '{"Accept":"application/json"}'::jsonb,
  NULL,
  5
);

SELECT status, body, duration_ms
FROM requests.rest_request(
  'POST',
  'https://api.example.invalid/events',
  '{"Content-Type":"application/json"}'::jsonb,
  '{"event":"refresh"}',
  5,
  NULL,
  NULL,
  0,
  300,
  2.0,
  0
);
```

`http_response` 结果的字段为 `status`、`headers`、`body`、`duration_ms`、`response_bytes`、`curl_error_code` 和 `error_message`。响应头名会转换为小写并以 JSON 数组保存，完整响应头块也会保留在 raw 字段中。

### 主要对象与重试语义

- `requests.rest_request` 接收方法、URL、JSONB 请求头、请求体、超时覆盖值、可选 Basic Auth 凭据、重试参数与 `cancel_mode`。
- 便捷包装函数包括 `requests.get_req`、`requests.get_request`、`requests.post_req`、`requests.post_request`、`requests.put_req`、`requests.put_request`、`requests.delete_req` 和 `requests.delete_request`。
- `requests.auth_bearer` 构造 Authorization 请求头对象。
- `cancel_mode` 为 0 时，取消查询会中止传输；为 1 时，进度回调会等待请求完成。

重试仅适用于幂等的 `GET`、`HEAD`、`PUT`、`DELETE` 和 `OPTIONS` 请求。它会在选定的短暂 libcurl 失败或 HTTP `408`、`429`、`5xx` 响应时发生。即使重试次数非零，`POST` 也不会重试。请求在调用后端中同步执行，指数退避会延长该后端的占用时间。

### 安全与事务边界

该扩展没有目标白名单。调用者可以选择任意 URL、跟随重定向、发送写方法、设置 Basic 或 bearer 凭据，并改变用户可设置的 TLS 行为。这会形成通向本机、云元数据端点、内部服务和公网的 SSRF 与数据泄露边界。应撤销 `PUBLIC` 对请求函数的 `EXECUTE`，仅授予范围很小的角色，并在 PostgreSQL 外部强制出站网络策略。

生产中绝不应禁用 `pg_https.verify_peer`，也不应把密钥放在 URL、查询文本、默认请求头或临时 SQL 参数中。实现会记录每次 curl 尝试，对慢请求或失败请求还会记录方法和 URL；普通 PostgreSQL 语句日志还可能暴露更多查询文本。

即使外部 HTTP 请求已成功，外层 SQL 事务仍可能随后回滚。应在远程 API 中设计幂等性与对账机制。`pg_https.max_response_size` 会限制响应缓冲，但不会使不可信端点变得安全，也不会转为异步调用。上游记录 1.1 在 PostgreSQL 17 上测试，并依赖已安装的 libcurl 构建；使用前应在准确服务器上验证取消、重定向、DNS、TLS、重试、内存限制和失败行为。
