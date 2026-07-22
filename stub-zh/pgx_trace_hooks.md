## 用法

来源：

- [官方仓库 README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/pgx_trace_hooks.control)
- [跟踪 hook 实现](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_trace_hooks/src/lib.rs)

`pgx_trace_hooks` 是一个诊断型预加载扩展，用于记录 PostgreSQL 执行器与工具命令 hook 的进入事件。它适合在受控测试服务器中观察语句是否进入 `ExecutorStart`、`ExecutorRun`、`ExecutorFinish`、`ExecutorEnd` 和 `ProcessUtility`；它记录的是 hook 阶段名称，而不是查询计划、耗时或语句归属。

### 核心流程

该库拒绝普通的运行时加载，因此需要在 `shared_preload_libraries` 中配置，并在创建 SQL 扩展前重启 PostgreSQL：

```conf
shared_preload_libraries = 'pgx_trace_hooks'
```

```sql
CREATE EXTENSION pgx_trace_hooks;

SELECT 1;
CREATE TEMP TABLE trace_probe(id integer);
```

执行有代表性的语句后检查 PostgreSQL 服务器日志。查询可能产生执行器阶段记录，工具命令则可能产生 `ProcessUtility`。对于每个 hook，实现都会保留并调用此前安装的回调；如果没有此前回调，则退回 PostgreSQL 的标准处理函数。

### 跟踪的 Hook

- `ExecutorStart` 标记执行器初始化。
- `ExecutorRun` 标记产生元组的执行过程。
- `ExecutorFinish` 标记执行器的完成工作。
- `ExecutorEnd` 标记执行器清理。
- `ProcessUtility` 标记工具语句处理。

### 运维说明

`pgx_trace_hooks` 不可重定位，创建时需要超级用户权限。除非通过 `shared_preload_libraries` 处理该库，否则 `_PG_init` 会报错；修改该设置需要重启服务器。Hook 消息使用 PostgreSQL 的 `INFO` 级别，可能显著放大普通工作负载的日志量，因此除非已经明确评估日志成本和数据暴露风险，否则不要把它用于繁忙的生产系统。

这些 hook 会串接本库初始化时已经存在的回调。因此，预加载顺序可能影响最终调用链；应按照最终顺序与所有其他使用 hook 的扩展一起测试。审阅的源码属于 PostgreSQL 15 时期的研究仓库，除加载库之外不提供 SQL API。
