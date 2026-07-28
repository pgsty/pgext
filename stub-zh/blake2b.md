## 用法

来源：

- [官方上游 README](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/README.md)
- [官方扩展控制文件 (blake2b.control)](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/blake2b.control)
- [官方扩展 SQL (blake2b--1.0.sql)](https://github.com/cardano-community/pg_blake2b/blob/c1144ef6ff938fff9aa49574aa345805b6ca656f/blake2b--1.0.sql)

`blake2b` — PostgreSQL 扩展，用于快速安全哈希。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION blake2b;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `blake2b(data bytea, digest_size integer DEFAULT NULL, key bytea DEFAULT NULL)` 是一个扩展函数，返回 `bytea`。
- `blake2b(data text, digest_size integer DEFAULT NULL, key bytea DEFAULT NULL)` 是一个扩展函数，返回 `bytea`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
