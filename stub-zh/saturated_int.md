## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/saturated_int.control)
- [官方上游文档](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/README.md)

`saturated_int` — 提供饱和算术语义的 PostgreSQL 整数类型扩展，溢出时钳制到 int4 边界。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "saturated_int";
SELECT extversion
FROM pg_extension
WHERE extname = 'saturated_int';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
