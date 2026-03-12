

## 用法

> [passwordcheck: 在 CREATE/ALTER ROLE 时检查密码强度](https://www.postgresql.org/docs/current/passwordcheck.html)

`passwordcheck` 在通过 `CREATE ROLE` 或 `ALTER ROLE` 设置密码时验证密码强度。弱密码将被拒绝并返回错误。

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'passwordcheck'
```

### 配置参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `passwordcheck.min_password_length` | `8` | 最小密码长度（字节）（仅超级用户可设置） |

### 工作原理

该模块检查通过 `CREATE ROLE` 或 `ALTER ROLE` 设置的密码：

```sql
-- 如果密码太短或太弱将被拒绝
CREATE ROLE myuser WITH LOGIN PASSWORD 'abc';
-- ERROR: password is too short

-- 足够强的密码将被接受
CREATE ROLE myuser WITH LOGIN PASSWORD 'Str0ng_P@ssword!';
```

### 默认检查

在没有 CrackLib 的情况下，该模块强制执行：
- 最小密码长度（可通过 `passwordcheck.min_password_length` 配置）
- 密码不能是用户名
- 基本复杂度要求

### 限制

- 客户端程序发送的预加密密码无法被完全验证
- 该模块只能猜测加密提交中的实际密码
- 为了更强的安全性，考虑使用外部认证方法（例如 GSSAPI）
- 无需 `CREATE EXTENSION` —— 这仅是一个共享库模块
