## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/README.md)
- [后台工作进程实现](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/src/worker.rs)
- [SQL 可调用函数](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/src/functions.rs)
- [扩展 control 文件](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/pg_walrus.control)

`pg_walrus` 监控请求检查点，并自动增减 `max_wal_size`。它必须加入 `shared_preload_libraries` 并重启；应在 `walrus.database` 指定的数据库中创建扩展，以保存调整历史。

```sql
CREATE EXTENSION pg_walrus;
ALTER SYSTEM SET walrus.dry_run = true;
SELECT pg_reload_conf();
SELECT jsonb_pretty(walrus.status());
SELECT walrus.recommendation();
```

应从 `walrus.dry_run = true` 开始。正常情况下，当阈值、上限、缩小、冷却与每小时限速规则决定调整时，工作进程会执行 `ALTER SYSTEM`、写入 `postgresql.auto.conf` 并发送 SIGHUP。`walrus.history` 记录决策；`walrus.analyze(apply := true)` 可立即应用一次，且要求超级用户。

改变 WAL 大小会影响检查点频率、磁盘占用、崩溃恢复、复制保留与备份操作。应在 WAL 峰值速率下验证算法，设置安全的 `walrus.max` 与 `walrus.min_size`，监控剩余空间，并与现有配置管理协调，避免自动改写与声明式设置相互冲突。
