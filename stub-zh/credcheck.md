

## 用法

> [credcheck: PostgreSQL 用户名和密码的凭证检查](https://github.com/MigOpsRepos/credcheck)

`credcheck` 在 `CREATE ROLE`、`ALTER ROLE` 和密码修改时，强制执行可配置的用户名和密码强度规则。它还支持密码重用策略和认证失败封禁。

### 配置参数

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'credcheck'
```

#### 用户名检查

| 参数 | 描述 | 示例 |
|-----------|-------------|---------|
| `credcheck.username_min_length` | 用户名最小长度 | `4` |
| `credcheck.username_min_special` | 最少特殊字符数 | `1` |
| `credcheck.username_min_digit` | 最少数字字符数 | `1` |
| `credcheck.username_min_upper` | 最少大写字符数 | `2` |
| `credcheck.username_min_lower` | 最少小写字符数 | `1` |
| `credcheck.username_min_repeat` | 相邻最大重复字符数 | `2` |
| `credcheck.username_contain` | 必须包含其中一个字符 | `a,b,c` |
| `credcheck.username_not_contain` | 不得包含这些字符 | `x,y,z` |
| `credcheck.username_contain_password` | 用户名不得包含密码 | `on` |

#### 密码检查

| 参数 | 描述 | 示例 |
|-----------|-------------|---------|
| `credcheck.password_min_length` | 密码最小长度 | `8` |
| `credcheck.password_min_special` | 最少特殊字符数 | `1` |
| `credcheck.password_min_digit` | 最少数字字符数 | `1` |
| `credcheck.password_min_upper` | 最少大写字符数 | `1` |
| `credcheck.password_min_lower` | 最少小写字符数 | `1` |
| `credcheck.password_min_repeat` | 相邻最大重复字符数 | `3` |
| `credcheck.password_contain_username` | 密码不得包含用户名 | `on` |
| `credcheck.password_valid_until` | VALID UNTIL 最少天数 | `60` |
| `credcheck.password_valid_max` | VALID UNTIL 最大天数 | `365` |
| `credcheck.whitelist` | 不受检查约束的用户名 | `admin,super` |

### 使用示例

```sql
-- 拒绝：用户名太短
CREATE USER abc WITH PASSWORD 'pass';
-- ERROR: username length should match the configured credcheck.username_min_length

-- 拒绝：密码包含用户名
CREATE USER abcd$ WITH PASSWORD 'abcd$xyz';
-- ERROR: password should not contain username
```

### 密码重用策略

```sql
SET credcheck.password_reuse_history = 2;
SET credcheck.password_reuse_interval = 365;  -- 天
```

查看密码历史：

```sql
SELECT rolename, password_hash FROM pg_password_history;
```

### 认证失败封禁

```sql
SET credcheck.max_auth_failure = 3;  -- 3次失败后封禁
```

重置被封禁的用户：

```sql
SELECT pg_banned_role_reset();              -- 重置所有
SELECT pg_banned_role_reset('username');     -- 重置特定用户
```
