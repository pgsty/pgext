## 用法

来源：

- [官方上游 README](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/README.md)

`pg_idna` — 为 PostgreSQL 提供 WHATWG URL 标准的国际化域名转换函数

已复核目录快照记录的版本为 `0.1.1-docs`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_idna";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_idna';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
