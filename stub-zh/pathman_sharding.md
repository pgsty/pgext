## 用法

来源：

- [扩展 control 文件](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/pathman_sharding.control)
- [通用 SQL 对象](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/init.sql)
- [哈希分片 SQL](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/hash.sql)
- [范围分片 SQL](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/range.sql)
- [必需的 postgres_fdw 补丁](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/postgres_fdw_execute.patch)
- [远程 DDL 与索引复制实现](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/src/pathman_sharding.c)
- [ProcessUtility 钩子实现](https://github.com/funbringer/pathman_sharding/blob/0c1bb0cfbb512f6850eb87de382e235c1db7b967/src/hooks.c)

`pathman_sharding` 0.1 是 PostgreSQL 10 时代的历史性 `pg_pathman` 配套扩展。它创建 `postgres_fdw` 分区，并在远程 PostgreSQL 服务器上复制实体表与父表索引；它无法配合原生 `postgres_fdw` 工作，也不是现代内置分区方案。

### 非标准前置条件

需要安装 `pg_pathman`、`postgres_fdw` 和 `pathman_sharding`，并将 `pathman_sharding` 放到与 `pg_pathman` 相同的模式中。还必须将仓库补丁应用到 PostgreSQL 自带的 `postgres_fdw`；补丁会增加 `postgres_fdw_execute_custom_command(text,text)`，扩展使用它执行远程 DDL。标准 PostgreSQL 并不提供该函数。

每个操作角色都需要在每台外部服务器上拥有可用的用户映射。补丁函数会通过当前 PostgreSQL 角色的映射执行命令，因此安装成功并不能证明远程创建或删除可用。

```sql
CREATE EXTENSION pg_pathman WITH SCHEMA public;
CREATE EXTENSION postgres_fdw WITH SCHEMA public;
CREATE EXTENSION pathman_sharding WITH SCHEMA public;

-- Assumes shard_a and shard_b are configured postgres_fdw servers.
CREATE TABLE events (id bigint NOT NULL, payload jsonb);

SELECT public.create_foreign_hash_partitions(
  'events'::regclass,
  'id',
  4,
  ARRAY['shard_a', 'shard_b'],
  false
);

SELECT partition, server
FROM public.pathman_foreign_partition_list;
```

只有在验证数据移动与容量后，才能将 `partition_data` 参数设为 true；false 会创建分区，但保留父表中的现有数据。

### 主要对象

- `create_foreign_hash_partitions` 创建哈希分区，默认以轮转方式将分区名分配到外部服务器。
- `add_foreign_range_partition` 为已配置范围分区的父表新增一个外部分区。
- `distribute_partitions_among_servers` 是默认分配过程，可通过哈希创建函数的过程参数替换。
- `pathman_foreign_partition_list` 将 pg_pathman 元数据与外部表及服务器元数据连接起来。

### 远程 DDL、权限与兼容性边界

ProcessUtility 钩子会响应 pg_pathman 外部分区的创建与删除。创建时会发送远程表 DDL，然后尝试复制每个父表索引；索引复制错误会被捕获并忽略。删除时会发送远程 `DROP TABLE ... CASCADE`。每次操作后都应比对本地与远程的表和索引，并协调每个分片的备份与恢复。

SQL 向公众授予 `pathman_foreign_partition_list` 读取权，其中包括外部服务器 `srvoptions`。postgres_fdw 补丁创建原始命令函数时没有显式限制权限。应撤销 `PUBLIC` 对 `postgres_fdw_execute_custom_command(text,text)` 和分片辅助函数的 `EXECUTE`，仅授予严格受控的管理角色。

上游没有 README、发行历史、升级脚本或当前兼容矩阵。源码使用 PostgreSQL 10 的 utility hook 和 relation API，目录也只标识 PostgreSQL 10；不能假设它可在更新大版本上构建。考虑维护遗留系统前，应在隔离集群中测试打补丁的服务器构建、钩子共存、角色映射、远程 DDL 失败、部分索引创建、路由、数据移动和拆除。
