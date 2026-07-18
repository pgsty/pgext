## 用法

来源：

- [固定提交的上游 README](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/README.md)
- [固定提交的 1.3 版安装 SQL](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/sql/pg_sessions--1.3.sql)
- [固定提交的 C 实现](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/src/pg_sessions.c)
- [固定提交的发行元数据](https://github.com/menardorama/pg_sessions/blob/6471e4f563f14f77f5354081d46f40e011f0ca9/META.json)
- [正式 PGXN 发行页面](https://pgxn.org/dist/pg_sessions/)

pg_sessions 是从旧版 pg_stat_statements 代码派生的遗留会话与规范化语句统计收集器。它以用户、数据库、进程 ID 和查询 ID 组合作为条目键，并把执行计数器与 Linux 进程的 CPU、内存和 I/O 指标结合起来。

### 预加载与查询

该模块需要共享内存与执行器钩子，因此要在服务器启动前配置：

```conf
shared_preload_libraries = 'pg_sessions'
pg_sessions.max = 5000
pg_sessions.track = 'all'
pg_sessions.track_utility = on
pg_sessions.track_all_steps = on
pg_sessions.track_system_metrics = on
```

重启 PostgreSQL 后：

```sql
CREATE EXTENSION pg_sessions;

SELECT datname, pid, queryid, status, calls, total_time,
       resident_memory_size, bytes_reads, bytes_writes
FROM pg_sessions
ORDER BY last_executed_timestamp DESC;
```

该视图被授予 PUBLIC，但 C 实现会对非超级用户隐藏属于其他用户的查询 ID 和查询文本。重置函数则被撤销了 PUBLIC 权限。

### 版本与运行限制

扩展控制文件和安装 SQL 使用对象版本 1.3，与本目录一致；仓库 META 文件和正式 PGXN 软件包则把发行版本标记为 0.0.5。上游要求 PostgreSQL 9.5，并说明更老版本无法工作；对于仍受维护的 PostgreSQL 版本没有当前兼容声明。源码使用已经过时的服务器内部接口，不能假定它能在现代主版本上编译或运行。

系统指标读取 Linux procfs，不能移植到其他操作系统。跟踪每个执行器步骤会给所有语句增加钩子与写入活动，共享哈希表和外部查询文本文件也会消耗内存与存储。C 默认会开启工具语句跟踪、关闭时保存、全步骤跟踪与系统指标，尽管 README 示例把保存设为关闭。任何运维使用前都要测量开销、限制视图访问、决定是否允许保留查询文本，并测试崩溃/重启后的持久化行为。
