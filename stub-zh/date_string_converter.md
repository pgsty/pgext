## 用法

来源：

- [官方扩展控制文件](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/date_string_converter.control)
- [官方上游文档](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/README.md)
- [官方 Rust 包清单](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/Cargo.toml)

`date_string_converter` — 用于日期字符串转换的 pgrx 扩展骨架；当前源码仅暴露 hello_date_string_converter()。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "date_string_converter";
SELECT extversion
FROM pg_extension
WHERE extname = 'date_string_converter';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
