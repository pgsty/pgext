## 用法

来源：

- [Official documentation](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/README.MD)
- [Control file](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans.control)
- [Version 1.0 SQL](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans--1.0.sql)
- [Runtime and access checks](https://github.com/postgredients/pg_stat_query_plans/blob/f0049ce1d02b1912cd7475c9fc91e1e1a441238c/pg_stat_query_plans.c)

`pg_stat_query_plans` 1.0 在共享内存中跟踪语句及其执行计划，可通过查询级和计划级视图定位执行行为变化。它是独立扩展，并非安装 pg_stat_statements 后附加的视图。

### 启用与核心流程

将模块加入现有预加载列表，重启 PostgreSQL，再在需要查询统计视图的数据库中创建扩展：

```conf
shared_preload_libraries = 'pg_stat_query_plans'
```

```sql
CREATE EXTENSION pg_stat_query_plans;
SELECT queryid, calls, total_exec_time, query
FROM pg_stat_query_plans_sql
ORDER BY total_exec_time DESC
LIMIT 10;
SELECT queryid, planid, calls, normalized_plan
FROM pg_stat_query_plans
ORDER BY calls DESC
LIMIT 10;
SELECT * FROM pg_stat_query_plans_info;
```

### 对象与配置

- `pg_stat_query_plans_sql` 按数据库、用户、查询标识与顶层执行状态聚合指标。
- `pg_stat_query_plans` 进一步区分计划标识，包含代表性查询文本、规范化计划和示例计划。
- `pg_stat_query_plans_info` 展示分配、文本存储、淘汰与重置信息。
- `pg_stat_query_plans.track` 选择采集范围；`pg_stat_query_plans.track_planning` 开启规划耗时统计，默认关闭。
- `pg_stat_query_plans_reset(userid, dbid, queryid)` 重置匹配的统计，参数全为零时选择全部记录；`pg_stat_query_plans_reset_minmax()` 清除极值。安装 SQL 撤销了这两个函数的公共执行权限。

### 运行边界

安装需要管理权限，修改预加载配置后需要重启。普通角色可查询视图，但 C 读取函数会按调用者检查权限，不应假定不同角色看到的记录相同。代表性查询和计划可能保留字面量，向监控用户开放前应审查权限。存储有容量限制，记录可能被淘汰，因此这些计数不能替代审计日志。固定版本源码没有给出完整的 PostgreSQL 主版本支持矩阵，应另行验证目标服务器版本。
