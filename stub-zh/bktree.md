## 用法

来源：

- [官方扩展控制文件](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree.control)
- [官方上游文档](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/README.md)

`bktree` — 面向 64 位感知哈希和汉明距离搜索的 SP-GiST BK-tree 索引扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "bktree";
SELECT extversion
FROM pg_extension
WHERE extname = 'bktree';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
