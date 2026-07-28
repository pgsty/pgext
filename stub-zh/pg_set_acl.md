## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/README.md)
- [官方扩展控制文件 (pg_set_acl.control)](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/pg_set_acl.control)
- [官方扩展 SQL (pg_set_acl--0.0.1.sql)](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/pg_set_acl--0.0.1.sql)

`pg_set_acl` — 一个实现 SET 命令访问控制列表的 PostgreSQL 扩展。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_set_acl;

select set_acl.grant(setting, user);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `set_acl.grant(cstring, cstring)` 是一个扩展函数。
- `set_acl.read_acl(cstring, cstring)` 是一个扩展函数。
- `set_acl.revoke(cstring, cstring)` 是一个扩展函数。
- `set_acl.privs` 是一个由扩展安装或管理的表。
- `set_acl` 是一个由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
