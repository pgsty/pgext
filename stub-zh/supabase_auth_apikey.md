## 用法

来源：

- [官方主文档](https://database.dev/martindonadieu/supabase_auth_apikey)

`supabase_auth_apikey` — 用于 Supabase Auth 的 API Key 创建、删除、校验与权限检查扩展。

已复核目录快照记录的版本为 `0.0.3`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "supabase_auth_apikey";
SELECT extversion
FROM pg_extension
WHERE extname = 'supabase_auth_apikey';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
