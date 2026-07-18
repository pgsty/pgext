## 用法

来源：

- [上游兼容性说明](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/README.md)
- [函数实现](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/src/jsonb_set_lax.c)
- [SQL 声明](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/sql/jsonb_set_lax.sql)

`jsonb_set_lax` 把 PostgreSQL 13 的同名函数回移到 PostgreSQL 9.5 至 12。它可改变 SQL NULL 替换值的处理方式，可选 `use_json_null`、`delete_key`、`return_target` 或 `raise_exception`。

```sql
CREATE EXTENSION jsonb_set_lax;
SELECT jsonb_set_lax('{"a":1}'::jsonb, '{a}', NULL, true, 'use_json_null');
SELECT jsonb_set_lax('{"a":1}'::jsonb, '{a}', NULL, true, 'delete_key');
```

该扩展安装到 `pg_catalog` 且不可迁移 schema。源码会主动拒绝 PostgreSQL 13 及以上版本，此时应使用内核实现。迁移到较新服务器或从其迁回时，应区分 SQL NULL 与 JSON `null`，校验处理模式字符串，并测试缺失路径和数组行为。
