## 用法

来源：

- [官方扩展控制文件](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity.control)
- [官方上游文档](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/README.md)

`similarity` — 使用 fstrcmp 类近似匹配算法计算两个字符串归一化相似度得分的 C 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "similarity";
SELECT extversion
FROM pg_extension
WHERE extname = 'similarity';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
