## 用法

来源：

- [官方上游文档](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/readme.md)

`town` — 用于创建 JSONB/数组时序表并配置 GIN 与 GiST 索引的轻量 PL/pgSQL 辅助扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`btree_gist`, `plpgsql`。

```sql
CREATE EXTENSION "town";
SELECT extversion
FROM pg_extension
WHERE extname = 'town';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
