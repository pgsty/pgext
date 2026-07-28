## 用法

来源：

- [官方上游 README](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/README.md)
- [官方扩展控制文件 (pg_cel.control)](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/pg_cel.control)
- [官方实现源码](https://github.com/thomasdarimont/pgauthz/blob/9d4351743ecce44a1a76cafe796d7c85fa03cf31/extensions/pg-cel/src/lib.rs)

`pg_cel` — 一个小型 pgrx PostgreSQL 扩展，用于评估 Common Expression Language (CEL) 表达式，因此可以使用 CEL 而不是原始 SQL 编写 pgauthz 条件。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_cel;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `cel_compile_check` 是一个扩展函数。
- `cel_eval_bool` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
