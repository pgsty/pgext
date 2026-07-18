## 用法

来源：

- [官方扩展控制文件](https://github.com/boringSQL/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/pg_viewctl.control)
- [官方上游文档](https://github.com/boringSQL/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/README.md)

`pg_viewctl` — pg_viewctl：用于视图依赖管理和安全模式演进的 PostgreSQL 工具扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_viewctl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_viewctl';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
