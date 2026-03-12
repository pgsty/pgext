


## 用法

> [omni_httpd: HTTP 服务器](https://docs.omnigres.org/omni_httpd/intro/)

`omni_httpd` 扩展是 PostgreSQL 的内嵌 HTTP 服务器，允许在 SQL 中处理 HTTP 请求。

### 定义处理器

```sql
CREATE FUNCTION my_handler(request omni_httpd.http_request)
    RETURNS omni_httpd.http_outcome
    RETURN omni_httpd.http_response(body => request.headers::text);
```

### 注册路由

```sql
INSERT INTO omni_httpd.urlpattern_router (match, handler)
VALUES (omni_httpd.urlpattern('/headers'), 'my_handler'::regproc);
```

### 配置

- **`omni_httpd.http_workers`** -- HTTP 工作进程数（默认为 CPU 数量，受 `max_worker_processes` 限制）
- **`omni_httpd.temp_dir`** -- 套接字临时目录（默认 `/tmp`）

### 控制

```sql
SELECT omni_httpd.stop();   -- 停止服务器
SELECT omni_httpd.start();  -- 启动服务器
-- 传入 immediate => true 可立即执行而无需等待事务提交
```

### 主要类型

- **`omni_httpd.http_request`** -- 请求类型，包含 `path`、`method`、`query_string`、`headers`、`body`
- **`omni_httpd.http_response`** -- 响应构造器
- **`omni_httpd.http_outcome`** -- 处理器返回类型
