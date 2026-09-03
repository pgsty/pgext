## 用法

来源：

- [官方文档](https://github.com/Alena0704/vacuum_statistics/blob/c40b73e3bcce0343b9e4c86eab353be26d6e006c/README.md)
- [扩展控制文件](https://github.com/Alena0704/vacuum_statistics/blob/c40b73e3bcce0343b9e4c86eab353be26d6e006c/vacuum_statistics.control)
- [官方仓库](https://github.com/Alena0704/vacuum_statistics)

`vacuum_statistics` 在共享内存中记录 VACUUM 的表、数据库与索引统计。

### 启用

将 `vacuum_statistics` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `vacuum_statistics`：

```ini
shared_preload_libraries = 'vacuum_statistics'
```

```sql
CREATE EXTENSION vacuum_statistics;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT relname, tuples_deleted, pages_scanned, pages_removed, total_time
FROM pg_stats_vacuum_tables
ORDER BY total_time DESC
LIMIT 10;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `extvac_reset_entry` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stats_get_vacuum_database` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stats_get_vacuum_indexes` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stats_get_vacuum_tables` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_stats_vacuum_database` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stats_vacuum_indexes` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stats_vacuum_tables` | VIEW | 扩展创建的检查或查询视图。 |
| `extvac_reset_db_entry` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 预加载 `vacuum_statistics` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
