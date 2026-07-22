## 用法

来源：

- [官方 README](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/README.md)
- [官方扩展控制文件](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url.control)
- [官方扩展 SQL](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url--1.0.sql)

`url`（项目名为 `pg_url`）是一个纯 SQL/PL/pgSQL 库，用于解析 URL、对值进行百分号编码，并把查询字符串作为结构化 PostgreSQL 值处理。它不会发起 HTTP 请求，也不会解析或验证远端主机。

### 核心流程

```sql
CREATE EXTENSION url;

SELECT url_encode('hello world/你好');
SELECT url_decode('hello%20world%2F%E4%BD%A0%E5%A5%BD');

SELECT url('https://alice@example.com:8443/docs?q=postgres&q=sql#usage');

SELECT qskv('q=postgres&q=sql&empty') -> 'q';
SELECT qskv('q=postgres&q=sql&empty') ->> 'q';
```

复合类型 `url` 包含 `scheme`、`authority`、`domain`、`port`、`path`、`querystring` 与 `fragment` 字段。可用 `url(text)` 或类型转换解析文本，再转回文本重建。`qskv` 保存 `kventry` 数组：`->` 返回某个键的第一个值，`->>` 返回全部匹配值组成的 `text[]`。`unnest(qskv)` 输出键值行。

### 解析边界

该实现是面向 PostgreSQL 9.4 时代语法的正则表达式解析器，不是完整的标准或安全验证器。采用前应测试国际化域名、IPv6 字面量、特殊用户信息、无效百分号转义、重复/空查询键、`+` 处理与往返转换。解析得到的 `domain` 或 `path` 仍是不可信输入；应用必须另行限制允许的协议与目标，以防 SSRF 或开放重定向。扩展不需要网络预加载、重启或服务器文件系统访问。
