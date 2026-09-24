## 用法

来源：

- [Official documentation](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql--1.3.sql)
- [Control file](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.control)
- [Runtime and configuration](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.c)
- [Upstream test configuration](https://github.com/polardb/PolarDB-for-PostgreSQL/blob/9dfb570c0940ed1db830cc05c2218f3e796c6d4e/external/polar_stat_sql/polar_stat_sql.conf)

`polar_stat_sql` 1.3 在官方 PolarDB PostgreSQL 11 分支中采集查询级内核资源、计划节点、执行阶段、锁等待与存储 I/O 指标。实现依赖 PolarDB 专用探针，不能据此认为兼容原生 PostgreSQL。

### 启用与核心流程

在匹配的 PolarDB 构建上，将两个模块加入已有预加载列表，依赖在前，然后重启：

```conf
shared_preload_libraries = 'pg_stat_statements,polar_stat_sql'
polar_stat_sql.enable_stat = on
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION polar_stat_sql;
SELECT query, datname, rolname, user_time, system_time, reads, writes
FROM polar_stat_sql
ORDER BY user_time DESC
LIMIT 10;
```

预加载顺序很重要，初始化需要读取 `pg_stat_statements.max` 设置。SQL 视图将扩展计数与语句文本、数据库名称和角色名称关联。`polar_stat_sql()` 提供原始标识和计数；`polar_stat_sql_reset()` 清除本扩展统计，其公共执行权限已被安装 SQL 撤销。

### 配置与限制

`polar_stat_sql.sample_rate` 控制采样；`polar_stat_sql.enable_getrusage` 启用内核资源测量；`polar_stat_sql.enable_gather_plan_info` 和 `polar_stat_sql.enable_plan_need_time` 控制计划节点细节和计时；`polar_stat_sql.save` 控制停机时的持久化。额外测量会增加查询执行工作量，应按实际需要选择详细程度。

安装和预加载配置需要管理权限。采集器使用共享状态，部分内核计数依赖平台，结果应作为监控数据而非审计记录。规范源码和版本化 SQL 位于 POLARDB_11_STABLE，本次检查的官方 15 与 17 分支没有该模块，不能假定更新的 PolarDB 或 PostgreSQL 服务器一定提供它。
