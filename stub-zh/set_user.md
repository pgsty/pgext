

## 用法

> [set_user: 带增强日志和控制的用户切换](https://github.com/pgaudit/set_user)

`set_user` 允许切换用户和可选的权限提升，并提供增强的审计日志。当非特权用户必须提升为超级用户或对象所有者角色进行维护任务时，它提供额外的控制层。

```sql
CREATE EXTENSION set_user;
```

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'set_user'
```

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `set_user.block_alter_system` | `on` | 提升权限后阻止 ALTER SYSTEM |
| `set_user.block_copy_program` | `on` | 提升权限后阻止 COPY PROGRAM |
| `set_user.block_log_statement` | `on` | 阻止 SET log_statement；对超级用户强制 `log_statement=all` |
| `set_user.superuser_allowlist` | `*` | 允许提升为超级用户的角色 |
| `set_user.nosuperuser_target_allowlist` | `*` | 允许作为非超级用户目标的角色 |
| `set_user.superuser_audit_tag` | `AUDIT` | 提升权限时附加到 log_line_prefix 的标签 |

### 函数

```sql
-- 切换到非超级用户角色
SELECT set_user('dbclient');

-- 提升为超级用户（需要 set_user_u 的 EXECUTE 权限）
SELECT set_user_u('postgres');

-- 使用令牌切换（重置时需要令牌）
SELECT set_user('dbclient', 'my_secret_token');

-- 重置回原始用户
SELECT reset_user();
SELECT reset_user('my_secret_token');  -- 如果使用了令牌

-- 不可撤销的会话认证切换
SELECT set_session_auth('target_role');
```

### 权限设置

```sql
-- 允许角色切换到非超级用户角色
GRANT EXECUTE ON FUNCTION set_user(text) TO admin;

-- 允许角色提升为超级用户
GRANT EXECUTE ON FUNCTION set_user_u(text) TO dba;
```

### 提升权限时的行为

当提升为超级用户角色时：
- 角色转换会以特定标记记录
- `ALTER SYSTEM` 和 `COPY PROGRAM` 被阻止（如已配置）
- `log_statement` 被强制设为 `all` 以实现完整审计追踪
- `AUDIT` 标签被附加到 `log_line_prefix`
