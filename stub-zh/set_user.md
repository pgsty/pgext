## 用法

来源：

- [README.md](https://github.com/pgaudit/set_user/blob/45fc481de6874c2afc5199d7efbd2e32c66d702e/README.md)
- [set_user.control](https://github.com/pgaudit/set_user/blob/45fc481de6874c2afc5199d7efbd2e32c66d702e/set_user.control)
- [CHANGELOG.md](https://github.com/pgaudit/set_user/blob/45fc481de6874c2afc5199d7efbd2e32c66d702e/CHANGELOG.md)

`set_user` 允许切换用户和可选的权限提升，并提供增强的审计日志。当非特权用户必须提升为超级用户或对象所有者角色进行维护任务时，它提供额外的控制层。

```sql
CREATE EXTENSION set_user;
```

### 配置

加入 `postgresql.conf`，保留其他预加载条目，然后重启 PostgreSQL：

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
-- Switch to a non-superuser role
SELECT set_user('dbclient');

-- Escalate to superuser (requires EXECUTE on set_user_u)
SELECT set_user_u('postgres');

-- Switch with a token (token required for reset)
SELECT set_user('dbclient', 'my_secret_token');

-- Reset back to original user
SELECT reset_user();
SELECT reset_user('my_secret_token');  -- if token was used

-- Irrevocable session auth switch
SELECT set_session_auth('target_role');
```

### 权限设置

```sql
-- Allow role to switch to non-superuser roles
GRANT EXECUTE ON FUNCTION set_user(text) TO admin;

-- Allow role to escalate to superuser
GRANT EXECUTE ON FUNCTION set_user_u(text) TO dba;
```

### 提升权限时的行为

当提升为超级用户角色时：
- 角色转换会以特定标记记录
- `ALTER SYSTEM` 和 `COPY PROGRAM` 被阻止（如已配置）
- `log_statement` 被强制设为 `all` 以实现完整审计追踪
- `AUDIT` 标签被附加到 `log_line_prefix`

### 4.3.0 安全与兼容性

4.3.0 要求 PostgreSQL 14 或更新版本。它修复 RESET ALL 导致的后端崩溃，阻止 DISCARD ALL 和函数 SET 子句绕过重置令牌检查，并以常量时间比较令牌。许可列表中的角色名现在区分大小写，混合大小写名称应正确加引号。

重置令牌只是尽力而为的保护，不能防御恶意超级用户或其他身份切换途径。应显式限制 EXECUTE 授权和两类角色许可列表。依赖修复前，需要安装匹配文件并替换预加载进程。
