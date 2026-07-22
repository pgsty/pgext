## 用法

来源：

- [官方 README](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/README.md)
- [官方 PGXN 文档](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/doc/jsoncdc.md)
- [PostgreSQL 逻辑解码文档](https://www.postgresql.org/docs/current/logicaldecoding-explanation.html)

`jsoncdc` 版本 `0.1.0` 是把 PostgreSQL WAL 变更转换为换行分隔 JSON 的逻辑解码输出插件。它由逻辑复制槽加载，不需要 `CREATE EXTENSION`，也不安装 SQL API。

### 核心流程

先配置逻辑解码：`wal_level` 必须是 `logical`，且 `max_replication_slots` 必须为该复制槽留有容量。然后使用有权创建并消费逻辑复制槽的角色连接目标数据库。

```sql
SELECT *
FROM pg_create_logical_replication_slot('jsoncdc_slot', 'jsoncdc');

-- After transactions modify user tables:
SELECT *
FROM pg_logical_slot_get_changes('jsoncdc_slot', NULL, NULL);
```

数据流依次生成事务开始对象、表模式与名称记录、插入/更新/删除记录、可选的逻辑消息记录，以及包含事务 ID 与时间戳的提交对象。每个 JSON 对象占一行输出。

### 复制槽操作

```sql
SELECT slot_name, active, restart_lsn, confirmed_flush_lsn
FROM pg_replication_slots
WHERE slot_name = 'jsoncdc_slot';

SELECT pg_drop_replication_slot('jsoncdc_slot');
```

逻辑复制槽独立于消费者持久存在，并会保留所需 WAL 与目录行。应监控延迟，让消费者能够容忍崩溃后的重复投递，并在不再需要时删除复制槽。

### 兼容性边界

上游代码是年代较久的 Rust 输出插件，最后记录的版本为 `0.1.0`。采用前应确认它能针对准确的 PostgreSQL 服务器构建。DDL 不是行变更流，客户端必须按插件记录的顺序解析，而不能假设每个事务对应一个独立完整的 JSON 文档。
