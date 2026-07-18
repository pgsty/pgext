## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/plv8/)
- [官方项目或服务商页面](https://plv8.github.io/)

`plls` — 基于 PL/v8 运行时编译并执行 LiveScript 的 PostgreSQL 过程语言

已复核目录快照记录的版本为 `3.1.10`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "plls";
SELECT extversion
FROM pg_extension
WHERE extname = 'plls';
```

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
