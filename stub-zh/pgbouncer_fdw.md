

## 用法

> [pgbouncer_fdw: 通过普通 SQL 视图查询 PgBouncer 统计信息并通过 SQL 函数运行 PgBouncer 命令的扩展](https://github.com/CrunchyData/pgbouncer_fdw)

### 创建服务器

```sql
CREATE EXTENSION pgbouncer_fdw;

CREATE SERVER pgbouncer FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host 'localhost', port '6432', dbname 'pgbouncer');
```

对于多个 PgBouncer 实例：

```sql
CREATE SERVER pgbouncer1 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.10', port '6432', dbname 'pgbouncer');
CREATE SERVER pgbouncer2 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.11', port '6432', dbname 'pgbouncer');

INSERT INTO pgbouncer_fdw_targets (target_host) VALUES ('pgbouncer1'), ('pgbouncer2');
UPDATE pgbouncer_fdw_targets SET active = false WHERE target_host = 'pgbouncer';
```

### 创建用户映射

```sql
CREATE USER MAPPING FOR PUBLIC SERVER pgbouncer
  OPTIONS (user 'ccp_monitoring', password 'mypassword');
```

### 可用视图

| 视图 | 描述 |
|------|-------------|
| `pgbouncer_clients` | 客户端连接详情 |
| `pgbouncer_pools` | 连接池统计 |
| `pgbouncer_servers` | 后端服务器状态 |
| `pgbouncer_stats` | 统计摘要 |
| `pgbouncer_databases` | 数据库定义 |
| `pgbouncer_config` | 配置参数 |
| `pgbouncer_lists` | 内部列表 |
| `pgbouncer_dns_hosts` | DNS 主机缓存 |
| `pgbouncer_dns_zones` | DNS 区域缓存 |
| `pgbouncer_sockets` | 套接字信息 |
| `pgbouncer_users` | 用户配置 |

```sql
SELECT * FROM pgbouncer_pools;
SELECT * FROM pgbouncer_stats;
SELECT database, cl_active, cl_waiting, sv_active FROM pgbouncer_pools;
```

监控多个实例时，每行包含一个 `pgbouncer_target_host` 列标识来源。

### 命令函数

管理函数（需要显式 `GRANT EXECUTE`）：

```sql
SELECT pgbouncer_command_reload();              -- 重载配置
SELECT pgbouncer_command_pause('mydb');          -- 暂停数据库
SELECT pgbouncer_command_resume('mydb');         -- 恢复数据库
SELECT pgbouncer_command_kill('mydb');           -- 终止连接
SELECT pgbouncer_command_disable('mydb');        -- 禁用数据库
SELECT pgbouncer_command_enable('mydb');         -- 启用数据库
SELECT pgbouncer_command_reconnect('mydb');      -- 重新连接后端
SELECT pgbouncer_command_set('key', 'value');    -- 设置参数
SELECT pgbouncer_command_shutdown();             -- 关闭 PgBouncer
SELECT pgbouncer_command_suspend();              -- 暂停操作
SELECT pgbouncer_command_wait_close('mydb');     -- 等待连接关闭
```

### 权限

```sql
GRANT USAGE ON FOREIGN SERVER pgbouncer TO monitoring_user;
GRANT SELECT ON pgbouncer_pools TO monitoring_user;
GRANT SELECT ON pgbouncer_stats TO monitoring_user;
GRANT EXECUTE ON FUNCTION pgbouncer_command_reload() TO admin_user;
```
