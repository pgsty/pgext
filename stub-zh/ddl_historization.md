## 用法

来源：[README](https://github.com/rodo/pg_ddl_historization/blob/master/README.md)，[releases](https://github.com/rodo/pg_ddl_historization/releases)

`ddl_historization` 是一个 PostgreSQL 扩展，用于把数据库中的 DDL 变更记录到历史化表中。上游 README 说明可通过 `make install`、`pgxn install ddl_historization`，以及 AWS RDS 上的 `pg_tle` 路径完成安装。

### 启用日志

```sql
CREATE EXTENSION ddl_historization;
```

README 将该扩展描述为使用 PostgreSQL event trigger 对数据库中的 DDL 变更做历史化记录。

### 上游当前文档

- 集群本地安装：`make install`
- PGXN 安装：`pgxn install ddl_historization`
- AWS RDS / `pg_tle`：用 `make pgtle` 构建 `pgtle.ddl_historization-0.3.sql`
- 测试套件：使用 pgTAP 的 `make test`

### 值得关注的发布说明

- `0.2` 是这次刷新任务要求对应的版本。
- `0.0.4` 的发布说明提到新增了启动和停止日志记录的函数。
- `0.0.6` 的发布说明提到新增了 `ddl_history_column` 表。
- `0.0.7` 的发布说明提到修复了与外键相关的日志记录问题。

### 注意事项

当前上游 README 仍较简略，没有给出 start/stop logging 函数的精确 SQL 签名，也没有完整说明后续版本新增历史化表的 schema。除非 README 或发布说明更明确，否则这里应保持保守描述。
