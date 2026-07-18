## 用法

来源：

- [官方上游 README](https://github.com/siose-innova/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/README.md)

`pg_landmetrics` — 用于计算景观格局指标的 PostgreSQL/PostGIS 扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`, `postgis`。

```sql
CREATE EXTENSION "pg_landmetrics";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_landmetrics';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
