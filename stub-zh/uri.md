

## 用法

> [uri: 支持验证和组件提取的 URI 数据类型](https://github.com/petere/pguri)

`uri` 扩展提供了一种用于存储 URI 的数据类型，支持 RFC 3986 语法验证、组件提取函数和人性化排序。

```sql
CREATE EXTENSION uri;

CREATE TABLE links (
    id   int PRIMARY KEY,
    link uri
);

INSERT INTO links VALUES (1, 'https://github.com/petere/pguri');
```

### 组件提取函数

| 函数 | 返回类型 | 描述 |
|------|---------|------|
| `uri_scheme(uri)` | `text` | 协议（http、ftp、mailto） |
| `uri_userinfo(uri)` | `text` | 用户信息；不存在时返回 NULL |
| `uri_host(uri)` | `text` | 主机名或 IP 地址 |
| `uri_host_inet(uri)` | `inet` | IP 主机转为 inet 类型；非 IP 时返回 NULL |
| `uri_port(uri)` | `integer` | 端口号；未指定时返回 NULL |
| `uri_path(uri)` | `text` | 路径组件（不会为 NULL） |
| `uri_path_array(uri)` | `text[]` | 按 `/` 分割的路径 |
| `uri_query(uri)` | `text` | 查询字符串；不存在时返回 NULL |
| `uri_fragment(uri)` | `text` | 片段标识；不存在时返回 NULL |

### 工具函数

```sql
-- 按 RFC 3986 标准化 URI
SELECT uri_normalize('HTTP://Example.COM/foo/../bar');

-- 百分号编码文本
SELECT uri_escape('hello world', true, false);  -- hello+world

-- 解码百分号编码文本
SELECT uri_unescape('hello+world', true, false);  -- hello world
```

### 示例

```sql
SELECT uri_scheme(link), uri_host(link), uri_path(link)
FROM links
WHERE uri_host(link) = 'github.com';
```
