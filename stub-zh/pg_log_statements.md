## 用法

来源：

- [官方 README](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/README.md)
- [扩展控制文件](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements.control)
- [0.0.2 扩展 SQL](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements--0.0.2.sql)
- [C 实现](https://github.com/pierreforstmann/pg_log_statements/blob/4ee44c244dc1fafe72d3f7060e8fe6970cc1f546/pg_log_statements.c)

`pg_log_statements` 可以针对指定服务器进程，或符合过滤条件的新认证会话，强制设置 `log_statement=all`。它改变服务器日志行为，并不提供可查询的语句历史表。

### 核心流程

预加载动态库、重启 PostgreSQL、安装 SQL 函数，并立即把执行权限制给审计管理员。

```conf
shared_preload_libraries = 'pg_log_statements'
```

```sql
CREATE EXTENSION pg_log_statements;

REVOKE ALL ON FUNCTION pgls_start(integer) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop(integer) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_start_filter(cstring, cstring) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop_filter(cstring, cstring) FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_state() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_conf() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_start_debug() FROM PUBLIC;
REVOKE ALL ON FUNCTION pgls_stop_debug() FROM PUBLIC;

SELECT pgls_start(12345);
SELECT * FROM pgls_state() AS s(pid integer, status text);
SELECT pgls_stop(12345);

SELECT pgls_start_filter('application_name', 'billing-worker');
SELECT * FROM pgls_conf() AS c(filter_type text, filter_value text);
SELECT pgls_stop_filter('application_name', 'billing-worker');
```

PID 模式使用 `pgls_start(integer)` 与 `pgls_stop(integer)`。过滤模式接受 `application_name`、`user_name`、`hostname`、`ip_address` 或 `database_name`，并且只影响安装过滤器之后建立的会话。`pgls_start_debug()` 与 `pgls_stop_debug()` 用于切换认证字段日志。由于 `pgls_state()` 和 `pgls_conf()` 返回匿名记录，调用方必须给出列定义。

### 安全与日志注意事项

扩展函数没有权限检查，若不撤销就会获得公众执行权。不受信用户可以为其他会话启用语句日志，把 SQL 文本及其中的秘密暴露到服务器日志，并制造大量日志。执行权只能授予经过审计的管理员，同时要保护日志存储、保留和访问。

PID 变更要等目标后端下次进入被挂钩的执行器或工具事件才会生效。启用会强制把 `log_statement` 设为 `all`；停用会强制设为 `none`，不会恢复会话原值。过滤器存放在共享内存中，其数量与值长度受实现限制，重启后会消失。目录中的扩展版本是 0.0.2，尽管较新的仓库元数据提到 0.0.4；行为和 SQL 必须以已安装软件包为准。上游声明验证到 PostgreSQL 16，更高版本需要另行测试。
