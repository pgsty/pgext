

## 用法

> [pg_permissions: 查看对象权限并与期望状态进行比较](https://github.com/cybertec-postgresql/pg_permissions)

pg_permissions 让你查看数据库对象的实际权限，并与期望的权限状态进行比较。

### 定义期望权限

向 `permission_target` 插入条目以描述应有的权限：

```sql
INSERT INTO permission_target (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{SELECT,INSERT,UPDATE,DELETE}', 'TABLE', 'appschema');

INSERT INTO permission_target (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{USAGE}', 'SCHEMA', 'appschema');

INSERT INTO permission_target (role_name, permissions, object_type, schema_name, object_name)
VALUES ('appuser', '{USAGE}', 'SEQUENCE', 'appschema', 'appseq');
```

将 `object_name` 或 `column_name` 设置为 NULL 可应用于模式中该类型的所有对象。

### 查找权限差异

```sql
SELECT * FROM permission_diffs();
```

返回 `missing = TRUE`（权限应存在但不存在）或 `missing = FALSE`（不应存在的多余权限）的行。

### 查看实际权限

可用视图（所有视图具有相同的列结构）：

- `database_permissions` -- 当前数据库的权限
- `schema_permissions` -- 模式的权限
- `table_permissions` -- 表的权限
- `view_permissions` -- 视图的权限
- `column_permissions` -- 表/视图列的权限
- `function_permissions` -- 函数的权限
- `sequence_permissions` -- 序列的权限
- `all_permissions` -- 以上所有的 UNION

```sql
SELECT * FROM table_permissions WHERE role_name = 'appuser' AND schema_name = 'appschema';
```

### 通过视图 Grant/Revoke

权限视图的 `granted` 列是可更新的——更新它会执行相应的 `GRANT` 或 `REVOKE` 命令。

注意：超级用户不显示在视图中（他们自动拥有所有权限）。
