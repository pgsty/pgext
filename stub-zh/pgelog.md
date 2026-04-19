## 用法

来源：[README](https://github.com/anfiau/pgelog/blob/master/README.md)，[control file](https://github.com/anfiau/pgelog/blob/master/pgelog.control)，[extension SQL 1.0.2](https://github.com/anfiau/pgelog/blob/master/pgelog--1.0.2.sql)，[Tag v1.0.2](https://github.com/anfiau/pgelog/tree/v1.0.2)

`pgelog` 通过使用 `dblink` 实现的 pseudo-autonomous transactions 写入可抵抗 rollback 的日志行。它会在 `pg_variables` 中缓存额外连接，因此同一 session 内重复写日志的成本更低。

### 依赖与安装

- PostgreSQL 11+
- `dblink`
- `pg_variables`
- 无密码本地 `dblink` 访问，通常通过 `peer`

```sql
CREATE EXTENSION IF NOT EXISTS dblink;
CREATE EXTENSION IF NOT EXISTS pg_variables;
CREATE EXTENSION pgelog;
```

### 主要对象与函数

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');

SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

- `pgelog_logs`：基础日志表。
- `pgelog_vw_logs`：带时间差的日志视图。
- `pgelog_params`：配置表。
- `pgelog_to_log(...)`：写入一条不会随 rollback 消失的日志。
- `pgelog_get_param(...)`、`pgelog_set_param(...)`、`pgelog_delete_param(...)`：管理扩展设置。

### 典型用法

官方 README 展示了在 PL/pgSQL exception handler 中使用 `pgelog_to_log(...)`：先用 `GET STACKED DIAGNOSTICS` 收集诊断信息，再写入一条 `FAIL` 日志，最后重新抛出错误。

### 注意事项

- 每个 session 最多会额外打开一个 `dblink` 连接，因此 `max_connections` 需要把这一点算进去。
- 上游 `v1.0.2` 仍然使用同一组用户可见对象；Pigsty 提到的运行时 `dblink` 与 `pg_variables` 依赖，已被 control file 与 README 共同确认。
