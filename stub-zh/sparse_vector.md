## 用法

来源：

- [官方扩展控制文件](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/sparse_vector.control)

`sparse_vector` — 为稀疏浮点向量提供 C 数据类型、归一化、点积与余弦相似度函数。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "sparse_vector";
SELECT extversion
FROM pg_extension
WHERE extname = 'sparse_vector';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
