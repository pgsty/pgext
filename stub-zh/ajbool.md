## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/ajbool.control)
- [官方上游文档](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/ajbool/)

`ajbool` — 带有伪 NULL 未知状态的三值布尔类型，可用于主键。

已复核目录快照记录的版本为 `0.0.3`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ajbool";
SELECT extversion
FROM pg_extension
WHERE extname = 'ajbool';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
