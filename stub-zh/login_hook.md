

## 用法

> [login_hook: 在用户登录时执行代码，类似于 Oracle 的 after logon 触发器](https://github.com/splendiddata/login_hook)

`login_hook` 允许在用户登录数据库时执行自定义 PL/pgSQL 代码。

```sql
CREATE EXTENSION login_hook;
```

### 配置

添加到 `postgresql.conf`：

```ini
session_preload_libraries = 'login_hook'
```

### 创建登录函数

定义一个 `login_hook.login()` 函数，该函数将在每次登录时执行：

```sql
CREATE OR REPLACE FUNCTION login_hook.login() RETURNS VOID LANGUAGE PLPGSQL AS $$
BEGIN
    IF NOT login_hook.is_executing_login_hook() THEN
        RAISE EXCEPTION 'Should only be invoked by login_hook';
    END IF;

    -- 你的登录逻辑：
    RAISE NOTICE 'Hello %', current_user;

EXCEPTION
    WHEN OTHERS THEN
        RAISE LOG 'Error in login_hook.login(): %', SQLERRM;
END
$$;
GRANT EXECUTE ON FUNCTION login_hook.login() TO PUBLIC;
```

需要 `PUBLIC` 授权，因为该函数会为每个连接用户运行。

### 函数

| 函数 | 返回类型 | 描述 |
|----------|---------|-------------|
| `login_hook.is_executing_login_hook()` | `boolean` | 仅在登录钩子执行期间调用时返回 true |
| `login_hook.get_login_hook_version()` | `text` | 返回 login_hook 的编译版本 |
| `login_hook.login()` | `void` | 用户提供的在登录时执行的函数 |

### 重要说明

- 该函数不会在后台进程或恢复模式期间被调用
- 在函数内部处理所有异常 —— 失败将阻止普通用户登录
- 超级用户在函数失败时会收到警告但仍可登录
- 对于 PostgreSQL 17+，可考虑使用原生登录事件触发器
