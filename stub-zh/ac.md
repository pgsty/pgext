## 用法

来源：

- [官方上游 README](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/README.md)
- [官方扩展控制文件 (ac.control)](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/src/ac.control)
- [官方扩展 SQL (ac--0.0.1.sql)](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/src/ac--0.0.1.sql)

`ac` — 访问控制 Postgres 扩展（包含开发环境）。在实现相应的安全、审计或访问控制工作流时使用它。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION ac;

SELECT ac_policy(
    'user:1',
    ARRAY['read', 'write'],
    ARRAY['delete']
);
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `ac_check(ac_operation, ac_subject, ac_policy[])` 是一个扩展函数，返回 `BOOLEAN`。
- `ac_check(op ac_operation , bindings ac_subject[] , list ac_policy[])` 是一个扩展函数，返回 `BOOLEAN`。
- `ac_list_cleanup(dirty ac_policy[])` 是一个扩展函数，返回 `ac_policy[]`。
- `ac_policy(ac_subject, ac_operation)` 是一个扩展函数，返回 `ac_policy`。
- `ac_policy(ac_subject, ac_operation, ac_operation)` 是一个扩展函数，返回 `ac_policy`。
- `ac_policy(ac_subject, ac_operation[])` 是一个扩展函数，返回 `ac_policy`。
- `ac_policy(ac_subject, ac_operation[], ac_operation)` 是一个扩展函数，返回 `ac_policy`。
- `ac_policy(ac_subject, ac_operation[], ac_operation[])` 是一个扩展函数，返回 `ac_policy`。
- `ac_policy` 是一个扩展定义的类型。
- `ac_operation` 是一个扩展定义的域。
- `ac_subject` 是一个扩展定义的域。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为可重定位。
- 上游明确表示该项目尚未准备好生产使用。
- 上游描述该项目仍处于开发阶段。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
