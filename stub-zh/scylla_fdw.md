## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/README.md)
- [1.0 版本 SQL 对象](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw--1.0.sql)
- [扩展 control 文件](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.control)
- [ScyllaDB 驱动依赖](https://github.com/scylladb/cpp-rs-driver)

`scylla_fdw` 通过 PostgreSQL 外部数据包装器 API 暴露 ScyllaDB 表。它依赖单独构建的 ScyllaDB C/C++ 驱动，支持读写、简单谓词下推、TLS、连接池和 `IMPORT FOREIGN SCHEMA`。

```sql
CREATE EXTENSION scylla_fdw;
CREATE SERVER scylla_server
  FOREIGN DATA WRAPPER scylla_fdw
  OPTIONS (host '127.0.0.1', port '9042', consistency 'local_quorum');
CREATE USER MAPPING FOR CURRENT_USER
  SERVER scylla_server
  OPTIONS (username 'cassandra', password 'replace-me');
```

外部表必须设置 `keyspace` 与 `table` 选项。执行 `UPDATE` 和 `DELETE` 时还要添加 `primary_key`；ScyllaDB 通常也要求按分区键等值过滤才能高效访问。应限制目录访问并使用专用映射角色，避免凭据对过多用户可见。

该项目仍属 alpha：唯一发行标签为 0.1.0-alpha，而 control 版本是 1.0。它不下推连接或聚合，仅简单条件可下推，集合类型支持也有限。任何非测试用途之前都必须验证查询计划和故障行为。
