## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/sr_plan.control)
- [官方上游文档](https://github.com/postgrespro/sr_plan/blob/19a6a425fc91bdfc5482650db9747bead67efd40/README.md)

`sr_plan` — 基于规划器钩子的扩展，可记录、展示并复用执行计划，包括参数化计划模板。

已复核目录快照记录的版本为 `1.2`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "sr_plan";
SELECT extversion
FROM pg_extension
WHERE extname = 'sr_plan';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
