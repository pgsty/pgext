## 用法

来源：

- [README](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/README.md)
- [Control](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/pgx_permission_sync.control)
- [Makefile](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/Makefile)
- [Base SQL](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.0.sql)
- [Security upgrade](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.0--1.1.sql)
- [Drift and reconciliation SQL](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.1--1.2.sql)

`pgx_permission_sync` 1.2 通过 dblink 向配置的 PostgreSQL 数据库同步权限 DDL，并提供漂移检测、校正、审计日志和重试。扩展以 SQL/PL/pgSQL 编写，支持 PostgreSQL 14 及以上版本；安装需要 `dblink`、`pgcrypto` 与超级用户权限，无需共享预加载。

### 安装与配置

固定模式为 `dba_pgx_permission_sync`。安装从 1.0 SQL 开始，沿提供的升级链到达 1.2。配置目标时先停用实时同步。保存密码前应设置至少 16 字符的加密密钥；未设置密钥时，实现会告警并以未加密形式保存密码。使用 psql 时，应私下提供 `encryption_key` 和 `sync_password` 变量。目标账户需要 `CREATEROLE` 及相关授权权限。

```sql
CREATE EXTENSION pgx_permission_sync CASCADE;
SELECT dba_pgx_permission_sync.disable_sync();
SELECT dba_pgx_permission_sync.enable_dry_run();
SELECT dba_pgx_permission_sync.set_encryption_key(:'encryption_key');
SELECT dba_pgx_permission_sync.add_environment(
  p_env_name := 'staging', p_host := 'staging-db.internal',
  p_dbname := 'appdb', p_username := 'sync_admin',
  p_password := :'sync_password', p_sslmode := 'verify-full'
);
SELECT * FROM dba_pgx_permission_sync.v_environment_status;
```

### 应用前检查漂移

`detect_drift()` 比较角色、成员关系、对象权限和默认权限。`reconcile()` 默认只预演，只有管理员可以实际应用；全局预演设置也会强制预演。撤销权限需要单独启用，不会删除目标独有角色，特权属性默认阻止，除非显式放行。缺少的角色以无密码方式创建。

```sql
SELECT dba_pgx_permission_sync.detect_drift('staging');
SELECT * FROM dba_pgx_permission_sync.v_drift_open_items;
SELECT kind, drift_type, action, status, fix_sql
FROM dba_pgx_permission_sync.reconcile('staging');
```

### 实时同步与运维

实时事件触发器覆盖支持的权限、策略及模式命令。`CREATE ROLE`、`ALTER ROLE` 和 `DROP ROLE` 是集群级命令，不会触发这些事件；应使用漂移校正或经过审查的 `manual_sync()` 调用。检查规则和预演结果后再启用实际投递。远程调用同步发生在源 DDL 事务内，慢目标会拖延原语句；dblink 在远端的修改与源库不构成分布式原子事务。

```sql
SELECT dba_pgx_permission_sync.disable_dry_run();
SELECT dba_pgx_permission_sync.enable_sync();
SELECT * FROM dba_pgx_permission_sync.get_sync_status();
SELECT * FROM dba_pgx_permission_sync.v_failed_syncs;
SELECT * FROM dba_pgx_permission_sync.process_retry_queue();
```

### 访问控制与限制

应把 `pgsync_admin` 和 `pgsync_operator` 视为可信管理角色：操作员能调用手工远程 SQL。`pgsync_reader` 提供只读可见性。加密密钥位于受限的 `secrets` 表中，加密不能防范该表所有者。漂移比较不覆盖所有权、数据库权限、大对象、密码、逐角色设置及 PostgreSQL 16 的成员关系选项。`role_pattern` 不会过滤授权或撤权命令的被授权者。重试与日志清理需要显式调用或独立配置的调度器。
