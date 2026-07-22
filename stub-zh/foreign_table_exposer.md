## 用法

来源：

- [Official extension control file](https://github.com/komamitsu/foreign_table_exposer/blob/1c57a554d1c50c872fac3b103317b11dec10e95f/foreign_table_exposer.control)
- [Official upstream documentation](https://pgxn.org/dist/foreign_table_exposer/1.0.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/foreign_table_exposer/)

`foreign_table_exposer` 1.0 让只查找普通表和视图的目录客户端也能发现 PostgreSQL 外部表。它通过改写符合条件的 `pg_catalog.pg_class` 查询实现这一点，不会改变外部表的存储或事务语义。

### 核心流程

让模块随服务器启动加载，重启 PostgreSQL，并在需要此行为的每个数据库中启用其 SQL 对象。

```conf
shared_preload_libraries = 'foreign_table_exposer'
```

```sql
CREATE EXTENSION foreign_table_exposer;
```

启用后，等价于 `relkind IN ('r', 'v')` 的目录条件会被改写，纳入 relkind 为 `f` 的外部表。它主要面向原本会隐藏外部表的旧 BI 与元数据工具。

### 运维边界

改写钩子会影响整个服务器中匹配的目录查询，可能让本来有意排除外部表的客户端得到意外结果。它不会让 FDW 操作变成本地、事务性或低开销操作。该版本发布于 2015 年；投入生产前，应在实际 PostgreSQL 大版本、其他使用钩子的扩展以及每种客户端查询上测试。
