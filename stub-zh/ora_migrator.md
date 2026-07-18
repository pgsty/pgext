## 用法

来源：

- [官方上游文档](https://github.com/cybertec-postgresql/ora_migrator/blob/c1f90bb8f1b1d50929b3e2211033467f39535980/README.md)

`ora_migrator` — ora_migrator：用于将 Oracle 数据库迁移到 PostgreSQL 的工具。

已复核目录快照记录的版本为 `1.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`db_migrator`, `oracle_fdw`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ora_migrator";
SELECT extversion
FROM pg_extension
WHERE extname = 'ora_migrator';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
