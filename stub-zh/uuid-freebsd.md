## 用法

来源：

- [已归档的上游仓库](https://github.com/RhodiumToad/uuid-freebsd)
- [上游 README](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/README.md)
- [2.0 版本 SQL 定义](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd--2.0.sql)
- [上游构建元数据](https://github.com/RhodiumToad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/Makefile)
- [当前 PostgreSQL UUID 函数](https://www.postgresql.org/docs/current/functions-uuid.html)

`uuid-freebsd` 是基于操作系统 UUID 实现的 FreeBSD 专用 UUID 生成器。扩展名包含连字符，因此必须在 SQL 中加引号：

```sql
CREATE EXTENSION "uuid-freebsd";

SELECT uuid_generate_v4();
SELECT uuid_generate_v5(uuid_ns_url(), 'https://example.com/resource/42');
```

`2.0` 版本提供 nil UUID、DNS/URL/OID/X.500 namespace 常量、随机 UUID v4、基于时间的 v1 与 v1mc，以及基于名称的 v3 与 v5 函数。它要求 FreeBSD 并链接 `libmd`，不是面向 Linux 或其他平台的可移植实现。

### 已淘汰状态

维护者将项目标记为 obsolete，并于 2023 年归档仓库。最后一次源码变更早于现代 PostgreSQL 版本。应优先使用当前 PostgreSQL 版本自带的 UUID 函数；需要其他算法时，可使用仍在维护的 `uuid-ossp` 包。`uuid-freebsd` 只适合兼容的 FreeBSD 遗留部署。
