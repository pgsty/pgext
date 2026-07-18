## 用法

来源：

- [官方扩展控制文件](https://github.com/asotolongo/vacuum_utils/blob/8de352f1960ad0d32c9d0d8cf24fa7417ef6c057/vacuum_utils.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/vacuum_utils/)

`vacuum_utils` — 提供检查 VACUUM/ANALYZE 阈值与活动状态、并执行表维护操作的 SQL 函数。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "vacuum_utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'vacuum_utils';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
