## 用法

来源：

- [当前上游固定版本 README](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/README.md)
- [固定版本扩展控制文件](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/pg_query_state.control)
- [固定版本执行器与消息实现](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/pg_query_state.c)
- [PostgreSQL 18 自定义信号补丁](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/patches/custom_signals_18.0.patch)
- [PostgreSQL 18 运行时 EXPLAIN 补丁](https://github.com/postgrespro/pg_query_state/blob/cfda145588a770601426491471d8122ed868a6ce/patches/runtime_explain_18.0.patch)

pg_query_state 1.2 可提取另一个后端中正在运行查询的实时执行状态。它能返回嵌套查询栈帧和并行工作进程计划，以及当前行数、循环、计时和缓冲区统计，还提供数值与循环打印形式的进度估计。

### 补丁服务器的配置与查询

这不是即插即用的普通扩展。必须对精确 PostgreSQL 大版本应用上游两个补丁，重新构建并部署服务器，然后预加载模块并重启：

```conf
shared_preload_libraries = 'pg_query_state'
pg_query_state.enable = on
pg_query_state.enable_timing = off
pg_query_state.enable_buffers = off
```

```sql
CREATE EXTENSION pg_query_state;

SELECT pid, frame_number, query_text, plan, leader_pid
FROM pg_query_state(
    12345,
    format => 'json'
);

SELECT pg_progress_bar(12345);
```

应把 PID 换成当前正在执行查询的后端。调用者必须是超级用户，或者与目标后端具有相同有效角色。

### 侵入式监控边界

服务器补丁向 PostgreSQL 内部加入自定义进程信号和运行时 EXPLAIN 状态。预加载模块安装执行器钩子、分配共享内存、向领导进程与并行工作进程发信号，并通过共享内存队列传输查询文本和计划状态。启用时始终要求行数插桩，计时与缓冲区收集会再增加开销。每次小版本更新及扩展交互都必须验证完整的补丁服务器构建，而不只是共享库。

输出含实时查询文本、字面量、对象名、嵌套函数查询和计划细节。因此同角色访问可能看到另一个会话中携带数据的 SQL。应限制角色共享与函数执行权，保护监控输出，并验证超时和取消行为，避免无响应目标拖住监控会话。

进度是根据计划行数推导的启发式值；对于不可计数操作、估计不准、锁等待和许多写计划，它可能不可用或误导。当前仓库包含至 PostgreSQL 18 的补丁，并在 2026 年仍有活动，但这不代表补丁生产服务器风险很低。应先部署到专用诊断构建，并用真实并行与嵌套负载测量开销。
