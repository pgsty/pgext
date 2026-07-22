## 用法

来源：

- [官方文档](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/docs/index.md)
- [官方 README](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/README.md)
- [Hook 与 GUC 实现](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/src/pgedge_safesession.c)

`pgedge_safesession` 1.0-alpha1 为配置的 PostgreSQL 角色强制只读会话。执行器和工具 hook 阻止 DML、DDL、危险工具命令与特定 C 函数，PostgreSQL 只读事务状态再提供第二层保护。它是面向 PostgreSQL 14+ 的预览安全控制，不能替代正常的最小权限授权。

### 预加载与配置

Hook 只有在服务器启动时加载才会工作：

```conf
shared_preload_libraries = 'pgedge_safesession'
```

重启 PostgreSQL 后，由超级用户配置。把已经加载的模块登记到 `pg_extension` 是可选操作，并不会激活 hook：

```sql
ALTER SYSTEM SET pgedge_safesession.roles =
  'readonly_user,reporting_role';
SELECT pg_reload_conf();
```

限制锚定到 `session_user`。配置角色的成员同样受限；`SET ROLE` 和具有特权的 `SECURITY DEFINER` 所有者都不能绕过检查。登录角色为超级用户的会话始终豁免。

### 保护控制

全部控制都是 `SUSET`。`pgedge_safesession.block_dml`、`block_ddl`、`block_c_functions` 和 `force_read_only` 默认开启。`block_all_c_functions` 默认关闭，因此默认只阻止 volatile C 函数；开启后，stable/immutable 扩展计算也会被阻止。

在默认设置下，受限会话可以使用 SELECT、不执行语句的 EXPLAIN、事务控制、安全 SET/SHOW、LISTEN/NOTIFY、游标、只读 PL/pgSQL/SQL 函数以及普通 `COPY TO`。写入、MERGE、DDL、COPY FROM、COPY TO PROGRAM、CTAS/SELECT INTO、GRANT/REVOKE、VACUUM/ANALYZE、强表锁、受保护只读 GUC 变更以及 volatile C 函数都会被阻止。

### 安全边界

角色授权仍是首要控制：SafeSession 不会让可读数据自动安全，也不进行行过滤、值脱敏、用户认证或防范豁免超级用户。应审查部署中每个允许的工具操作；若不信任 C 函数的易变性标记，应启用 `block_all_c_functions`。由于这是依赖服务器 hook 的 alpha 软件，在把它作为安全边界前，应针对精确 PostgreSQL 大版本验证命令覆盖、扩展交互、预备语句、连接池、重载和故障转移。
