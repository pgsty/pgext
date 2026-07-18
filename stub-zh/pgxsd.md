## 用法

来源：

- [官方扩展控制文件](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.control)

`pgxsd` — pgxsd 用数据库内保存的 XSD 模式校验 PostgreSQL XML 值。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgxsd";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgxsd';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
