

## 用法

> [pg_fkpart: 基于外键的表分区工具](https://github.com/lemoineat/pg_fkpart)

`pg_fkpart` 基于外键关系对 PostgreSQL 表进行分区。所有函数位于 `pgfkpart` schema 中。

### 创建扩展

```sql
CREATE EXTENSION pg_fkpart;
```

### 按外键分区表

```sql
SELECT pgfkpart.partition_with_fk(
    'public',           -- 要分区的表的 schema
    'my_table',         -- 要分区的表
    'public',           -- 外键表的 schema
    'fk_table',         -- 外键表
    true                -- 支持 SQL RETURNING
);
```

### 取消分区

```sql
SELECT pgfkpart.unpartition_with_fk('public', 'my_table');
```

### 索引管理

将父表的索引和约束传播到所有子表：

```sql
SELECT pgfkpart.dispatch_index('public', 'my_table');
```

删除所有子表上的索引：

```sql
SELECT pgfkpart.drop_index('public', 'my_table', 'my_index_name');
```

删除所有子表上的唯一约束：

```sql
SELECT pgfkpart.drop_unique_constraint('public', 'my_table', 'my_constraint_name');
```

### 分区追踪

该扩展维护一个 `pgfkpart.partition` 表，记录所有已分区的表，包括 schema、表名、外键列和外键表信息。
