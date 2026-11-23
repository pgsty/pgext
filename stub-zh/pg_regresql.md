## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/README.md)
- [pg_regresql.control](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql.control)
- [安装脚本](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql--2.2.1.sql)
- [从 2.2 升级的脚本](https://api.pgxn.org/src/pg_regresql/pg_regresql-2.2.1/pg_regresql--2.2--2.2.1.sql)

`pg_regresql` 让查询规划使用目录统计信息，不再按物理文件大小缩放估算，适合在数据很少但统计信息具有代表性的数据库中做执行计划回归测试。2.2.1 版将 SQL 对象安装在固定的 `pg_catalog` 模式中。

### 启用规划器钩子

钩子支持 PostgreSQL 15 及以上版本。以超级用户在测试会话中加载：

```sql
LOAD 'pg_regresql';
CREATE TABLE test_orders (created_at date);
EXPLAIN SELECT * FROM test_orders WHERE created_at > DATE '2024-06-01';
```

如需在测试实例的新会话中加载，将库加入现有会话预加载列表：

```ini
session_preload_libraries = 'pg_regresql'
```

钩子用 `pg_class` 的值替换表和索引的页数、行数估计，并根据目录统计计算可见性比例。比较计划前应注入有意义的统计信息；单有空表不能代表生产负载。

### 恢复统计信息

在 PostgreSQL 16 和 17 上，先创建扩展，再加载新版 pg_dump 导出的统计信息：

```sql
CREATE EXTENSION pg_regresql;
SELECT pg_restore_relation_stats(
    'schemaname', 'public', 'relname', 'test_orders',
    'relpages', 123513::integer, 'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);
```

2.2 为这两个版本增加 `pg_restore_extended_stats()` 与 `pg_clear_extended_stats()`，并让 `pg_ndistinct` 和 `pg_dependencies` 接受 pg_dump 19 输出的 JSON 输入。在 PostgreSQL 16 上须删除备份中不支持的 transaction_timeout 设置。PostgreSQL 18 已内置表与属性统计恢复，19 开始内置扩展统计恢复；本扩展不会在 18 及以上版本安装这些回移函数。

### 升级与清理

从 SQL 版本 2.0 升级时，上游要求删除并重新创建扩展；从 2.1 升级使用 `ALTER EXTENSION pg_regresql UPDATE`。删除前应检查依赖。扩展不可重定位，安装需要超级用户。

16/17 上名为 `regresql_pg_*_in` 的替代输入函数会刻意脱离扩展成员关系。`DROP EXTENSION` 会保留这些函数，使已安装的类型转换继续工作，因此删除扩展并不等于将目录完全恢复原状。应在受控的规划测试中使用，不能假定它与其他规划器钩子的组合已经过验证。

### PostgreSQL 19 规划器钩子

2.2.1 在 PostgreSQL 19 上使用 `build_simple_rel_hook`，在更早的受支持主版本上使用 `get_relation_info_hook`。对于已启用 2.2 的数据库，安装匹配的文件后执行 `ALTER EXTENSION pg_regresql UPDATE TO '2.2.1';`。这次升级更新扩展版本记录，保留原有 SQL 定义。
