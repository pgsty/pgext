## 用法

来源：

- [腾讯云 WAL 分析官方文档](https://cloud.tencent.com/document/product/409/130512)
- [腾讯云 PostgreSQL 官方产品页](https://cloud.tencent.com/product/postgres)

`tencentdb_wal_stat` 是腾讯云数据库 PostgreSQL 的诊断扩展，读取近期 WAL 文件并归因记录数、数据字节、全页镜像、资源管理器类型和关系标识。它用于定位与 WAL 突增相关的数据库、模式、表或索引，并不是长期 WAL 历史存储。

### 要求与核心流程

腾讯云文档说明版本 `1.0` 适用于 TencentDB PostgreSQL 15 或更高版本，只能由 `pg_tencentdb_superuser` 成员在主库执行：

```sql
CREATE EXTENSION IF NOT EXISTS tencentdb_wal_stat;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'tencentdb_wal_stat';
```

传入要检查的近期 WAL 文件数量，范围是 `1` 到 `500`。下面的查询找出最近 50 个仍可用文件中贡献最大的表和索引：

```sql
SELECT database_name,
       schema_name,
       relation_name,
       relation_kind,
       rmgr_name,
       wal_records,
       pg_size_pretty((wal_bytes + wal_fpi_bytes)::bigint) AS total_size
FROM tencentdb_wal_stat(50)
WHERE relation_kind IN ('table', 'index')
ORDER BY (wal_bytes + wal_fpi_bytes) DESC
LIMIT 20;
```

用 `wal_bytes + wal_fpi_bytes` 汇总归因总大小。按 `database_name`、`schema_name` 或 `rmgr_name` 分组，可以从实例级排查逐步下钻到数据库、模式或写入类型。

### 输出与解释

- `database_oid`、`schema_oid` 和 `relation_oid` 标识对象，配套名称列会在系统目录可用时解析。
- `relation_kind` 区分表、索引、TOAST、序列、物化视图、共享或系统对象以及未知对象。
- `rmgr_name` 按资源管理器分类 WAL，例如 `Heap`、`Btree` 或 `Transaction`。
- `wal_records` 与 `wal_bytes` 统计记录和不含 FPI 的字节，`wal_fpi` 与 `wal_fpi_bytes` 统计全页镜像及其字节。

### 运维边界

- 函数只能在主库使用，并要求服务商角色。它是腾讯云数据库组件，不能假定社区 PostgreSQL 中存在同一软件包或角色。
- 较大的 `wal_num` 会读取更多 WAL 并耗时更长。腾讯云建议大范围分析避开高峰；事件排查应从较小范围开始，并监控查询本身。
- 只能分析尚未回收的 WAL 文件。参数是文件数而不是时间区间，因此流量和 WAL 段大小会影响覆盖时长。
- 其他数据库中的对象名可能显示为 `<oid:N>` 或 `<filenode:N>`，因为当前数据库无法解析其系统目录。应保留 OID，并在正确数据库中关联名称，不要把占位符当作缺失 WAL。
- 归因是诊断证据，不是因果证明。调整索引、限流租户或修改 `max_wal_size` 前，应结合工作负载、检查点、复制槽与延迟、归档、磁盘使用和查询统计一起判断。
