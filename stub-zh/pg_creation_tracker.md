## 用法

来源：

- [README.md](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/README.md)
- [pg_creation_tracker--1.0.sql](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/pg_creation_tracker--1.0.sql)
- [pg_creation_tracker.control](https://github.com/12vishnu10/pg_creation_tracker/blob/1e5dc91a17b8802fabd58b127fb30b8a5d93462d/pg_creation_tracker.control)

`pg_creation_tracker` 1.0 是 SQL/PLpgSQL 扩展，用于记录表 DDL，并估算既有表和数据库的时间戳。GitHub 发布标记为 v1.0.0，实际 control 和 SQL 版本为 1.0。

### 核心用法

```sql
CREATE EXTENSION pg_creation_tracker;
SELECT schema_name, object_name, creation_time, recorded_time, detection_method
FROM timestamp_tracking.object_creation_log
ORDER BY id;
SELECT database_name, creation_time, detection_method
FROM timestamp_tracking.database_creation_info;
```

### 运行边界

安装需要超级用户，会创建固定的 `timestamp_tracking` 模式、事件触发器，并立即回填历史记录。不需要原生库或预加载。需要跟踪的数据库须分别安装；安装到模板库会影响之后从该模板复制的数据库。

`object_creation_log` 提供 `creation_time`、`recorded_time` 和 `detection_method`。`ddl_command_end` 触发器以 `statement_timestamp()` 记录表及外部表 DDL，同一对象 OID 只保留第一条记录，因此不是完整且不可变的审计日志。`backfill_existing_tables()` 先尝试目录事务提交时间，再读取文件元数据。`database_creation_info` 根据目录元数据推断集群内数据库的时间。

历史值只是估算：SQL 读取文件修改时间而非诞生时间；目录行的 `xmin` 可能对应后续变更，分区根表使用最早的子表修改时间。提交时间缺失或路径不可读时可能返回 NULL，不应将这些值作为最初创建时间的证明。辅助函数使用 `SECURITY DEFINER`，部署前须检查执行授权和搜索路径安全。所检查的发布版没有给出完整的大版本支持矩阵，也未提供明确的许可证文件。
