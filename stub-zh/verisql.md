## 用法

来源：

- [官方上游 README](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/README.md)
- [官方扩展控制文件 (verisql.control)](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/verisql.control)
- [官方扩展 SQL (verisql--0.1.0.sql)](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/verisql--0.1.0.sql)

`verisql` — 确定性验证 oracle 用于 AI 生成的 SQL，**在数据库内部**——无需额外的服务或客户端库。纯 PL/pgSQL 实现，作为 CREATE EXTENSION 安装。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION verisql;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `verisql.check(p_sql text)` 是一个扩展函数，返回 `TABLE`。
- `verisql.diff(p_sql_a text, p_sql_b text)` 是一个扩展函数，返回 `boolean`。
- `verisql.explain_sanity(p_sql text)` 是一个扩展函数，返回 `TABLE`。
- `verisql.fingerprint(p_sql text)` 是一个扩展函数，返回 `text`。
- `verisql.history_check(p_sql text, p_tag text)` 是一个扩展函数，返回 `TABLE`。
- `verisql.history_record(p_sql text, p_tag text)` 是一个扩展函数，返回 `void`。
- `verisql.query_history` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 控制文件标记该扩展为可信的。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
