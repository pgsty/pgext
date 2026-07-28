## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/README.md)
- [官方扩展控制文件 (hello-world.control)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/hello-world.control)
- [官方扩展 SQL (hello-world--0.0.1.sql)](https://github.com/constructive-io/supabase-test-suite/blob/44ceff6be0b62e1d1a60524acc0c65c1b8343726/packages/hello-world/sql/hello-world--0.0.1.sql)

`hello-world` — **🛠 构建于 Constructive 团队 — Postgres 工具化模块的创造者，致力于构建安全、可组合的后端。如果您喜欢我们的工作，请在 GitHub 上贡献。**。在实现相应的安全、审计或访问控制工作流时使用它。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "hello-world";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `rls_test.update_updated_at_column()` 是一个扩展函数，返回 `trigger`。
- `rls_test.pets` 是由扩展安装或管理的表。
- `rls_test` 是由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 请先安装并验证确认的扩展依赖项：`plpgsql`、`pgcrypto`、`supabase`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
