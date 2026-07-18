## 用法

来源：

- [官方上游文档](https://database.dev/danahartweg/pg_protect_columns)

`pg_protect_columns` — 通过触发器保护指定列，防止其在初始赋值后被更新

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "pg_protect_columns";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_protect_columns';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
