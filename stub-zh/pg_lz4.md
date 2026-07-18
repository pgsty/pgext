## 用法

来源：

- [官方扩展控制文件](https://github.com/zilder/pg_lz4/blob/f3259686848a61c544dbe58df5627ed3d77a2144/pg_lz4.control)

`pg_lz4` — 面向打过自定义压缩方法补丁的 PostgreSQL 开发分支的 LZ4 压缩访问方法。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_lz4";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lz4';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
