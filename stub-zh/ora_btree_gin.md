

## 用法

> [ora_btree_gin: 支持在 GIN 中索引 Oracle 数据类型](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ora_btree_gin)

`ora_btree_gin` 扩展为 IvorySQL 提供的 Oracle 兼容数据类型添加了 GIN（广义倒排索引）操作符类支持。

### 启用

```sql
CREATE EXTENSION ora_btree_gin;
```

### 在 Oracle 类型上创建 GIN 索引

该扩展允许您在 Oracle 兼容数据类型（如 `NUMBER`、`VARCHAR2` 和 Oracle 风格 `DATE`）上创建 GIN 索引，类似于标准 `btree_gin` 扩展对原生 PostgreSQL 类型的作用。

```sql
CREATE TABLE t (val NUMBER);
CREATE INDEX t_val_gin_idx ON t USING gin (val);
```

### 使用场景

Oracle 兼容类型的 GIN 索引特别适用于：

- 包含 Oracle 数据类型列的多列索引查询
- 结合全文搜索和 Oracle 类型列过滤的查询
- 任何需要 Oracle 兼容数据类型倒排索引支持的场景

该扩展是 IvorySQL Oracle 兼容套件的一部分，需要 `ivorysql_ora` 可用。
