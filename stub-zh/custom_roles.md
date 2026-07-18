## 用法

来源：

- [官方上游文档](https://database.dev/garyaustin/custom_roles)

`custom_roles` — 面向用户角色管理的 Supabase RLS 辅助表和函数。

已复核目录快照记录的版本为 `0.0.3`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "custom_roles";
SELECT extversion
FROM pg_extension
WHERE extname = 'custom_roles';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
