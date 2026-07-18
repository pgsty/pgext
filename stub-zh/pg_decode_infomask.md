## 用法

来源：

- [官方扩展控制文件](https://github.com/frost242/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/pg_decode_infomask.control)
- [官方上游文档](https://github.com/frost242/pg_decode_infomask/blob/3cac9efcb408743601c9eb98b24704854210e027/README.md)

`pg_decode_infomask` — 面向 PostgreSQL 9.6 的 C 扩展，将堆元组 infomask 与 infomask2 提示位解码为可查询字段。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_decode_infomask";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_decode_infomask';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
