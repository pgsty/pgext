## 用法

来源：

- [上游使用文档](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/doc/shard_manager.md)
- [0.0.1 版本 SQL 实现](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/sql/shard_manager.sql)
- [扩展控制文件](https://github.com/bonesmoses/shard_manager/blob/05989f903c72874285f751704988fd0111635d82/shard_manager.control)

`shard_manager` 是一个已废弃的 PL/pgSQL 工具包，用于每分片一个模式的布局，以及 64 位时间/分片/序列 ID。它把模板记录在 `shard_table`，物理映射记录在 `shard_map`，生成器设置记录在 `shard_config`；它不会路由查询，也不会创建远程外部表。

```sql
CREATE SCHEMA shard;
CREATE EXTENSION shard_manager WITH SCHEMA shard;
SELECT shard.register_base_table('comm', 'yell', 'id');
SELECT shard.create_next_shard('comm', 'localhost');
SELECT shard.init_shard_tables('comm', 1);
```

必须在初始化任何分片之前配置 `epoch`、`shard_count` 和 `ids_per_ms`。扩展会把后两者向下调整为 2 的幂，并在初始化后拒绝相关修改，因为改变位分配可能产生冲突。应用必须自行维持 `shard_map` 的全局一致性并完成路由。

管理函数使用 `SECURITY DEFINER`，会创建模式、序列、表、默认值和授权。部分动态 SQL 拼接模式名时没有标识符引用。只能允许严格受控的标识符和角色，应审计每个定义者函数及搜索路径，并且不要把管理辅助函数授予不可信用户。

使用 `LIKE ... INCLUDING ALL` 克隆表并不构成完整的分布式模式或迁移系统。依赖这一历史设计之前，应测试约束、序列、索引、触发器、所有权、DDL 发布、并发分片创建、故障切换、时钟行为、ID 耗尽、备份恢复及应用路由。
