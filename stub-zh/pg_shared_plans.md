## 用法

来源：

- [官方 README](https://github.com/rjuju/pg_shared_plans/blob/master/README.md)
- [扩展 SQL 与权限定义](https://github.com/rjuju/pg_shared_plans/blob/master/pg_shared_plans--0.0.1.sql)
- [预加载与配置实现](https://github.com/rjuju/pg_shared_plans/blob/master/pg_shared_plans.c)

`pg_shared_plans` 0.0.1 是明确标注的概念验证扩展，用共享内存透明缓存执行计划。它可能减少合格语句的重复规划开销，但上游声明尚未达到生产可用状态，而且无法在备库上可靠地因 DDL 变更失效计划。

### 预加载与设置

模块支持 PostgreSQL 12 及以上版本，依赖 `pg_stat_statements`，并且必须先预加载，再创建扩展对象：

```conf
shared_preload_libraries = 'pg_stat_statements,pg_shared_plans'
```

```sql
CREATE EXTENSION pg_shared_plans CASCADE;

SELECT * FROM pg_shared_plans_info;
SELECT * FROM pg_shared_plans;
```

修改预加载列表后必须重启 postmaster。PostgreSQL 14 及以上版本还需要启用查询标识符计算；当核心设置允许时，模块会调用服务器钩子启用它。

### 控制项与视图

重要控制项包括 `pg_shared_plans.enabled`、`pg_shared_plans.max`、`pg_shared_plans.min_plan_time`、`pg_shared_plans.threshold`、`pg_shared_plans.rdepend_max` 和 `pg_shared_plans.read_only`。当前 C 源码中 `pg_shared_plans.max` 的默认值是 1000，而 README 仍写 200。

`pg_shared_plans` 提供摘要，`pg_shared_plans_relations` 提供关系 OID，`pg_shared_plans_explain` 提供计划文本，`pg_shared_plans_all` 同时提供两者。`pg_shared_plans_reset(userid, dbid, queryid)` 删除匹配的缓存项，安装脚本已撤销 PUBLIC 对它的权限。

### 安全与运维说明

缓存的查询与计划文本可能暴露模式名、关系细节和谓词。底层 `pg_shared_plans(boolean, boolean, oid, oid)` 函数仍保留默认 PUBLIC 执行权限，而经过整理的视图权限更窄；应撤销该函数权限，并只向受信角色授予所需监控视图。重启会清空共享内存内容。DDL 失效、RLS 身份处理、预备语句行为、淘汰、内存尺寸与故障转移都应视为实验范围；未经源码级验证，绝不能在备库或生产集群启用该概念验证实现。
