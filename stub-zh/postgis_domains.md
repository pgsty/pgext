## 用法

来源：

- [官方扩展控制文件](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains.control)
- [官方上游文档](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/README.md)

`postgis_domains` — 定义可复用的 PostGIS geometry/geography 域，约束为有效、简单的 SRID 4326 值

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`postgis`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postgis_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgis_domains';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
