## 用法

来源：

- [官方上游文档](https://database.dev/supabase/dbdev)
- [官方项目或服务商页面](https://database.dev/)
- [官方上游 README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`dbdev` — 用于从 database.dev 注册表安装软件包的数据库内客户端。

已复核目录快照记录的版本为 `0.0.5`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`http`, `pg_tle`。

```sql
CREATE EXTENSION "dbdev";
SELECT extversion
FROM pg_extension
WHERE extname = 'dbdev';
```

该上游项目与 `Supabase` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
