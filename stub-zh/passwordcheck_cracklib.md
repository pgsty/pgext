

## 用法

> [passwordcheck_cracklib: 使用 cracklib 加强 PostgreSQL 用户密码检查](https://github.com/devrimgunduz/passwordcheck_cracklib)

`passwordcheck_cracklib` 类似于标准 PostgreSQL `passwordcheck` 模块，但使用 cracklib 进行更严格的密码检查。每当通过 `CREATE ROLE` 或 `ALTER ROLE` 设置密码时检查用户密码。如果密码被认为太弱，将被拒绝并以错误终止命令。

### 配置

在 `postgresql.conf` 中添加到 `shared_preload_libraries`：

```ini
shared_preload_libraries = '$libdir/passwordcheck_cracklib'
```

重启 PostgreSQL 以生效。

### 工作原理

加载后，任何设置密码的 `CREATE ROLE` 或 `ALTER ROLE` 命令都会根据 cracklib 字典检查密码。弱密码或容易猜到的密码将被自动拒绝。

```sql
-- 如果密码太弱将被拒绝
CREATE ROLE myuser WITH LOGIN PASSWORD 'password123';
-- ERROR: password is easily cracked

-- 强密码将被接受
CREATE ROLE myuser WITH LOGIN PASSWORD 'X9#kLm$vQ2!pR7';
```

无需 `CREATE EXTENSION` —— 这仅是一个共享库模块。
