

## 用法

> [supautils: 在云环境中保护集群安全的扩展](https://github.com/supabase/supautils)

`supautils` 是一个可加载库，允许非超级用户安全地创建事件触发器、发布和扩展。它完全通过配置管理 —— 不会向数据库添加表、函数或安全标签。

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

或按角色启用：

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### 关键 GUC 参数

| 参数 | 描述 |
|-----------|-------------|
| `supautils.privileged_role` | 超级用户操作的代理角色 |
| `supautils.superuser` | 实际的超级用户（默认为引导用户） |
| `supautils.privileged_extensions` | 允许非超级用户安装的扩展 |
| `supautils.privileged_role_allowed_configs` | 特权角色可更改的超级用户专属设置 |
| `supautils.reserved_roles` | 受保护不可被 CREATEROLE 用户修改的角色 |
| `supautils.reserved_memberships` | 受限不可授予的角色成员关系 |
| `supautils.constrained_extensions` | 定义扩展资源约束的 JSON |
| `supautils.extensions_parameter_overrides` | 覆盖 CREATE EXTENSION 参数的 JSON |
| `supautils.policy_grants` | 授予非所有者 RLS 策略管理权限的 JSON |
| `supautils.drop_trigger_grants` | 授予非所有者触发器删除权限的 JSON |

### 非超级用户发布

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

### 特权扩展

```ini
supautils.privileged_extensions = 'hstore'
```

非超级用户可以创建通常需要超级用户权限的扩展：

```sql
CREATE EXTENSION hstore;
```

### 保留角色

```ini
supautils.reserved_roles = 'connector, storage_admin'
```

具有 CREATEROLE 权限的用户无法 ALTER 或 DROP 这些角色。

### 表所有权绕过（RLS 策略管理）

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

允许 `my_role` 管理其不拥有的表上的 RLS 策略。
