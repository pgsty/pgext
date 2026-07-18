## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed.control)
- [官方上游文档](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/README.md)

`wildspeed` — 使用置换词索引加速 LIKE 通配符搜索的 GIN 操作符类扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "wildspeed";
SELECT extversion
FROM pg_extension
WHERE extname = 'wildspeed';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
