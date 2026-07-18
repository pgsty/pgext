## 用法

来源：

- [锁定提交的上游 README](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/README.md)
- [0.3.0 版安装 SQL](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats--0.3.0.sql)
- [0.3.0 版破坏性升级脚本](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats--0.2.0--0.3.0.sql)
- [锁定提交的扩展控制文件](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats.control)

mv_stats 0.3.0 使用数据库级事件触发器，记录物化视图的创建、修改、刷新次数和刷新耗时。它通过 mv_stats 视图暴露收集的行，并为安装前已存在的每个物化视图初始化一行空数据。因为要创建事件触发器，安装需要超级用户。

### 查看与重置统计

```sql
CREATE EXTENSION mv_stats;

CREATE MATERIALIZED VIEW public.mv_example AS
SELECT current_database() AS database_name;

REFRESH MATERIALIZED VIEW public.mv_example;

SELECT mv_name, refresh_count, refresh_mv_last,
       refresh_mv_time_last, refresh_mv_time_total
FROM mv_stats
WHERE mv_name = 'public.mv_example';

SELECT *
FROM mv_activity_reset_stats('public.mv_example');
```

安装前已存在的物化视图，其创建时间戳为 NULL。重置函数接受一个带模式限定的名称，或使用默认星号处理每一行；它会清空刷新计数和耗时，同时记录 reset_last。

### 事件触发器、暴露与升级限制

三个事件触发器会在数据库中每次 CREATE、ALTER、REFRESH 或 DROP MATERIALIZED VIEW 前后运行。刷新计时保存在事务局部自定义设置中，结果是普通表行，而不是具有崩溃安全性的累积服务器统计。回滚的 DDL 也会回滚对应行更改。并发刷新会更新同一统计行，因此应在实际负载下验证计数和总时间，不应将它用于计费或严格审计。

扩展向 PUBLIC 授予了视图和后端表的查询权，会向每个数据库角色暴露物化视图名称和耗时。如果对象目录或工作负载耗时属于敏感信息，应撤销这些授权。事件函数使用 SECURITY DEFINER，但没有设置受保护的 search_path。应复核所有权和函数权限，不让不可信模式进入特权 search_path，并限制重置和破坏性辅助函数的执行权。

随附的 0.2.0 到 0.3.0 升级脚本会显式删除并重建表、视图、函数和触发器，从而删除全部收集统计。上游 README 要求用户在更新前复制视图。应捕获并验证该副本，在维护窗口内运行精确的版本化 ALTER EXTENSION 命令，随后确认所有事件触发器。

README 记录了 PostgreSQL 10 及更高版本，但没有提供当前主版本矩阵或性能特性。在投入生产前，应在精确 PostgreSQL 版本上测量 DDL 开销、长时间与失败刷新、模式重命名、引号标识符、并发 DDL、扩展重定位、备份/恢复和升级。
