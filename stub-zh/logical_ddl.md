## 用法

来源：[README](https://github.com/samedyildirim/logical_ddl/blob/master/README.md)

`logical_ddl` 会捕获一组受限的 `ALTER TABLE` 变更，将其写入可复制的 shadow table，并在 logical replication 的 subscriber 上回放等价 DDL。

### 支持的 DDL

- `ALTER TABLE ... RENAME TO ...`
- `ALTER TABLE ... RENAME COLUMN ... TO ...`
- `ALTER TABLE ... ADD COLUMN ...`
- `ALTER TABLE ... ALTER COLUMN ... TYPE ...`
- `ALTER TABLE ... DROP COLUMN ...`

内建类型、arrays、composite types、domains 与 enums 可以作为列类型使用，但该扩展不会复制这些自定义类型本身的定义。

### 发布端与订阅端设置

```sql
CREATE EXTENSION logical_ddl;

-- Publisher
INSERT INTO logical_ddl.settings VALUES (true, 'source1');
INSERT INTO logical_ddl.publish_tablelist (relid) VALUES ('table1'::regclass);
ALTER PUBLICATION log_pub_1 ADD TABLE logical_ddl.shadow_table;

-- Subscriber
INSERT INTO logical_ddl.settings VALUES (false, 'source1');
INSERT INTO logical_ddl.subscribe_tablelist (source, relid)
VALUES ('source1', 'table1'::regclass);
ALTER SUBSCRIPTION log_sub_1 REFRESH PUBLICATION;
```

### 主要表

- `logical_ddl.settings`：声明当前节点是 publisher 还是 subscriber，以及 source name。
- `logical_ddl.publish_tablelist`：定义要捕获的表与命令类型。
- `logical_ddl.subscribe_tablelist`：定义要回放的目标表与命令类型。
- `logical_ddl.shadow_table`：复制用的命令日志。
- `logical_ddl.applied_commands`：回放历史与失败跟踪。

### 注意事项

- 该扩展在 superuser 权限下工作。
- 类型变更的 `USING` 表达式、default expressions、constraints 与 indexes 均未实现。
- Pigsty 提到上游在 `0.1.0` 修复了一个 `RAISE WARNING` 拼写错误，但这不影响这里记录的用户用法。
