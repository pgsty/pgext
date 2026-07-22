## 用法

来源：

- [Official usage documentation](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/doc/pgbouncer_wrapper.md)
- [Version 1.2.2 control file](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/pgbouncer_wrapper.control)
- [Views, foreign server, and control functions](https://github.com/davidfetter/pgbouncer_wrapper/blob/07be9547479c765588098843671abdcadbf45e33/sql/pgbouncer_wrapper.sql)

`pgbouncer_wrapper` 1.2.2 通过 `dblink_fdw` 把 PgBouncer 管理控制台暴露为 PostgreSQL 视图和 SQL 控制函数。它适合基于 SQL 的监控和严格受控的运维，但会把 PgBouncer 管理身份委托给数据库调用者，因此必须作为高权限管理接口对待。

### 连接包装层

安装 SQL 会创建名为 `pgbouncer` 的外部服务器：Unix socket 主机 `/tmp`、端口 `6432`、管理数据库 `pgbouncer`，并创建一个使用 `pgbouncer` 用户的 `PUBLIC` 用户映射：

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pgbouncer_wrapper;

SELECT * FROM pgbouncer.version;
SELECT database, "user", cl_active, cl_waiting, sv_active, sv_idle
FROM pgbouncer.pools;
SELECT database, total_query_count, avg_query_time
FROM pgbouncer.stats;
```

如果管理控制台使用不同 socket、端口、用户或认证方式，应在查询前通过高权限迁移修改 `pgbouncer` 外部服务器和用户映射。每个视图都为某条 `SHOW` 命令声明固定列布局，因此必须测试准确的 PgBouncer 版本。

### 监控与控制对象

监控视图包括 `clients`、`servers`、`pools`、`databases`、`stats`、`stats_averages`、`stats_totals`、`totals`、`config`、`lists`、`mem`、`sockets`、`active_sockets`、`users`、`version` 以及 DNS/state 视图。

控制函数包括 `disable`、`enable`、`kill`、`pause`、`reconnect`、`reload`、`resume`、`shutdown`、`suspend`、`wait_close` 和 `pgbouncer.set`。它们会直接向管理控制台发送命令；`kill`、`shutdown`、`pause` 和 `suspend` 可能中断服务。

### 权限与兼容性边界

安装过程创建 `CREATE USER MAPPING FOR PUBLIC`，因此以后获准访问 `pgbouncer` 模式的每个数据库角色都能使用同一个 PgBouncer 映射身份。应撤销默认函数执行权限，避免广泛授予模式使用权，并把只读视图与控制函数分开授权。不要随意暴露 `pgbouncer.fds`：其声明输出包含密码与 SCRAM key 字段，底层 `SHOW FDS` 命令还可能阻塞 PgBouncer 事件循环。

`pgbouncer.set(key, value)` 会把 key 作为原始命令文本格式化，因此只允许可信且经过校验的 key。所有控制函数参数都应视为管理员输入。

PgBouncer 不同版本会改变 `SHOW` 列。版本不匹配可能使视图失败或错误解释字段；自动化控制调用前，应按部署版本文档验证每个视图，并演练故障切换和重启。该包装器不是遥测缓存：每次查询都会发起实时 `dblink` 请求，继承管理控制台的可用性和延迟。
