## 用法

来源：

- [官方上游文档](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/README)
- [官方 PGXN 分发页](https://pgxn.org/dist/pgvihash/)

`vihash` — 提供版本无关的哈希函数。

已复核目录快照记录的版本为 `1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "vihash";
SELECT extversion
FROM pg_extension
WHERE extname = 'vihash';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
