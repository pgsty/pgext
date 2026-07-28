## 用法

来源：

- [官方上游 README](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/README.md)
- [官方扩展控制文件 (rls_oso.control)](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/rls_oso.control)
- [官方实现源代码](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/src/lib.rs)

`rls_oso` — 插件，用于在 PostgreSQL 的行级安全策略中使用 Oso 授权库。当实现相应的安全、审计或访问控制工作流时，请使用此插件。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION rls_oso;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `oso_configure_rls` 是一个扩展函数。
- `oso_is_allowed` 是一个扩展函数。
- `oso_reload()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本为 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将该项目描述为一个概念验证。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
