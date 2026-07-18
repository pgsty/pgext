## 用法

来源：

- [官方扩展控制文件](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/safecast.control)
- [官方上游文档](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/README.md)

`safecast` — 纯 SQL 安全类型转换辅助函数；部分无效文本转数值时返回 NULL。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "safecast";
SELECT extversion
FROM pg_extension
WHERE extname = 'safecast';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
