## 用法

来源：

- [上游 README](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/README.md)
- [SQL 用法与运维指南](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/doc/USAGE.md)
- [扩展 control 文件](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/duckpipe-pg/pg_duckpipe.control)
- [扩展 Cargo 清单](https://github.com/relytcloud/pg_duckpipe/blob/86d5c248d80370fb2e10ddc1abd780fd94c3a52b/duckpipe-pg/Cargo.toml)

`pg_duckpipe` 使用逻辑解码复制堆表中的现有行，并把后续 WAL 变更持续写入 DuckLake 列式表。安装 `pg_ducklake` 及其捆绑的 `pg_duckdb` 后，配置逻辑 WAL 并预加载两个动态库：

```conf
wal_level = logical
shared_preload_libraries = 'pg_duckdb,pg_duckpipe'
```

重启 PostgreSQL，启用扩展，再把带主键的表加入默认同步组：

```sql
CREATE EXTENSION pg_duckdb;
CREATE EXTENSION pg_ducklake;
CREATE EXTENSION pg_duckpipe;

CREATE TABLE orders (
  id bigint PRIMARY KEY,
  total numeric NOT NULL
);

SELECT duckpipe.add_table('public.orders');
SELECT source_table, state, rows_synced, applied_lsn
FROM duckpipe.status();
```

### 兼容性与 WAL 安全

版本 `1.0.0` 的 Cargo 构建特性覆盖 PostgreSQL 14 至 18，而顶层 README 表示 CI 测试 PostgreSQL 17 和 18；应验证实际使用的软件包组合。默认 upsert 模式要求主键，append 模式则可以处理无主键表。每个同步组都会占用逻辑复制槽，停滞或遗留的同步组可能持续保留 WAL，直至磁盘耗尽。应配置 `max_slot_wal_keep_size`、监控复制槽延迟，并在停用空组时调用 `duckpipe.drop_group()`。
