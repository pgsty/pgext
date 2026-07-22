## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/README.md)
- [1.0 版本 SQL 对象](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw--1.0.sql)
- [扩展 control 文件](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.control)
- [FDW 实现](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.c)
- [连接实现](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_connection.cpp)
- [ScyllaDB 驱动依赖](https://github.com/scylladb/cpp-rs-driver)

`scylla_fdw` 通过外部数据包装器 API 将 ScyllaDB 表映射到 PostgreSQL。已复核代码实现扫描、插入、更新、删除和 `IMPORT FOREIGN SCHEMA`，并支持有限的谓词下推。它需要 PostgreSQL 14 或更高版本、C++17 工具链，以及单独构建的 ScyllaDB C/C++ 驱动。

### 核心流程

创建服务器和按角色限定的用户映射，再明确声明远端表结构：

```sql
CREATE EXTENSION scylla_fdw;

CREATE SERVER scylla_server
  FOREIGN DATA WRAPPER scylla_fdw
  OPTIONS (host '127.0.0.1', port '9042', consistency 'local_quorum');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER scylla_server
  OPTIONS (username 'app_reader', password 'replace-me');

CREATE FOREIGN TABLE users (
  user_id uuid,
  username text,
  created_at timestamptz
)
SERVER scylla_server
OPTIONS (keyspace 'app', table 'users', primary_key 'user_id');

EXPLAIN (VERBOSE)
SELECT * FROM users
WHERE user_id = '00000000-0000-0000-0000-000000000001';
```

`keyspace` 与 `table` 标识远端关系。更新或删除需要构造远端键谓词时，应设置 `primary_key`。务必确认 ScyllaDB 实际的分区键和聚簇键要求，不要假定该选项能让任意谓词变得高效。

### 对象与选项索引

- `scylla_fdw_handler()` 和 `scylla_fdw_validator(text[], oid)` 为 `scylla_fdw` 外部数据包装器提供底层实现。
- 服务器选项包括 `host`、`port`、`consistency`、连接池大小以及 TLS 相关值。
- 用户映射选项是 `username` 与 `password`。
- 外部表选项包括 `keyspace`、`table` 和 `primary_key`。
- `IMPORT FOREIGN SCHEMA` 根据 ScyllaDB 元数据生成 PostgreSQL 外部表定义。

### 查询与事务边界

- 反解析器仅下推部分简单条件。连接与聚合不会下推；含有 `OR`、`NOT` 或 `NULL` 的表达式可能留在本地。请用 `EXPLAIN (VERBOSE)` 验证每条重要查询。
- 集合类型处理能力有限。应使用有代表性的数据测试每种映射类型及其空值行为。
- 远端写入通过驱动执行，不会加入 PostgreSQL 的事务或回滚机制。因此，后续本地错误可能发生在 ScyllaDB 变更已经提交之后。
- 源码接受 `request_timeout` 选项，但已复核的连接实现没有将它应用到驱动。不要依赖此设置强制查询期限。
- README 将 `ssl_cert`、`ssl_key` 与 `ssl_ca` 描述为文件路径，但已复核的连接代码并不读取文件，而是把选项字符串直接传给驱动的证书与密钥设置函数。使用前须用确切构建和选项表示验证 TLS。
- 已复核的 `IMPORT FOREIGN SCHEMA` 路径会建立独立的默认连接，不会把配置的 TLS、超时或一致性值带入该元数据请求。尤其对于仅允许 TLS 的集群，应单独测试。

### 运维注意事项

- 凭据保存在 PostgreSQL 用户映射目录中。应限制目录可见性，并只向专用角色授予服务器使用权限。
- 仓库唯一的发行标签是 `0.1.0-alpha`，而扩展 control 文件报告版本 `1.0`。在完成验证前，应将 API 和磁盘元数据都视为 alpha 质量。
- 驱动调用在 PostgreSQL 后端进程内运行。任何非测试部署前，都应演练 DNS 故障、认证错误、超时、远端节点丢失、取消和后端重启行为。
