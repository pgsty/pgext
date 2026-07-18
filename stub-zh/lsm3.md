## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/README.md)
- [1.0 版本 SQL 对象](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3--1.0.sql)
- [C 实现](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3.c)
- [扩展 control 文件](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3.control)

`lsm3` 使用三个普通 B-tree 结构实现类似 LSM 的索引访问方法：活动顶层索引、合并顶层索引与基础索引。创建 SQL 扩展前，必须将其加入 `shared_preload_libraries` 并重启。

```sql
CREATE EXTENSION lsm3;
CREATE TABLE event_log (id bigint, payload text);
CREATE INDEX event_log_id_lsm ON event_log USING lsm3 (id);
```

每个 LSM3 索引都会启动独立合并后台进程。应相应规划 `max_worker_processes`、`lsm3.max_indexes` 与 `lsm3.top_index_size`，并评估扫描合并三个索引带来的读放大。

它不支持并行扫描与数组键。最重要的是，LSM3 无法强制唯一性：索引选项 `unique=true` 只是查询优化，不会拒绝重复值，绝不能用作唯一约束。1.0 没有当前兼容矩阵；应使用可丢弃数据验证崩溃恢复、WAL 重放、vacuum、并发合并、重建索引与升级行为。
