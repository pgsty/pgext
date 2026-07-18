## 用法

来源：

- [官方上游文档](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/doc/unnest_ordinality.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/unnest_ordinality/)
- [官方上游 README](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/README.md)

`unnest_ordinality` — 将数组展开并为每个元素返回序号位置的旧版函数扩展。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "unnest_ordinality";
SELECT extversion
FROM pg_extension
WHERE extname = 'unnest_ordinality';
```

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
