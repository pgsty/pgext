## 用法

来源：

- [v3.4.4 README](https://github.com/supabase/supautils/blob/v3.4.4/README.md)
- [v3.4.4 release](https://github.com/supabase/supautils/releases/tag/v3.4.4)
- [Version restriction implementation](https://github.com/supabase/supautils/blob/v3.4.4/src/extensions.c)

`supautils` 是一个可加载库，允许通过配置把部分原本仅限超级用户的 PostgreSQL 能力安全地开放给非超级用户。上游特别强调：它不会在数据库里创建表、函数或安全标签。

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

README 记录了一个特权代理角色，可在不授予 `SUPERUSER` 的前提下创建发布、外部数据包装器、事件触发器和特权扩展。

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

对于事件触发器，README 说明这些触发器会对非超级用户生效，但会跳过超级用户和保留角色；同时它也明确记录了一条限制：在创建发布、外部数据包装器或扩展时，这些触发器不会触发。

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

允许非超级用户创建指定特权扩展：

```ini
supautils.privileged_extensions = 'hstore'
```

允许某个角色管理自己并不拥有的表上的 RLS 策略：

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

在 `CREATE EXTENSION` 时强制把扩展装入指定模式：

```ini
supautils.extensions_parameter_overrides = '{ "pg_cron": { "schema": "pg_catalog" } }'
```

保护托管服务角色不被 `CREATEROLE` 用户修改：

```ini
supautils.reserved_roles = 'connector, storage_admin'
supautils.reserved_memberships = 'pg_read_server_files'
```

### 版本选择与运行边界

`supautils.restrict_extension_versions` 控制非超级用户是否可以显式指定版本：`off` 允许；`warn` 忽略指定值，发出警告并选择控制文件默认版本；`error` 拒绝。这同时适用于扩展创建和升级；超级用户及配置的代理超级用户不受约束。不显式指定版本的操作仍可执行，但须通过普通权限检查。

集群预加载需要重启；角色级会话预加载对新连接生效。不要为 supautils 本身执行 CREATE EXTENSION。源码发布版 3.4.4 是库更新，没有 SQL 扩展升级步骤。它避免在允许列表中的表上检查策略时获取 ACCESS EXCLUSIVE 锁，并在退出提权区间的每条路径上恢复调用者角色。

允许的扩展和自定义脚本会使用代理超级用户权限，应把它们作为可信代码审查。带标签的 README 说明 PostgreSQL 18 的视图不支持增强权限提示。扩大授权前，应测试角色切换、事件触发器所有权和保留角色保护。
