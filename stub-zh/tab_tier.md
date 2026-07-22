## 用法

来源：

- [PGXN 官方使用文档](https://pgxn.org/dist/tab_tier/doc/tab_tier.html)
- [PGXN 官方发行页](https://pgxn.org/dist/tab_tier/)

`tab_tier` 是一个历史 PL/pgSQL 分区维护与归档工具包。它不同于触发器路由的分区：近期行留在根表，计划任务随后把旧行移入继承子表，可选的长期存储可以是本地表或外部表。

```sql
CREATE EXTENSION tab_tier;
SELECT tab_tier.register_tier_root('comm', 'yell', 'created_dt');
SELECT tab_tier.bootstrap_tier_parts('comm', 'yell');
SELECT tab_tier.migrate_tier_data('comm', 'yell', '201301');
```

应以高于 `part_period` 的频率调度 `cap_tier_partitions()`，再运行 `migrate_all_tiers()`，确保移动数据前已有符合条件的目标。默认设计在根表保留 `root_retain` 范围的数据，用 `part_period` 划分子表边界，并在 `lts_threshold` 后归档；可通过 `tier_root`、`tier_part` 和 `tier_config` 检查实际状态。

迁移先复制行，再从根表删除。大表引导应使用可恢复的逐分区事务：`flush_tier_data()` 和 `flush_all_tiers()` 可能在一个巨大事务中移动所有符合条件的行，任何错误都会回滚这些工作。应测试锁、WAL、复制延迟、磁盘余量、约束、索引、触发器、迟到或未来时间戳以及并发写入。

归档可能把数据移到 `lts_target` 并删除旧子表，启用前必须核对行数和持久备份。该扩展使用旧式继承分区，通过 `tab_tier_role` 授予较宽管理权限，并带有 2014 年时代假设；新系统应优先采用受维护的声明式分区工具。
