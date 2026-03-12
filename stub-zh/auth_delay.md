

## 用法

> [auth_delay: 在报告认证失败前暂停](https://www.postgresql.org/docs/current/auth-delay.html)

`auth_delay` 在报告认证失败前使服务器短暂暂停，使暴力破解密码攻击更加困难。

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'auth_delay'
auth_delay.milliseconds = '500'
```

### 配置参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `auth_delay.milliseconds` | `0` | 报告认证失败前等待的毫秒数 |

### 注意事项

- 必须通过 `shared_preload_libraries` 加载
- 无法防止拒绝服务攻击（等待的进程仍然占用连接槽位）
- 无需 `CREATE EXTENSION` —— 这仅是一个共享库模块
