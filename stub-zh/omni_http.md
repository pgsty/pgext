


## 用法

> [omni_http: 基础 HTTP 类型](https://docs.omnigres.org/omni_httpc/reference/)

`omni_http` 扩展提供 `omni_httpc` 和 `omni_httpd` 使用的基础 HTTP 类型。

### 主要类型

- **`omni_http.http_method`** -- HTTP 方法枚举（GET、POST、PUT、DELETE 等）
- **`omni_http.http_headers`** -- HTTP 头部数组类型。头部名称为 null 时不创建头部；值为 null 时序列化为空字符串。
- **`omni_http.http_request`** -- 表示 HTTP 请求的复合类型
- **`omni_http.http_response`** -- 表示 HTTP 响应的复合类型

这些类型构成了 HTTP 客户端（`omni_httpc`）和 HTTP 服务器（`omni_httpd`）扩展之间的共享接口。
