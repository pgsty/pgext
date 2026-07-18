## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_ethiopian_calendar/)

`pg_ethiopian_calendar` — 在公历日期与埃塞俄比亚历日期之间进行双向转换

已复核目录快照记录的版本为 `1.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ethiopian_calendar";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ethiopian_calendar';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
