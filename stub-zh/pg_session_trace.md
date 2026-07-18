## 用法

来源：

- [上游 README](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/README.md)
- [1.0 版安装 SQL](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/sql/pg_session_trace--1.0.sql)
- [会话管理实现](https://github.com/charmitro/pg_session_trace/blob/efb2a0bc68d3d44723e5259d7792755987f452d2/src/pg_session_trace_session.c)

pg_session_trace 将选定 PostgreSQL 会话的详细事件记录到每会话二进制跟踪文件中。它面向 PostgreSQL 15 及更高版本，可捕获语句、执行时间、计划、等待，并在最高级别下捕获绑定值。

### 服务器配置

在安装 SQL 对象前，必须预加载模块并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_session_trace'
pg_session_trace.enabled = on
pg_session_trace.output_directory = 'pg_session_trace'
pg_session_trace.level = 4
```

可用以下语句为后端启用跟踪、检查活动跟踪会话、调整详细程度，以及停止跟踪：

```sql
CREATE EXTENSION pg_session_trace;

SELECT pg_session_trace_enable_session(pg_backend_pid());
SELECT pg_session_trace_set_level(pg_backend_pid(), 4);
SELECT * FROM pg_session_trace_list_sessions();
SELECT pg_session_trace_disable_session(pg_backend_pid());
```

0 级记录基本语句，1 级增加执行细节，4 级包含绑定值。应使用能解答诊断问题的最低级别。

### 注意事项

- 会话激活与控制是特权操作。应严格限制执行权，并记录每次跟踪的启动者。
- 4 级可能捕获作为绑定参数提交的密码、令牌、个人数据和应用载荷。只能跟踪经批准的会话，保护输出目录，并及时删除文件。
- 生成跟踪会增加 CPU、内存、存储和 I/O 开销。二进制环形缓冲区在压力下可能丢失事件；应监控报告的 events_dropped 计数。
- 输出文件是服务器本地运维产物，不是审计日志。文件模式可减少意外访问，但不能取代主机访问控制、加密与保留管理。
- 最大会话容量等启动设置需要重启。在故障发生前，应测试解码工具以及扩展与服务器版本兼容性。
