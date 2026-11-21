## 用法

来源：

- [README.md](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/README.md)
- [pg_permissions.control](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions.control)
- [pg_permissions--1.5--1.6.sql](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions--1.5--1.6.sql)

`pg_permissions` 比较实际权限与预期基线，并可生成或执行修正授权。1.6 要求 PostgreSQL 15 及以上版本。它是纯 SQL 扩展，无需预加载，不可重定位，须在首次安装时选定模式；安装者需要该模式的 CREATE 权限。

### 定义并检查基线

下面示例中的应用角色和模式须已存在：

```sql
CREATE EXTENSION pg_permissions SCHEMA public;
INSERT INTO public.permission_target
  (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{SELECT,INSERT,UPDATE,DELETE}', 'TABLE', 'appschema'),
       ('appuser', '{USAGE}', 'SCHEMA', 'appschema');
SELECT * FROM public.permission_diffs();
SELECT * FROM public.permission_diff_summary ORDER BY objects DESC;
SELECT public.permission_apply();
```

missing 为 true 表示缺少要求的权限，为 false 表示权限多余。对象名为 NULL 表示覆盖该模式下指定类型的全部对象。审计范围限于目标对象、受管理角色及继承其权限的角色；更广泛的检查可用 `all_permissions`。超级用户因隐含拥有全部权限而不显示。

### 先检查再执行

`permission_fix_sql()` 返回修正 SQL。`permission_apply()` 默认只预演；`permission_apply(false)` 使用调用者权限实际执行。权限视图中可写的 `granted` 列同样会执行 GRANT 或 REVOKE。基线写入权限与修改实际授权的权限都应受控。

对象视图覆盖数据库、模式、表、视图、列、函数和序列，`source` 列区分所有权、直接授权、PUBLIC 和继承权限。`PUBLIC` 可显式作为目标被授权者。`permission_target_snapshot(role, schema)` 将当前授权采纳为基线，会改变预期状态，应有意执行。

`default_acl_target`、`default_acl_permissions` 与 `default_acl_diffs()` 单独处理 ALTER DEFAULT PRIVILEGES。默认权限影响未来对象，不能修复已有授权。基线表允许 PUBLIC 读取，仅扩展所有者及明确授权的角色可写。`role_dependencies()` 检查依赖；`role_drop()` 可转移所有权并删除角色，应另行审阅。

### 升级边界

按版本逐步执行升级脚本。已安装 1.4 的数据库可执行：

```sql
ALTER EXTENSION pg_permissions UPDATE TO '1.5';
ALTER EXTENSION pg_permissions UPDATE TO '1.6';
```

1.6 补齐缺失的 1.2 到 1.3 升级路径。1.5 到 1.6 迁移会重写基线：删除空权限列表，将未指定列名的 COLUMN 条目转为表级条目，并合并重复项。升级前应备份并检查预期状态。PostgreSQL 17 之前不接受 MAINTAIN 权限。
