## 用法

来源：

- [官方扩展控制文件](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.control)

`file_fixed_length_fdw` — file_fixed_length_fdw：读取固定长度字段文件的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "file_fixed_length_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'file_fixed_length_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
