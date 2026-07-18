## 用法

来源：

- [官方扩展控制文件](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/grants_manager.control)
- [官方上游文档](https://github.com/shereshevsky/pg_grants_manager/blob/d77b782f71bb1f7417aff3424680e3de7c54e760/README.md)

`grants_manager` — 用于快照、报告并对齐数据库对象授权的声明式 PL/pgSQL 工具集。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "grants_manager";
SELECT extversion
FROM pg_extension
WHERE extname = 'grants_manager';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
