## 用法

来源：

- [项目 README](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/README.md)
- [扩展软件包元数据](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/Cargo.toml)
- [触发器与状态实现](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/src/lib.rs)
- [扩展管理的表](https://github.com/vkrinitsyn/rppd/blob/51e263d221d0f1dbafa8667eda5eaf176e389788/rppd-pg/pg_setup.sql)

`rppd` 0.2.2 是 Rust Python Postgres Discovery 的 PostgreSQL 触发器组件。表事件会发送到外部 RPPD 服务；该服务选择扩展管理表中存储的 Python 源码，并针对 PostgreSQL 或可选的 etcd 队列执行。它还建模了节点注册、执行日志和 cron 调度。

### 检查扩展状态

安装本地库并创建扩展后，应先检查配置，再为应用表挂载触发器：

```sql
CREATE EXTENSION rppd;

SELECT id, host, host_name, active_since, master
FROM rppd_config;

SELECT id, schema_table, topic, queue, fn_logging
FROM rppd_function;

SELECT rppd_info();
```

应用表通过触发器函数连接服务：

```sql
CREATE TRIGGER orders_rppd
AFTER INSERT OR UPDATE OR DELETE ON orders
FOR EACH ROW EXECUTE FUNCTION rppd_event();
```

`rppd_function` 中由管理员控制的匹配记录决定 Python 代码、表或队列 topic、串行方式、日志和清理行为。启用触发器前，应先启动并验证外部 RPPD 服务。

### 远程代码与事务边界

该扩展标记为需要超级用户、目标平台为 Linux，并且有意把数据库记录转化为可执行 Python。任何能够修改 `rppd_function`、其触发器路径、服务配置或 etcd 消息的主体，都可能利用守护进程的数据库和操作系统权限影响代码执行。应撤销直接写表权限，使用独立沙箱化的服务身份，对每个网络链路进行认证和加密，在数据库外对已审查代码签名，并审计每次配置变更。

触发器投递与外部执行并不处于源记录变更的同一事务。上游 README 指出服务可能尚看不到刚插入的记录，并建议延迟访问，但这不是正确性保证。必须显式设计幂等、重试、顺序、崩溃恢复和可观测性，不要假设 exactly-once 执行。存储的输出和错误也可能暴露应用数据，因此应限制 `rppd_function_log` 的保留期和访问权限。
