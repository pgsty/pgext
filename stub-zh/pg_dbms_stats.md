## 用法

来源：

- [14.0 版手册](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/pg_dbms_stats-en.md)
- [14.0 版对象参考](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/objects-en.md)
- [扩展控制文件](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/pg_dbms_stats.control)

`pg_dbms_stats` 14.0 可以保存、恢复和锁定规划器统计信息，让 PostgreSQL 14 优化器持续看到选定的统计快照。它用于减少 `ANALYZE` 或数据分布漂移引起的计划变化；它不会固定执行计划，也不会冻结表数据。

### 启用扩展

在库文件会拦截规划器目录读取的每个数据库中安装并注册 `pg_dbms_stats`。在没有 `dbms_stats` 模式的数据库中加载库可能让普通语句失败，因此要先完成逐库注册，再启用集群级预加载。

```sql
CREATE EXTENSION pg_dbms_stats;

SHOW pg_dbms_stats.use_locked_stats;
```

若要自动加载，应将其加入 `shared_preload_libraries` 并重启 PostgreSQL。

```conf
shared_preload_libraries = 'pg_dbms_stats'
```

做受控会话测试时，上游手册也允许在该数据库已创建扩展后执行 `LOAD 'pg_dbms_stats'`。

### 备份、恢复与锁定

```sql
SELECT dbms_stats.backup_database_stats('before scheduled analyze');

ANALYZE;

SELECT dbms_stats.restore_database_stats(now() - interval '1 day');

SELECT *
FROM dbms_stats.backup_history
ORDER BY time DESC;

SET pg_dbms_stats.use_locked_stats = off;
EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SET pg_dbms_stats.use_locked_stats = on;
```

恢复操作也会锁定恢复出的统计信息。备份、恢复、锁定、解锁和导入都提供数据库、模式、表和列级变体。目标对象必须先有 PostgreSQL 统计信息，备份或锁定才有实际意义。

### 重要对象

- `dbms_stats.backup_history` 记录备份集。
- `dbms_stats.relation_stats_effective` 和 `dbms_stats.column_stats_effective` 展示当前提供给规划器的值。
- `dbms_stats.stats` 汇总锁定的统计信息。
- `dbms_stats.clean_up_stats()` 删除表或列被移除后遗留的锁定行。
- 导出脚本和导入函数通过服务端文件搬运统计信息；导入要求 PostgreSQL 服务账号可读的绝对路径。

### 运维边界

锁定统计信息不能保证计划不变：索引、模式、规划器 GUC、PostgreSQL 代码、参数值和表大小仍会影响规划。应监控查询计划，并在固定快照带来损害时迅速解锁。

删除对象不会自动清理其锁定行，因此应定期执行 `dbms_stats.clean_up_stats()`。包含多态 `anyarray` 值的统计信息需要遵循手册中的特殊转储和恢复流程。14.0 是 PostgreSQL 14 分支；应使用匹配的上游分支并对升级做回归测试，不能加载到其他大版本。
