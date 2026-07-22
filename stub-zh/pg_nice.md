## 用法

来源：

- [扩展控制文件](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice.control)
- [0.0.1 版扩展 SQL](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice--0.0.1.sql)
- [C 实现](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/src/nice.c)

`pg_nice` 0.0.1 通过执行 `ionice -c 3 -p <backend-pid>`，降低当前 PostgreSQL 后端的 Linux I/O 调度优先级。它不会修改 CPU nice 级别、其他会话或整个 PostgreSQL 实例。

### 核心流程

```sql
CREATE EXTENSION pg_nice;

SELECT pg_backend_pid(), pg_nice();
```

`pg_nice() RETURNS integer` 声明为 `VOLATILE STRICT`，返回 C 库 `system()` 调用的退出状态。零通常表示外部命令成功；非零值是原始进程状态，不是 PostgreSQL 错误，也不是可移植诊断码。

### 要求与效果

数据库服务器必须运行 Linux，PostgreSQL 服务账号的 `PATH` 中必须能找到 `ionice`，并且该账号必须获准将当前进程调整到空闲 I/O 类 3。修改会作用于后端进程的剩余生命周期，因此可能超过调用它的事务或请求。

由于一个后端服务一个数据库连接，连接池复用会让后续工作继续使用较低 I/O 优先级。如果需要区分，应为低优先级工作打开专用会话。

### 注意事项

实现会从 PostgreSQL 内同步调用 shell。虽然命令只包含后端数字 PID，但它仍会在启动外部进程期间阻塞后端，并依赖 PostgreSQL 之外的操作系统行为。它不能移植到 macOS、BSD 或 Windows，缺少 `ionice` 时也没有回退方案。

应在部署内核上检查返回状态和实际调度器状态。只允许理解服务账号和连接池影响的管理员安装和执行；需要持久、可观测的资源治理时，应使用这个小扩展之外的工作负载管理方案。
