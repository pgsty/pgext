## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/pg_compat19.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/pg_compat19--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_compat19/meson.build)

`pg_compat19` `1.0` 提供面向 AlohaDB PostgreSQL 18 内核的 SQL 兼容函数。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_compat19;
SELECT random('2026-01-01'::date, '2026-12-31'::date);
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

该模块由超级用户安装到 pg_catalog，提供 error_on_null、元组标识辅助函数、日期时间随机范围、范围相减以及 UUID/bytea 转换。它依赖匹配的 AlohaDB 内部函数。升级到基于 PostgreSQL 19 的 AlohaDB 内核前需移除此兼容扩展，因为新内核已内置相应对象；移除前应检查依赖对象。

此模块对应上游旧名 `alohadb_pg19compat`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
