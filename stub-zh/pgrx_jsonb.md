## 用法

来源：

- [官方扩展控制文件](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/pgrx_jsonb.control)
- [官方上游文档](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/README.md)
- [官方 Rust 包清单](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/Cargo.toml)

`pgrx_jsonb` — 用于遍历 PostgreSQL JSONB 内部结构并重建 JSON 值的 pgrx 概念验证函数

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgrx_jsonb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgrx_jsonb';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
