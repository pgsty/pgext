

## 用法

> [ora_btree_gist: 支持在 GiST 中索引 Oracle 常见数据类型](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ora_btree_gist)

`ora_btree_gist` 扩展为 IvorySQL 提供的 Oracle 兼容数据类型添加了 GiST（广义搜索树）操作符类支持。

### 启用

```sql
CREATE EXTENSION ora_btree_gist;
```

### 在 Oracle 类型上创建 GiST 索引

该扩展允许您在 Oracle 兼容数据类型（如 `NUMBER`、`VARCHAR2` 和 Oracle 风格 `DATE`）上创建 GiST 索引，类似于标准 `btree_gist` 扩展对原生 PostgreSQL 类型的作用。

```sql
CREATE TABLE t (val NUMBER);
CREATE INDEX t_val_gist_idx ON t USING gist (val);
```

### 使用场景

Oracle 兼容类型的 GiST 索引支持：

- 使用 Oracle 数据类型的排他约束
- Oracle 类型列的最近邻搜索
- Oracle 兼容类型的 GiST 优化范围查询

该扩展是 IvorySQL Oracle 兼容套件的一部分，需要 `ivorysql_ora` 可用。
