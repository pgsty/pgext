

## 用法

> [pg_drop_events: 记录 DROP TABLE、DROP COLUMN、DROP MATERIALIZED VIEW 语句的事务 ID](https://github.com/bolajiwahab/pg_drop_events)

`pg_drop_events` 扩展使用事件触发器自动记录对表、列和物化视图的 DROP 操作详情。记录的信息可用于意外删除后的时间点恢复（PITR）。

### 跟踪的操作

- `DROP TABLE`
- `DROP COLUMN`（通过 `ALTER TABLE`）
- `DROP MATERIALIZED VIEW`

### 记录的信息

| 列 | 描述 |
|--------|-------------|
| `pid` | 进程标识符 |
| `usename` | 执行命令的数据库用户 |
| `query` | SQL 语句 |
| `xact_id` | 事务标识符 |
| `wal_position` | 预写日志位置 |
| `objid` | 对象标识符 |
| `object_name` | 被删除对象的完全限定名称 |
| `object_type` | 对象分类（表、表列等） |
| `xact_time` | 事务时间戳 |

### 示例

```sql
CREATE EXTENSION pg_drop_events;

-- 删除一个表
DROP TABLE t.t3;
-- NOTICE: table t.t3 dropped by transaction 1085.

-- 查询事件日志
SELECT * FROM pg_drop_events;
```

### 时间点恢复

记录的数据直接映射到 PostgreSQL 恢复参数：

| pg_drop_events 列 | 恢复参数 |
|-----------------------|-------------------|
| `xact_id` | `recovery_target_xid` |
| `xact_time` | `recovery_target_time` |
| `wal_position` | `recovery_target_lsn` |

在 `postgresql.conf` 或 `recovery.conf` 中使用这些值将数据库恢复到意外删除之前的时间点。
