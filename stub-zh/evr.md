## 用法

来源：

- [官方扩展控制文件](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/evr.control)
- [官方上游文档](https://github.com/theforeman/postgresql-evr/blob/ff3ad993a913e403f2966a86ba9c0a751fb24357/README.md)

`evr` — evr：用于解析 RPM epoch/version/release 的 PostgreSQL 数据类型与辅助函数。

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "evr";
SELECT extversion
FROM pg_extension
WHERE extname = 'evr';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
