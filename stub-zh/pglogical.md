

## 用法

> [pglogical: PostgreSQL 逻辑复制](https://github.com/2ndQuadrant/pglogical)

使用发布/订阅模型的 PostgreSQL 逻辑复制系统。无需触发器或外部程序。

### 启用

添加到 `postgresql.conf`：

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'pglogical'
```

```sql
CREATE EXTENSION pglogical;
```

### 提供者（发布者）设置

```sql
-- 在提供者上创建节点
SELECT pglogical.create_node(
    node_name := 'provider1',
    dsn := 'host=providerhost port=5432 dbname=mydb'
);

-- 将 public 模式中所有表添加到默认复制集
SELECT pglogical.replication_set_add_all_tables('default', ARRAY['public']);

-- 添加 public 模式中所有序列
SELECT pglogical.replication_set_add_all_sequences('default', ARRAY['public']);
```

### 订阅者设置

```sql
-- 在订阅者上创建节点
SELECT pglogical.create_node(
    node_name := 'subscriber1',
    dsn := 'host=subscriberhost port=5432 dbname=mydb'
);

-- 创建到提供者的订阅
SELECT pglogical.create_subscription(
    subscription_name := 'subscription1',
    provider_dsn := 'host=providerhost port=5432 dbname=mydb'
);
```

### 复制集管理

```sql
-- 创建自定义复制集
SELECT pglogical.create_replication_set('my_set');

-- 添加特定表
SELECT pglogical.replication_set_add_table('my_set', 'my_table', true);

-- 移除表
SELECT pglogical.replication_set_remove_table('my_set', 'my_table');
```

### 行和列过滤

```sql
-- 行过滤：仅复制匹配条件的行
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    row_filter := 'id > 1000'
);

-- 列过滤：仅复制特定列
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    columns := '{id, name, updated_at}'
);
```

### 订阅管理

```sql
-- 检查订阅状态
SELECT * FROM pglogical.show_subscription_status();

-- 删除订阅
SELECT pglogical.drop_subscription('subscription1');
```

### 关键特性

- 选择性复制（按表、行过滤、列过滤）
- 不同 PostgreSQL 主版本之间的复制
- 延迟复制
- 订阅者无需超级用户
