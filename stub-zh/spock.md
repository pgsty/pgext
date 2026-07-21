## 用法

来源：

- [Spock v5.0.10 README](https://github.com/pgEdge/spock/blob/v5.0.10/README.md)
- [Getting started](https://github.com/pgEdge/spock/blob/v5.0.10/docs/getting_started.md)
- [Configuration reference](https://github.com/pgEdge/spock/blob/v5.0.10/docs/configuring.md)
- [Limitations](https://github.com/pgEdge/spock/blob/v5.0.10/docs/limitations.md)
- [Release notes](https://github.com/pgEdge/spock/blob/v5.0.10/docs/spock_release_notes.md)

`spock` 为 PostgreSQL 15 至 18 提供了主动-主动逻辑复制。每个参与的数据库都是一个 Spock 节点；通过在节点之间创建有向订阅，形成一个多主拓扑结构。

### 配置

在 `postgresql.conf` 中：

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'spock'
track_commit_timestamp = on
spock.enable_ddl_replication = on
spock.include_ddl_repset = on
```

### 启用

```sql
CREATE EXTENSION spock;
```

### 创建节点

在每个节点上创建节点身份：

```sql
-- Node 1
SELECT spock.node_create(
    node_name := 'n1',
    dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);

-- Node 2
SELECT spock.node_create(
    node_name := 'n2',
    dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);
```

### 创建订阅

对于多主模式，每个节点订阅其他所有节点：

```sql
-- On n1: subscribe to n2
SELECT spock.sub_create(
    subscription_name := 'sub_n1n2',
    provider_dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);

-- On n2: subscribe to n1
SELECT spock.sub_create(
    subscription_name := 'sub_n2n1',
    provider_dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);
```

### 复制集管理

```sql
-- Add table to replication
SELECT spock.repset_add_table('default', 'my_table');

-- Remove table from replication
SELECT spock.repset_remove_table('default', 'my_table');

-- Add all tables in a schema
SELECT spock.repset_add_all_tables('default', '{public}');
```

### 关键功能

- 多主（主动-主动）复制
- 自动 DDL 复制
- 使用提交时间戳进行冲突检测和解决
- 行和列过滤
- 支持 PostgreSQL 15、16、17 和 18
- 表必须在所有节点上具有主键且模式匹配

### 操作与注意事项

- 在创建节点或订阅之前，在每个参与的服务器上安装 `spock` 并将其添加到 `shared_preload_libraries` 中。
- 确保各节点上的表定义、数据类型、主键和相关唯一索引一致。即使启用了 DDL 复制，也要协调 DDL 操作。
- 被复制的表需要具有主键或其他可用的副本标识符。临时表和未记录表不是复制目标。
- Spock 是按数据库操作的。对于每个参与的数据库，重复扩展名和拓扑结构设置。
- 主动-主动冲突处理依赖于提交时间戳和策略。在生产使用前，请测试同时插入和更新操作，特别是空值唯一键。
- 上游文档中在 README 中列出了平台/构建要求；请验证每个节点上使用的 PostgreSQL 构建和 Spock 包是兼容的。

### 版本 5.0.10

`5.0.10` 是 5.0 系列的一个补丁版本。其发布说明包括修复了包含 `NULL` 的唯一索引、`NULLS NOT DISTINCT` 冲突处理、在索引删除后刷新缓存的索引元数据、异常路径内存处理以及用于滚动补丁升级期间使用的数值版本检查。将每个节点升级到兼容的补丁级别，并在滚动更改后验证订阅。
