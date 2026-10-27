## 用法

来源：

- [README](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/README.md)
- [C implementation](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/pg_query_plan.c)
- [SQL installation](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/pg_query_plan--1.0.sql)
- [Makefile](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/Makefile)

`pg_query_plan` 是 PostgreSQL 18 模块，用于查看其他后端正在执行的计划和实际行数。核对的源码带有版本 1.0 SQL，但没有控制文件，也不通过 CREATE EXTENSION 安装。README 标注的发布日期尚在未来，此条目记录当前源码预览。

### 启用与查看

将 pg_query_plan 加入 shared_preload_libraries 并重启 PostgreSQL，再在每个需要检查函数的数据库中运行随附的 pg_query_plan--1.0.sql。应保留已有预加载项。SQL 脚本会删除并重建 public.pg_query_plan(integer)，执行前应检查已有函数及其依赖。

```sql
SELECT pid FROM pg_stat_activity WHERE pid <> pg_backend_pid();
SELECT * FROM public.pg_query_plan(pg_backend_pid());
```

检查其他后端时，应把从 pg_stat_activity 选出的 PID 作为参数。调用者必须是超级用户或目标角色的成员。结果包括数据库、leader/worker 身份、查询 ID、SQL 文本和实时计划；空闲目标不返回行。

### 运行边界

模块为每条查询启用行数测量，存在额外开销。它发送 SIGUSR2 请求，在执行器安全点处理；单个计划节点中的长操作可能延迟响应，超时为五秒。可选 query_progress.py 客户端需要 Python 和 psycopg2，进度估算取决于规划器估计。没有后台工作进程或自动保留机制；当前仓库未能确立许可证条款。
