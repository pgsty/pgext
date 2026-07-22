## 用法

来源：

- [官方 README](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/README.md)
- [官方 0.0.1 版 SQL](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views--0.0.1.sql)
- [官方 control 文件](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views.control)
- [所需的 Postgres-XL 游标补丁](https://git.postgresql.org/gitweb/?p=postgres-xl.git;a=commit;h=1b6ada225da45c82529d56f71e3c6a62fabcfd55)

`xl_global_views` 通过 `EXECUTE DIRECT` 在每个节点执行本地 PostgreSQL 系统视图，为 Postgres-XL 创建集群级 `pgxl_` 监控视图。每条生成记录前都会增加 `node_name` 与 `node_type`，从而能在一条查询中比较协调器和数据节点。它只适用于 Postgres-XL，不适用于普通 PostgreSQL。

### 核心流程

在所需的每个 Postgres-XL 节点上安装版本 `0.0.1`，在 `public` 中创建不可重定位扩展，再显式生成视图：

```sql
CREATE EXTENSION xl_global_views;
SELECT public.pgxl_create_views();

SELECT node_name, node_type, relname, heap_blks_read, heap_blks_hit
FROM public.pgxl_statio_all_tables
WHERE node_type = 'D' AND relname = 'pg_class';
```

底层系统视图结构改变后，应重新运行 `pgxl_create_views()`；它会根据 `information_schema.columns` 删除并重建匹配的 `pgxl_` 视图。

### 对象

- `pgxl_global_view(localtable text, fields text [, cond text, limitGiven int])`：在每个协调器/数据节点执行指定本地系统视图，并返回带节点前缀的记录。
- `pgxl_create_views()`：为匹配的 `pg_stat*`/`pg_statio*` 来源及部分系统目录视图创建 public 视图。
- 文档列出的输出包括 `pgxl_locks`、`pgxl_stat_activity`、`pgxl_stat_database`、`pgxl_stat_replication`、`pgxl_stat_ssl`、`pgxl_stats`，以及表、索引、函数的 I/O 与事务统计视图系列。

### 注意事项

Postgres-XL 构建必须包含上游游标补丁 `1b6ada225da45c82529d56f71e3c6a62fabcfd55`；缺少该补丁时，大多数视图会报告 SPI 错误。在目标 Postgres-XL 实现中，`EXECUTE DIRECT` 需要超级用户权限，因此仅向普通角色授予生成视图的 `SELECT` 并不能使其安全可用。

通用函数会根据文本构造动态 SQL，并默认使用较高的单节点行数上限，只应向可信管理员开放。`anyarray` 与 `pg_node_tree` 类型的字段会被省略，因此全局视图不一定与对应本地系统视图完全一致。项目面向旧版 Postgres-XL 分支，不能把它当作当前 PostgreSQL 监控扩展来安装或测试。
