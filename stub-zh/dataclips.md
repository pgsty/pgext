## 用法

来源：

- [官方扩展控制文件](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/dataclips.control)
- [官方上游文档](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/docs/dataclip.md)
- [官方上游 README](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/README.md)

`dataclips` — 已归档且未完成的 C 扩展原型，计划通过函数和 FDW 查询 Heroku Dataclips。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "dataclips";
SELECT extversion
FROM pg_extension
WHERE extname = 'dataclips';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
