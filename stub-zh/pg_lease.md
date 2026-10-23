## 用法

来源：

- [PGXN 0.1.0](https://pgxn.org/dist/pg_lease/0.1.0/)
- [官方 README](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/README.md)
- [扩展 SQL](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/sql/pg_lease--0.1.0.sql)
- [控制文件](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/pg_lease.control)

`pg_lease` 为具名资源提供可过期的所有权与递增的隔离令牌。它是纯 SQL/PL/pgSQL 扩展，采用惰性过期，无后台工作进程，也不需要预加载。PGXN 将 0.1.0 标为 testing；上游要求 PostgreSQL 14 或更高版本，文档记录了 PostgreSQL 14 和 15 的兼容性测试。

### 核心流程

扩展标为可信：拥有数据库内创建权限的角色可以安装。它创建固定的 `lease` 模式，不能迁移到其他模式。

```sql
CREATE EXTENSION pg_lease;

SELECT acquired, epoch, expires_at
FROM lease.acquire('job-42', 'worker-a', interval '30 seconds');

SELECT status, expires_at
FROM lease.renew('job-42', 'worker-a', 1, interval '30 seconds');

SELECT status FROM lease.release('job-42', 'worker-a', 1);
SELECT held, owner, epoch, expires_at FROM lease.inspect('job-42');
```

示例使用首次获取租约返回的令牌 1。应用必须把实际取得的令牌传给后续续租、释放操作以及外部资源的校验逻辑。成功释放会推进令牌，再次获取后必须使用新返回值。

### 操作接口

| 操作 | 行为 |
|---|---|
| `lease.acquire(key, owner, ttl, wait)` | 获取空闲或已过期的资源。等待时间默认零，资源被占用时返回失败；正值表示轮询直到成功或超时。 |
| `lease.renew(key, owner, epoch, ttl)` | 延长有效租约，保持令牌不变。 |
| `lease.release(key, owner, epoch)` | 释放有效租约并推进令牌。 |
| `lease.inspect(key)` | 查询所有权、到期时间与当前令牌；未知资源不返回记录。 |

续租与释放返回 `ok`、`not_owner`、`lapsed` 或 `epoch_mismatch`。无效输入抛出 SQLSTATE `22023`；阻塞等待超时抛出 `57014`。同一持有者再次获取尚未过期的租约是幂等操作，不延长到期时间。

### 事务与边界

阻塞获取应使用 READ COMMITTED。状态变更属于调用者事务，回滚会撤销变更。过期判断使用数据库服务器时钟；会话结束不会立即释放租约。函数采用固定搜索路径的 `SECURITY DEFINER`，并撤销公众直接访问 `lease.leases` 的权限。

持有者字符串由应用提供，不是经过认证的数据库身份。应用必须限制租约接口的访问，并在受保护资源处检查隔离令牌。扩展不保证主备切换或独立数据库之间的安全性。删除扩展会丢失租约状态，重新安装会重置令牌历史；0.1.0 没有为租约表注册扩展配置转储。
