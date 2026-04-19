## 用法

来源：[README](https://github.com/supabase/supautils/blob/master/README.md)，[homepage](https://supabase.github.io/supautils/)，[releases](https://github.com/supabase/supautils/releases)

`supautils` 是一个可加载库，允许通过配置把部分原本仅限超级用户的 PostgreSQL 能力安全地开放给非超级用户。上游特别强调：它不会在数据库里创建表、函数或 security label。

### 加载方式

集群级：

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

按角色启用：

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### 特权代理角色能力

README 记录了一个 privileged proxy role，可在不授予 `SUPERUSER` 的前提下创建 publication、foreign data wrapper、event trigger 和 privileged extension。

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

对于 event trigger，README 说明这些触发器会对非超级用户生效，但会跳过超级用户和 reserved roles；同时它也明确记录了一条限制：在创建 publication、foreign data wrapper 或 extension 时，这些触发器不会触发。

### 重要配置项

- `supautils.superuser`
- `supautils.privileged_role`
- `supautils.privileged_role_allowed_configs`
- `supautils.privileged_extensions`
- `supautils.extension_custom_scripts_path`
- `supautils.constrained_extensions`
- `supautils.extensions_parameter_overrides`
- `supautils.policy_grants`
- `supautils.drop_trigger_grants`
- `supautils.reserved_roles`
- `supautils.reserved_memberships`
- `supautils.hint_roles`
- `supautils.log_skipped_evtrigs`

### 常见示例

允许非超级用户创建指定 privileged extension：

```ini
supautils.privileged_extensions = 'hstore'
```

允许某个角色管理自己并不拥有的表上的 RLS policy：

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

在 `CREATE EXTENSION` 时强制把扩展装入指定 schema：

```ini
supautils.extensions_parameter_overrides = '{ "pg_cron": { "schema": "pg_catalog" } }'
```

保护托管服务角色不被 `CREATEROLE` 用户修改：

```ini
supautils.reserved_roles = 'connector, storage_admin'
supautils.reserved_memberships = 'pg_read_server_files'
```

### 发布说明

- `v3.2.1` 发布于 2026-04-02，公开说明以维护类改动为主，没有新增用户可见 SQL 接口。
- `v3.2.0` 新增了缺失 `GRANT` 权限时的 hint。

### 注意事项

这是一个强配置驱动的扩展。编写说明时应优先依据 README 中的 GUC 和行为保证，不要暗示任何上游已明确声明“不会创建”的数据库对象。
