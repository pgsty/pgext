

## 用法

> [pgpool_regclass: regclass 的替代实现](https://pgpool.net/)

`pgpool_regclass` 扩展提供一个替代的 `regclass` 函数，由 Pgpool-II 内部使用，用于处理跨多个后端的关系名称解析。

### 函数

```sql
-- 将关系名称解析为其 OID，类似于 PostgreSQL 的 regclass 强制转换
SELECT pgpool_regclass('my_table');
SELECT pgpool_regclass('my_schema.my_table');
```

### 用途

在标准 PostgreSQL 中，将字符串转换为 `regclass`（例如 `'my_table'::regclass`）会将关系名称解析为 OID。然而，Pgpool-II 需要判断 SQL 语句引用的是临时表还是普通表，以正确路由查询。

`pgpool_regclass` 函数以普通函数调用（而非类型转换）的形式提供此解析能力，使 Pgpool-II 能够：

- 判断引用的表是否存在于后端
- 区分临时表和永久表以实现正确的查询路由
- 在连接池中正确处理带模式限定的表名

### 注意事项

- 此扩展主要由 Pgpool-II 内部使用，通常不由应用程序直接调用
- 应安装在 Pgpool-II 管理的所有 PostgreSQL 后端节点上
- 对于 Pgpool-II 3.0+，此函数有助于 `check_temp_table` 功能
