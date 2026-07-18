## 用法

来源：

- [官方上游文档](https://github.com/a5221985/PhaseZero_Scripts/blob/443e3a5e59d2816b7b8d7ee0129c8685f3f11d8a/README.md)

`scripts` — 以 PostgreSQL 扩展形式打包的 PhaseZero 售后业务应用表结构。

已复核目录快照记录的版本为 `1.4`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "scripts";
SELECT extversion
FROM pg_extension
WHERE extname = 'scripts';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
