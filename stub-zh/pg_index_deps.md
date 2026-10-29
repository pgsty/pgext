## 用法

来源：

- [README.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/README.md)
- [INSTALL.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/INSTALL.md)
- [pg_index_deps.control](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/pg_index_deps.control)
- [pg_index_deps--0.7.sql](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/pg_index_deps--0.7.sql)
- [Makefile](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/sql_edition/Makefile)
- [ci.yml](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/.github/workflows/ci.yml)

`pg_index_deps` 是 `pg_stat_index_queries` 的纯 SQL 版本，SQL 版本为 0.7，使用固定模式 `index_deps`。它通过解析已捕获的计划文本判断索引依赖，不使用原生执行钩子。上游测试 PostgreSQL 16–18；Aurora 指南还说明支持 14.10/15.5 或更新版本，但删除影响函数需要 16 或更新版本。

### 核心流程

在已安装扩展文件和计划捕获依赖的自建服务器上：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_stat_plans;
CREATE EXTENSION hypopg;
CREATE EXTENSION pg_index_deps;
SELECT indexname, index_role, queries, verdict
FROM index_deps.index_usage;
SELECT queryid, calls, scan_types, index_role
FROM index_deps.index_dependents WHERE indexname = 'orders_customer_idx';
SELECT * FROM index_deps.index_drop_impact('public.orders_customer_idx');
```

原生 `pg_stat_index_queries` 扩展是可选项。在托管服务器上应使用官方纯 SQL 安装器或 pg_tle 安装器，而非假定可以安装任意服务端控制文件。pg_tle 方式有自己的预加载要求。Aurora 捕获计划需要 `aurora_compute_plan_id`，应按供应商的参数组流程激活。

### 数据来源与对象

`index_deps.config` 可选择 auto、aurora_stat_plans、pg_stat_plans 或 none。`index_deps.plan_index_usage` 解析索引引用，`index_deps.index_dependents` 按查询与索引汇总。`index_deps.index_usage`、`index_deps.index_role`、`index_deps.index_table_stats`、`index_deps.index_column_stats` 和 `index_deps.index_hints` 补充用途与统计背景。`index_deps.unresolved_plan_indexes` 指出已无法解析的历史索引名。

没有计划来源、查询被淘汰、索引改名或删除，以及负载未被捕获，都会限制结果。尤其在没有计划来源时，空依赖集合并不能证明索引未被使用。查询耗时仍然属于整条语句，并非逐索引耗时。

### 删除影响边界

`index_deps.index_drop_impact` 用 HypoPG 在当前会话隐藏索引，再比较 `EXPLAIN (GENERIC_PLAN)` 的成本。它不会执行语句，也不会删除真实索引。成本比率是规划器估计，而非延迟测量；陈旧统计和带参数化谓词的部分索引可能使结果无法判断。采取行动前应审查阻碍、真实参数值与代表性负载。

### 升级

安装匹配的 0.7 文件后，由扩展管理的安装可以执行：

```sql
ALTER EXTENSION pg_index_deps UPDATE TO '0.7';
```

pg_tle 安装应先注册新版安装器和更新路径，再更新扩展。纯 SQL 安装通过其安装器维护，不会自动在 `pg_extension` 中登记。
