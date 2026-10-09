## 用法

来源：

- [README.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/README.md)
- [sql/pg_stat_statement_context--1.0.sql](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/sql/pg_stat_statement_context--1.0.sql)
- [docs/sql-interface.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/sql-interface.md)
- [docs/limitations.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/limitations.md)
- [docs/configuration.md](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/docs/configuration.md)
- [pg_stat_statement_context.control](https://github.com/benchub/pg_stat_statement_context/blob/63c012c63e9a8ec1ebbb5271ba8c44f0b53e0fc1/pg_stat_statement_context.control)

`pg_stat_statement_context` 在滚动时间桶中按查询身份、用户、数据库与应用标签记录调用次数和执行时间。支持 PostgreSQL 14–18，可配合 `pg_stat_statements` 使用。

### 核心用法

```ini
shared_preload_libraries = 'pg_stat_statements, pg_stat_statement_context'
```

```sql
CREATE EXTENSION pg_stat_statement_context;
SELECT 1 /*controller:users,action:show*/;
SELECT queryid, tags, calls, total_exec_time
FROM pg_stat_statement_context_totals ORDER BY calls DESC;
```

### 运行边界

将下方配置合并到现有预加载列表并重启；使用 `pg_stat_statements` 时必须将其放在此库前面，以保留工具语句的查询 ID。预加载后会采集整个集群，SQL 对象按数据库安装。`pg_stat_statement_context_totals` 汇总时间桶，`pg_stat_statement_context_activity` 展示当前标签，`pg_stat_statement_context_info` 提供容量与跳过统计。重置和提取辅助函数默认仅允许超级用户调用，查看其他用户标签需要 `pg_read_all_stats` 的权限。标签应保持低基数并避免个人信息。预备语句会保留准备时的注释，可用 `pg_stat_statement_context.tags_override` 提供执行时的请求上下文。统计容量有限且可能被淘汰，失败语句和没有查询 ID 的语句不能构成完整审计记录。
