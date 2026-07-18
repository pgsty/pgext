## 用法

来源：

- [官方扩展控制文件](https://github.com/aquameta/pg_catalog_get_defs/blob/master/pg_catalog_get_defs.control)

`pg_catalog_get_defs` — 提供 PL/pgSQL 函数，根据 pg_catalog 元数据重建 CREATE TYPE 定义。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pg_catalog_get_defs";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_catalog_get_defs';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
