## 用法

来源：

- [Pinned official README](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/README.md)
- [Pinned extension SQL](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/twitter_fdw--1.1.0.sql)

`twitter_fdw` 是面向 Twitter 免认证搜索 API 的历史只读外部数据包装器，把搜索结果映射为预定义的 `public.twitter` 外部表。1.1.0 版使用的端点和响应格式已经退役，因此该扩展只能说明一种旧集成方式，不是今天 X API 的可用接口。

### 历史流程

创建扩展还会一并创建包装器、服务器、当前用户映射和外部表：

```sql
CREATE EXTENSION twitter_fdw;

SELECT from_user, created_at, text
FROM public.twitter
WHERE q = '#postgresql';
```

虚拟列 `q` 与等号一起使用时会被识别，并经过百分号编码后放入远程搜索查询。它是输入参数，而不是每条消息返回的属性。

### 安装的对象

- 函数：`twitter_fdw_handler()` 和 `twitter_fdw_validator(text[], oid)`。
- 外部数据包装器：`twitter_fdw`。
- 外部服务器：`twitter_service`。
- 用户映射：面向执行 `CREATE EXTENSION` 的角色。
- 外部表：`public.twitter`，包含 `id`、`text`、`from_user`、`from_user_id`、`to_user`、`to_user_id`、`iso_language_code`、`source`、`profile_image_url`、`created_at` 及虚拟列 `q`。

### 兼容性与安全边界

固定版本的 C 代码通过 libcurl 调用 `http://search.twitter.com/search.json`，并期待旧版 JSON 搜索字段。这个免认证 HTTP 端点已不再代表受支持的 Twitter/X API。扩展没有 OAuth 选项、API 密钥、bearer token、TLS 端点配置，也没有现代响应对象和限流规则的映射。

项目已经弃用，目标是早期 PostgreSQL FDW API。它依赖系统 libcurl，并内置旧版 libjson。即使改接其他端点，也应在加载原生代码前审计 HTTP 传输、解析器、内存处理和 PostgreSQL API 兼容性。

不要用它访问当前 X，也不要把成功执行 `CREATE EXTENSION` 当作远程查询可用的证明。应使用受维护的 API 客户端，或按照当前认证 API 契约构建新 FDW。它不要求预加载或重启。
