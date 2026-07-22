## 用法

来源：

- [上游参考手册](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/doc.pdf)
- [扩展控制文件](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/cassandra_fdw.control)
- [官方上游仓库](https://github.com/pgsql-io/cassandra_fdw)

`cassandra_fdw` 通过 Cassandra C/C++ 驱动把 Cassandra 3+ 数据暴露为 PostgreSQL 外部表。权威服务器选项包括必需的 `host`（可用逗号分隔多个地址）、`port`（默认 9042）和 `protocol`（文档默认 4）。

```sql
CREATE EXTENSION cassandra_fdw;
CREATE SERVER cassandra_server
  FOREIGN DATA WRAPPER cassandra_fdw
  OPTIONS (host '10.0.0.10,10.0.0.11', port '9042', protocol '4');

CREATE USER MAPPING FOR app_user SERVER cassandra_server;
```

应根据已安装版本精确映射 Cassandra keyspace、表、列和类型来定义外部表。所复核的 2016 年参考文档没有说明下推、写入、认证、TLS、一致性级别、分页和超时支持，因此必须从精确 SQL 和 C 驱动构建验证每项能力，不能假定普通 PostgreSQL 语义。

远程读取不参与 PostgreSQL MVCC 快照或跨系统原子事务。应规划 Cassandra 一致性、模式变化、部分失败、重试、重复副作用、驱动 ABI、节点发现及网络延迟。构建支持时应把凭据放入受限用户映射，限制服务器使用权，并在生产使用前通过 `EXPLAIN` 对谓词下推做基准验证。
