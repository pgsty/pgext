## 用法

来源：

- [官方上游 README](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/README.md)
- [官方扩展控制文件 (supabase_custom_claims.control)](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/supabase_custom_claims.control)
- [官方扩展 SQL (supabase_custom_claims--1.0.sql)](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/supabase_custom_claims--1.0.sql)

`supabase_custom_claims` — 这是实现 Supabase 项目自定义声明的一种方式。在实现相应的安全、审计或访问控制工作流时，请使用它。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION supabase_custom_claims;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
