

## 用法

> [pgoutput: 逻辑复制输出插件](https://www.postgresql.org/docs/current/protocol-logical-replication.html)

PostgreSQL 原生逻辑复制系统使用的内置逻辑解码输出插件。它是 `CREATE PUBLICATION` / `CREATE SUBSCRIPTION` 的默认插件。

### 概述

`pgoutput` 是 PostgreSQL 原生的逻辑复制输出插件。与第三方插件不同，它不需要单独安装——它内置于 PostgreSQL 核心中（10+ 版本）。

### 使用原生逻辑复制

在 **发布端**：

```sql
-- 为特定表创建发布
CREATE PUBLICATION my_pub FOR TABLE orders, customers;

-- 或发布所有表
CREATE PUBLICATION my_pub FOR ALL TABLES;

-- 带行过滤的发布（PG 15+）
CREATE PUBLICATION filtered_pub FOR TABLE orders WHERE (status = 'active');

-- 带列过滤的发布（PG 15+）
CREATE PUBLICATION col_pub FOR TABLE users (id, name, email);
```

在 **订阅端**：

```sql
-- 创建订阅
CREATE SUBSCRIPTION my_sub
    CONNECTION 'host=publisher dbname=mydb user=replicator'
    PUBLICATION my_pub;

-- 检查订阅状态
SELECT * FROM pg_stat_subscription;
```

### 直接使用 pgoutput 与复制槽

```sql
-- 使用 pgoutput 创建逻辑复制槽
SELECT * FROM pg_create_logical_replication_slot('my_slot', 'pgoutput');

-- 消费变更（需要协议级参数）
-- pgoutput 设计用于流复制协议，
-- 而非 SQL 槽函数（SQL 测试请使用 test_decoding）
```

### 发布管理

```sql
-- 向现有发布添加表
ALTER PUBLICATION my_pub ADD TABLE new_table;

-- 移除表
ALTER PUBLICATION my_pub DROP TABLE old_table;

-- 查看发布
SELECT * FROM pg_publication;
SELECT * FROM pg_publication_tables;
```

### 订阅管理

```sql
-- 禁用订阅
ALTER SUBSCRIPTION my_sub DISABLE;

-- 刷新订阅（获取新表）
ALTER SUBSCRIPTION my_sub REFRESH PUBLICATION;

-- 删除订阅
DROP SUBSCRIPTION my_sub;
```

### 主要功能

- 内置于 PostgreSQL 核心（10+）
- 二进制格式用于高效数据传输
- 行和列过滤（PostgreSQL 15+）
- 支持初始表同步
- 通过订阅刷新处理模式变更
- 支持每个订阅多个发布
