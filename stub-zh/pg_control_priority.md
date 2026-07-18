## 用法

来源：

- [上游 README](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/README.md)
- [1.0 版安装 SQL](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/pg_control_priority--1.0.sql)
- [优先级控制实现](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/pg_control_priority.c)

pg_control_priority 1.0 读取或提高 PostgreSQL 进程的 POSIX nice 数值。nice 数字越大，CPU 调度优先级越低；它不控制查询、锁、存储或网络优先级。

### 降低当前测试会话优先级

```sql
CREATE EXTENSION pg_control_priority;
SELECT pg_get_priority(pg_backend_pid());
SELECT pg_set_priority(pg_backend_pid(), 5);
SELECT pg_get_priority(pg_backend_pid());
```

应在可丢弃会话中执行，因为后端会一直保持较低优先级直至退出。安装 SQL 撤销了两个函数的公共执行权限，因此调用需要超级用户权限或显式授权。

### 启动配置

要持续应用调度优先级设置，应预加载动态库并重启：

```conf
shared_preload_libraries = 'pg_control_priority'
pg_control_priority.scheduling_priority = 5
```

### 注意事项

- PostgreSQL 进程通常可以降低自己的 CPU 优先级，但没有操作系统权限就无法再提高。常见 Unix 系统的有效 nice 范围最大为 19，尽管该设置接受 20。
- 在操作系统权限允许时，传入其他 PostgreSQL 进程标识符可能影响客户端后端、辅助进程或主进程。必须严格限制执行权限。
- 代码使用 PostgreSQL 进程检查与 Unix 调度接口。它无法移植到 Windows，且必须针对精确 PostgreSQL 与操作系统目标构建和测试。
