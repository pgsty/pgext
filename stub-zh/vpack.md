## 用法

来源：

- [官方扩展控制文件](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack.control)
- [官方上游文档](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/README.md)

`vpack` — 为 PostgreSQL 提供 VelocyPack 数据类型、路径查询、运算符、类型转换以及 JSON/BSON 转换支持。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "vpack";
SELECT extversion
FROM pg_extension
WHERE extname = 'vpack';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
