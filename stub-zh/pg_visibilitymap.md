## 用法

来源：

- [上游 README](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/README.md)
- [扩展 control 文件](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/pg_visibilitymap.control)
- [SQL 安装脚本](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/pg_visibilitymap--1.0.sql)
- [受支持的 pg_visibility 文档](https://www.postgresql.org/docs/current/pgvisibility.html)

`pg_visibilitymap` `1.0` 版检查表或物化视图可见性映射中的全可见位。它通过 `pg_visibilitymap(regclass)` 返回全部映射项，或用 `pg_is_all_visible(regclass,bigint)` 检查单个堆块。

### 示例

```sql
CREATE EXTENSION pg_visibilitymap;
SELECT * FROM pg_visibilitymap('public.orders'::regclass);
SELECT pg_is_all_visible('public.orders'::regclass, 0);
```

安装脚本已从 `PUBLIC` 撤销两个函数，只应向可信诊断角色授权。可见性映射位是维护元数据，不能独立证明每个元组对特定快照可见。这个停留在 2015 年、依赖旧版 PostgreSQL 内部接口的扩展已弃用；当前系统应优先使用 PostgreSQL 仍在维护的 `pg_visibility` 扩展。
