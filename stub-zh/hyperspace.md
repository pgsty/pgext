## 用法

来源：

- [官方扩展控制文件](https://github.com/JonM0/hyperspace/blob/99ebb600ef1b5e606a54e79978a6824a33492d3f/hyperspace.control)

`hyperspace` — 提供四维点、盒和圆类型，以及 B-tree 与 SP-GiST/KD-tree 操作符类。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "hyperspace";
SELECT extversion
FROM pg_extension
WHERE extname = 'hyperspace';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
