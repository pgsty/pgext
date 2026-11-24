## 用法

来源：

- [README](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/README.md)
- [Control](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login.control)
- [SQL 1.0](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login--1.0.sql)
- [Authentication hook](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login.c)

`pg_stat_login` 1.0 使用 PostgreSQL 自定义累积统计 API，按用户与数据库组合记录认证成功和失败次数，需要 PostgreSQL 18 及以上版本。它与面向 PostgreSQL 15–17 的旧扩展 `pg_login_stat` 是不同的扩展。

### 启用与检查

将库追加到现有预加载列表，重启 PostgreSQL，再由超级用户安装 SQL 对象：

```ini
shared_preload_libraries = 'pg_stat_login'
```

```sql
CREATE EXTENSION pg_stat_login;
SELECT username, datname, login_ok, login_fail,
       last_login_ok, last_login_fail
FROM pg_stat_logins;
SELECT * FROM pg_stat_login_info();
```

`pg_read_all_stats` 成员可以读取视图和检查函数，只有超级用户能够重置统计。扩展不可重定位。

### 容量与维护

共享表最多容纳 10,000 个不同组合，容量固定。`pg_stat_login_info()` 返回 `max_entries`、`num_entries`、`n_dropped` 和 `stats_reset`。容量耗尽后，已有组合继续累计，新组合的事件则被丢弃；失败尝试同样占用槽位。

```sql
SELECT pg_stat_login_reset();
ALTER SYSTEM SET pg_stat_login.enable = off;
SELECT pg_reload_conf();
```

重置会清除全部累计计数并回收槽位。可重载的 `pg_stat_login.enable` 默认开启。统计能跨正常重启保留，但崩溃恢复后会丢弃，因此它是累积监控数据，不能充当持久审计日志。

### 认证边界

`login_ok` 表示认证通过，不代表成功建立了可用会话。认证之后出现的数据库不存在、仅剩保留连接槽等错误，仍可能计入认证成功。在认证钩子运行前被拒绝的连接，包括连接数耗尽和 `pg_hba.conf` 拒绝，不会被统计；客户端尚未提供密码就断开的尝试也被排除。

名称按客户端提交的内容保存，包括不存在的角色或数据库。数据库编码中的非法字节会被转义；合法但客户端编码无法表示的名称可能导致查询统计失败。解释计数时应结合 PostgreSQL 目录核对名称，并查看服务器日志中不经过钩子的事件。本扩展使用自定义统计类型 ID 30。
