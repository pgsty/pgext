## 用法

来源：

- [已复核 commit 的 pgduck_rs GOAL](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/GOAL.md)
- [已复核 commit 的 pgduck_rs.control](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/pgduck_rs.control)
- [Cargo 软件包元数据](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/Cargo.toml)
- [read_parquet 实现](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/read_parquet.rs)
- [扩展模块导出](https://github.com/qsliu2017/pgduck-rs/blob/5c64da8bc6b7496e909580545b34d718b8686311/src/lib.rs)

`pgduck_rs` 是一个探索性的 pgrx 扩展，会在 PostgreSQL 后端进程中嵌入 DuckDB。在已复核 commit 中，唯一的数据读取操作是 `read_parquet(path text)`，它以 `SETOF jsonb` 逐行返回 Parquet 数据；`hello_pgduck_rs()` 只是演示函数。

### 核心流程

以超级用户安装扩展，再传入数据库服务器上可见的文件路径：

```sql
CREATE EXTENSION pgduck_rs;

SELECT row
FROM read_parquet('/srv/postgres/import/events.parquet') AS t(row);
```

路径由服务器进程而非 SQL 客户端解析。PostgreSQL 操作系统账户必须有权读取该文件。

### 对象索引

- `read_parquet(path text) RETURNS SETOF jsonb`：通过内存 DuckDB 连接打开本地 Parquet 文件，每行输出一个 JSON 对象。
- `hello_pgduck_rs() RETURNS text`：返回固定问候语，不属于 Parquet 工作流。

### 运维边界

- 上游 GOAL 明确把它描述为“学习，不交付”的原型工作。仓库可访问并不代表可用于生产。
- 每次调用 `read_parquet` 都会新建一个内存 DuckDB 连接，并先将完整结果收集到 Rust 向量中。因此在返回任何一行之前，大结果就会占用后端内存。
- 已复核实现使用 DuckDB 的本地 `read_parquet(?)` 调用，不会加载 `httpfs`、创建 secret，也没有远程对象存储流程。应将路径参数视为服务器文件访问能力。
- DuckDB 在 PostgreSQL 后端进程内执行，因此解析器、文件解码器、内存分配或原生代码故障都与后端共享进程边界。
- control 文件要求超级用户安装，并将扩展标记为不可重定位。Cargo 和 control 文件均报告版本 `0.0.0`；为 PostgreSQL 13 至 18 声明的 pgrx 构建 feature 并不构成兼容性或支持保证。
