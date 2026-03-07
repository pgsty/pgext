
## 用法

首先将扩展添加到 shared_preload_libraries 中：

```ini
shared_preload_libraries = 'pg_documentdb_core, pg_stat_statements, auto_explain'
```

创建扩展并执行 DDL 与 CRUD 操作的示例：

```sql
-- CASCADE 会自动安装 documentdb_core、pg_cron、vector 等依赖扩展
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

目前 DocumentDB 可与 FerretDB 2.0+ 配合使用，作为 MongoDB 兼容的后端。
