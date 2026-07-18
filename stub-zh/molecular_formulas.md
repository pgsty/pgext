## 用法

来源：

- [官方扩展控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/molecular_formulas.control)
- [官方上游文档](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/README.md)
- [官方 Rust 包清单](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/Cargo.toml)

`molecular_formulas` — 分子式解析器和 PostgreSQL 强类型数据类型。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "molecular_formulas";
SELECT extversion
FROM pg_extension
WHERE extname = 'molecular_formulas';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
