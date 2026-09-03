## 用法

来源：

- [官方文档](https://github.com/mbpcore/pg_login_guard/blob/6f69b9eff49b3d03bc365314884b3fd4e182dd19/README.md)
- [扩展控制文件](https://github.com/mbpcore/pg_login_guard/blob/6f69b9eff49b3d03bc365314884b3fd4e182dd19/pg_login_guard.control)
- [官方仓库](https://github.com/mbpcore/pg_login_guard)

`pg_login_guard` 在连续认证失败后按角色执行类似 fail2ban 的锁定。

### 启用

将 `pg_login_guard` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_login_guard`：

```ini
shared_preload_libraries = 'pg_login_guard'
```

```sql
CREATE EXTENSION pg_login_guard;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- See everyone currently being tracked (has a recent failure, or is locked):
SELECT * FROM pg_login_guard_status();

--  role_name | failed_attempts | window_started_at      | locked_until
-- -----------+------------------+-------------------------+-------------------------
--  alice     |                3 | 2026-08-15 10:00:01+00  |
--  bob       |                5 | 2026-08-15 10:01:40+00  | 2026-08-15 10:16:40+00

-- Manually unlock a role before its lockout expires (superuser only):
SELECT pg_login_guard_unlock('bob');

-- Forget all tracking history for a role (superuser only):
SELECT pg_login_guard_reset('bob');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_login_guard_status` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_login_guard_reset` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_login_guard_unlock` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_login_guard` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
