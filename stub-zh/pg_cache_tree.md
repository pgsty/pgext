## 用法

来源：

- [官方扩展控制文件](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/pg_cache_tree.control)
- [官方上游文档](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/doc/pg_cache_tree.md)
- [官方上游 README](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/README.md)

`pg_cache_tree` — 通过触发器和辅助函数把递归父节点、子节点 ID 缓存在数组列中。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_cache_tree";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cache_tree';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
