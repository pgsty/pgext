## 用法

来源：

- [官方上游 README](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/README.md)
- [官方扩展控制文件 (pg_argon2id.control)](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/pg_argon2id.control)
- [官方扩展 SQL (pg_argon2id--1.0.sql)](https://github.com/polyaklaci/pg_argon2id/blob/df949026ea01e9978092b24582ba6f7cb3897a6d/sql/pg_argon2id--1.0.sql)

`pg_argon2id` — 一个使用 Argon2id 算法提供安全密码哈希的 PostgreSQL 扩展。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_argon2id;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `argon2id_hash(password text, salt_length integer DEFAULT 16, time_cost integer DEFAULT 3, memory_cost integer DEFAULT 262144, parallelism integer DEFAULT 4)` 是一个扩展函数，返回 `text`。
- `argon2id_verify(password text, hash text)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
