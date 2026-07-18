## 用法

来源：

- [官方扩展控制文件](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx.control)
- [官方上游文档](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/README.org)

`sphinx` — 从 PostgreSQL 查询和更新外部 SphinxSearch 服务的 C 连接器。

已复核目录快照记录的版本为 `0.3`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "sphinx";
SELECT extversion
FROM pg_extension
WHERE extname = 'sphinx';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
