

## 用法

> [spock: PostgreSQL 多主逻辑复制扩展](https://github.com/pgEdge/spock)

PostgreSQL 15+ 的多主逻辑复制。每个节点同时充当发布者和订阅者。

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
-- 节点 1
SELECT spock.node_create(
    node_name := 'n1',
    dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);

-- 节点 2
SELECT spock.node_create(
    node_name := 'n2',
    dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);
```

### 创建订阅

对于多主，每个节点订阅其他所有节点：

```sql
-- 在 n1 上：订阅 n2
SELECT spock.sub_create(
    subscription_name := 'sub_n1n2',
    provider_dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);

-- 在 n2 上：订阅 n1
SELECT spock.sub_create(
    subscription_name := 'sub_n2n1',
    provider_dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);
```

### 复制集管理

```sql
-- 将表添加到复制
SELECT spock.repset_add_table('default', 'my_table');

-- 从复制中移除表
SELECT spock.repset_remove_table('default', 'my_table');

-- 添加模式中的所有表
SELECT spock.repset_add_all_tables('default', '{public}');
```

### 关键特性

- 多主（双活）复制
- 自动 DDL 复制
- 使用提交时间戳进行冲突检测和解决
- 行和列过滤
- 支持 PostgreSQL 15、16、17 和 18
- 表必须有主键且跨节点模式匹配
