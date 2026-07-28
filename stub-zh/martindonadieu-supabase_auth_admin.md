## 用法

来源：

- [Official database.dev 包页面](https://database.dev/martindonadieu/supabase_auth_admin)

`martindonadieu-supabase_auth_admin` — Supabase Auth 工具，用于检查用户是否为平台管理员。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "martindonadieu-supabase_auth_admin";
```

在目标数据库中安装扩展，当可用时运行上方最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `public"."is_admin` 是一个扩展函数。
- `public"."set_admin` 是一个扩展函数。
- `public.is_admin()` 是一个扩展函数，并返回 `boolean`。

### 要求与注意事项

- 该目录记录版本为 `0.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
