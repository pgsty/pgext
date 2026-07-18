## 用法

来源：

- [官方扩展控制文件](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.control)
- [官方上游文档](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/doc/utinyint.md)
- [官方上游 README](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/README.md)

`utinyint` — 为 PostgreSQL 提供 1 字节无符号整数类型，以及与 smallint、integer 和 boolean 的类型转换。

已复核目录快照记录的版本为 `0.1.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "utinyint";
SELECT extversion
FROM pg_extension
WHERE extname = 'utinyint';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
