## 用法

来源：

- [官方 README](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/README.md)
- [扩展控制文件](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench.control)
- [扩展 SQL](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench--1.0.sql)
- [钩子与 perf 实现](https://github.com/AyoubKaz07/pg_microbench/blob/b33b11fc47da522b26dde2bf3023c2b48ee1374a/pg_microbench.c)

`pg_microbench` 使用 `perf_event_open` 计数器检测 Linux 上的 PostgreSQL 后端。它可以测量规划器、执行器、工具语句或显式全局 SPI 范围，并以 `NOTICE` 消息输出硬件计数器结果，适合开发实验。

### 核心流程

操作系统必须允许 PostgreSQL 账号打开所需性能事件。应在隔离主机上使用范围尽可能小的能力或安全策略，不要放宽整个主机的 perf 策略。

```sql
CREATE EXTENSION pg_microbench;

REVOKE ALL ON FUNCTION pg_microbench_run(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_microbench_run(text) TO benchmark_admin;

SET pg_microbench.enable = on;
SET pg_microbench.log = on;
SET pg_microbench.scopes = 'planner,executor';
SELECT count(*) FROM generate_series(1, 100000);

SET pg_microbench.scopes = 'global';
SELECT pg_microbench_run('SELECT count(*) FROM generate_series(1, 100000)');
```

`pg_microbench.enable` 打开或关闭每个后端的 perf 文件描述符。`pg_microbench.log` 控制通知输出，`pg_microbench.scopes` 接受由 `global`、`planner`、`executor`、`utility` 组成的逗号分隔子集。报告字段包括墙钟时间、周期、指令、缓存引用与未命中、L1 数据与指令事件、分支和分支预测失败。

### 基准边界

计数器只属于当前会话，输出不会持久化。不受支持或受限制的硬件事件可能不可用，测量值还会随 CPU、内核、PostgreSQL 构建、缓存状态、调度器与并发负载变化。应先预热查询、重复试验、记录环境并比较分布，不能把单次通知当作保证。

`pg_microbench_run(text)` 会使用调用者权限同步执行 SQL 字符串，其中可以包含写入或昂贵工作。若只有基准管理员应发起此类负载，就必须限制执行权。启用钩子范围也会给当前会话中匹配的查询增加测量开销。固定源码明确面向 Linux 和 2026 年 1 月时的 PostgreSQL 开发主干，应针对确切服务器源码构建测试，并没有广泛的已发布大版本兼容承诺。安装扩展会在会话中加载动态库，因此不需要预加载。
