## 用法

来源：

- [官方 README 与协议文档](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/README.rst)
- [扩展 SQL](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/sql/replisome.sql)
- [输出插件实现](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/src/replisome.c)

`replisome` 0.1 是预生产逻辑解码输出插件，以 JSON 流式发送 INSERT、UPDATE 和 DELETE 变更。它可以选择表、列与行，并提供可选 Python 消费框架。它是变更集成原型，不是完整复制方案。

### 配置逻辑解码

```conf
wal_level = logical
max_replication_slots = 1
max_wal_senders = 1
```

修改后重启，在 `pg_hba.conf` 中配置最小权限复制条目，安装输出插件库，并可选登记 SQL 扩展。SQL 扩展只暴露用于协议检查的 `replisome_version()`。

```sql
CREATE EXTENSION replisome;
SELECT replisome_version();

SELECT *
FROM pg_create_logical_replication_slot('integration_slot', 'replisome');
```

通过复制协议消费，例如：

```sh
pg_recvlogical -d appdb --slot integration_slot --start \
  -o pretty-print=1 -f -
```

每个已提交事务表示为 JSON，其中包含操作代码、模式/表、值，以及可用时的旧键数据。

### 过滤与输出选项

重复 `include` 和 `exclude` JSON 选项按顺序处理，最后一次匹配生效。include 可以按表/模式名或正则表达式匹配，选择/跳过列，并应用行级 `where` 表达式。其他选项包括 `include-xids`、`include-lsn`、`include-timestamp`、`include-schemas`、`include-types`、`include-empty-xacts` 和 `write-in-chunks`。分块输出在消费者重新组装之前可能不是有效 JSON。

### 可靠性与兼容性边界

复制槽会保留 WAL，直到消费者推进；应监控延迟和磁盘用量，并有计划地删除废弃槽。重连/重试可能重放数据，因此消费者必须幂等。过滤和省略列可能使输出不足以重建行，行过滤还必须针对 UPDATE/DELETE 键变化测试。

插件不处理 TRUNCATE、DDL、序列、模式演进或冲突解决。上游将其标为预生产，面向 PostgreSQL 9.4 时代逻辑解码，并记录 Python 2.7 消费者。旧私有服务器 API 在当前 PostgreSQL 上需要维护中的移植版。若没有完整协议、崩溃、升级、故障转移和重放测试，不能把它用于灾难恢复或无损复制。
