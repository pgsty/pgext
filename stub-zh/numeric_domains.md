## 用法

来源：

- [官方扩展控制文件](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains.control)
- [官方上游文档](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/README.md)

`numeric_domains` — numeric_domains：为 PostgreSQL 提供常用数值范围域类型。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "numeric_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'numeric_domains';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
