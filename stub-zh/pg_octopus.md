## 用法

来源：

- [官方 README](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/README.md)
- [官方扩展 SQL](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/sql/pg_octopus.sql)
- [官方健康检查实现](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/src/pg_octopus.c)

`pg_octopus` 是一个实验性后台工作进程，用于记录它能否通过 libpq 连接到一组 PostgreSQL 主机/端口。它只检查 TCP/数据库登录可达性，不评估复制状态、读写就绪、延迟、仲裁或应用健康。

### 启用工作进程

必须预加载该工作进程并重启 PostgreSQL。其实现固定连接本地名为 `postgres` 的数据库，因此应在该库中创建扩展和节点表：

```ini
shared_preload_libraries = 'pg_octopus'
pg_octopus.health_check_period = 10000
pg_octopus.health_check_timeout = 2000
pg_octopus.health_check_max_retries = 2
pg_octopus.health_check_retry_delay = 1000
```

```sql
CREATE EXTENSION pg_octopus;

INSERT INTO octopus.nodes(node_name, node_port)
VALUES ('db-a.internal', 5432), ('db-b.internal', 5432);

SELECT node_name, node_port, health_status
FROM octopus.nodes
ORDER BY node_name, node_port;
```

`health_status` 为 `1` 表示可达，为 `0` 表示重试后仍不可达，为 `-1` 表示尚无结果。

### 配置索引

- `pg_octopus.health_check_period` 设置一轮检查的毫秒时长。
- `pg_octopus.health_check_timeout` 设置每次连接的毫秒超时。
- `pg_octopus.health_check_max_retries` 限制首次尝试后的重试次数。
- `pg_octopus.health_check_retry_delay` 设置尝试之间的延迟。
- `pg_octopus.nodes_table` 修改工作进程读写的表名。

超时加上所有重试延迟与重试超时应小于检查周期，避免每轮持续超期。

### 运维边界

远程检查始终连接数据库 `postgres`，连接串中不含用户名或密码；认证因此由 libpq 默认值、工作进程环境和 `pg_hba.conf` 决定。连接成功只证明该路径和凭据可以登录。工作进程还配置为 `BGW_NEVER_RESTART`，所以崩溃后会停止检查，直至 PostgreSQL 重启。本次复核的是年代较早的 Citus 演示源码，并使用动态拼接 SQL 处理配置表和节点值。应限制节点表写权限，在外部监控工作进程及结果时间戳，测试重载行为，并且不能仅凭 `health_status` 做自动故障转移。
