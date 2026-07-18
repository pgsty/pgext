## 用法

来源：

- [已复核 commit 的 pgduck_rs GOAL](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/GOAL.md)
- [已复核 commit 的 pgduck_rs.control](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/pgduck_rs.control)
- [Cargo 软件包元数据](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/Cargo.toml)
- [read_parquet 实现](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/read_parquet.rs)
- [扩展模块导出](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/lib.rs)

`pgduck_rs` 是一个探索性的 pgrx 扩展，用于从 PostgreSQL 调用嵌入式 DuckDB。在已复核 commit 中，已实现的数据操作是 `read_parquet(path text)`，它以 `SETOF jsonb` 返回每一行 Parquet 数据；模块还导出了演示函数 `hello_pgduck_rs`。

### 读取本地 Parquet 文件

```sql
CREATE EXTENSION pgduck_rs;
SELECT row
FROM read_parquet('/var/lib/postgresql/data/events.parquet') AS t(row);
```

路径在数据库服务器上解析，PostgreSQL 操作系统账户必须有读取权限。

### 注意事项

- 上游 GOAL 明确把项目描述为“学习，不交付”的原型工作。目录生命周期 `active` 只表示仓库可访问，不能覆盖上游对成熟度的表述。
- 每次调用 `read_parquet` 都会新建一个内存 DuckDB 连接，并在返回行之前物化完整结果。大文件可能消耗大量后端内存。
- 已复核代码只支持本地文件，不会加载 `httpfs`，也不支持凭据、secret 或远程对象存储。应执行常规服务器文件访问控制。
- Cargo 和控制文件都报告版本 `0.0.0`。pgrx feature 列出了 PostgreSQL 13 至 18，但这并不构成生产兼容性保证。
