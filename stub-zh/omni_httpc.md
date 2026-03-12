


## 用法

> [omni_httpc: HTTP 客户端](https://docs.omnigres.org/omni_httpc/reference/)

`omni_httpc` 扩展提供 HTTP/1、HTTP/2 和 HTTP/3（实验性）客户端功能。

### 准备和执行请求

```sql
SELECT version >> 8 AS http_version, status, headers, convert_from(body, 'utf-8')
FROM omni_httpc.http_execute(
    omni_httpc.http_request('https://example.com'),
    omni_httpc.http_request('https://example.org')
);
```

**`http_request(url, method, headers, body)`** -- 准备请求。方法默认为 GET。

**`http_execute(VARIADIC requests)`** -- 执行一个或多个请求。返回 `version`、`status`、`headers`、`body` 和 `error` 列。

### 执行选项

```sql
SELECT * FROM omni_httpc.http_execute_with_options(
    omni_httpc.http_execute_options(http2_ratio => 100),
    omni_httpc.http_request('https://example.com')
);
```

| 选项 | 类型 | 默认值 | 描述 |
|:-----|:-----|:-------|:-----|
| `http2_ratio` | smallint | 0 | 0-100，HTTP/2 使用百分比 |
| `http3_ratio` | smallint | 0 | 0-100，HTTP/3 使用百分比 |
| `force_cleartext_http2` | bool | false | 使用 h2c |
| `first_byte_timeout` | int | 5000 | 毫秒 |
| `timeout` | int | 5000 | 毫秒 |
| `follow_redirects` | bool | true | 跟随 HTTP 重定向 |
| `cacerts` | text[] | null | PEM 编码的 CA 证书 |
| `clientcert` | client_certificate | null | PEM 编码的客户端证书 |

`http2_ratio` 和 `http3_ratio` 之和不得超过 100。

### 连接池

```sql
SELECT * FROM omni_httpc.http_connections;
-- 返回: http_protocol, url
```
