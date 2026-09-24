## 用法

来源：

- [README](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/README.md)
- [Control file / 控制文件](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/pg_ros2.control)
- [Cargo.toml](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/Cargo.toml)
- [src/lib.rs](https://github.com/sweatybridge/pg_ros2/blob/4886868056540ecaf8d389401a202260902e177e/src/lib.rs)

`pg_ros2` 以 PostgreSQL 表暴露 ROS 2 图快照，并通过通知传递选定主题的消息。0.2.0 文档面向 Linux/Ubuntu 22.04、PostgreSQL 18、ROS 2 Humble 和 rclrs 运行环境。

### 图快照

PostgreSQL 服务必须获得 ROS 运行环境。将 `pg_ros2` 加入 `shared_preload_libraries`，把 `pg_ros2.database` 设为已有数据库并重启，再以超级用户安装。

```sql
CREATE EXTENSION pg_ros2;
SELECT * FROM nodes;
SELECT * FROM topics;
SELECT * FROM parameters;
SELECT * FROM worker_status;
SELECT * FROM parameter_status;
```

### 主题订阅

监听会话订阅与完整 ROS 主题名相同的通知频道，再在另一个自动提交会话中启动长时间运行的过程：

```sql
LISTEN "/chatter";
```

```sql
SET statement_timeout = 0;
SET client_connection_check_interval = '1s';
CALL subscribe('/chatter');
```

订阅过程可独立于图工作进程运行。可选的 `pg_durable` 集成可将其作为工作流管理，但并非硬性依赖。

### 访问与投递

仅为适当角色授予表读取和过程执行权限，远程参数快照可能含有秘密。通知频道没有逐频道访问控制，同库的其他用户可以监听频道。

图快照最终刷新，工作进程或 ROS 故障后可能陈旧。主题投递是临时且尽力而为的，缓冲有限并可能丢弃消息，不提供重放或持久消费者位点。服务器失败后需重新启动订阅过程，并监控工作状态，不能把旧快照当成当前状态。
