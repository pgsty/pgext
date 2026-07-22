## 用法

来源：

- [官方 README](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/README.md)
- [扩展控制文件](https://github.com/pierreforstmann/pg_logqueryid/blob/306da69ba0dc73f49ab08feaf61b83e5f521cdd1/pg_logqueryid.control)
- [PGXN 发行版与版本](https://pgxn.org/dist/pg_logqueryid/)

`pg_logqueryid` 会把 `pg_stat_statements` 查询标识符写到 `auto_explain` 输出的计划旁边，使运维人员能将日志计划与规范化语句统计关联起来。它没有添加可查询的 SQL API，也没有自己的自定义 GUC。

### 核心流程

加载并配置 `pg_stat_statements` 与 `auto_explain`，启用日志收集，再在会话中或通过预加载设置加载 `pg_logqueryid`。

```conf
logging_collector = on
shared_preload_libraries = 'pg_stat_statements,auto_explain,pg_logqueryid'
pg_stat_statements.track = 'all'
auto_explain.log_min_duration = '1s'
```

```sql
SELECT queryid, query
FROM pg_stat_statements
WHERE queryid = 5917340101676597114;
```

也可以通过 `LOAD 'pg_logqueryid'` 或 `session_preload_libraries` 只在选定连接中启用模块。如果两个前置模块未加载并配置，它会保持不活动状态。

### 版本与日志边界

目录中的控制文件报告 `0.0.1`，而 PGXN 发布了直到 `1.0.1` 的后续发行版；应固定实际安装的制品，不能从仓库 README 推断其版本。上游报告已验证到 PostgreSQL 18，但钩子交互与日志格式仍需在准确构建上测试。

在 PostgreSQL 16 及以后版本中，启用 `auto_explain.log_verbose` 后，`auto_explain` 已能直接打印查询 ID，因此该扩展可能是冗余的。用零时长阈值记录每个计划会带来很高开销，也可能暴露 SQL 文本或敏感值。应从选择性阈值与采样策略开始，保护日志访问和保留周期，并验证日志查询 ID 与目标 `pg_stat_statements` 条目一致。
