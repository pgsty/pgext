## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_gsl/)

`pg_gsl` — 为 PostgreSQL 封装部分 GNU Scientific Library 例程，包括实数 FFT 变换及辅助函数。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "pg_gsl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_gsl';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
