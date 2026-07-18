## 用法

来源：

- [PostgreSQL 开发版文档](https://www.postgresql.org/docs/devel/pgstashadvice.html)
- [已复核 PostgreSQL commit 的扩展 control 文件](https://github.com/postgres/postgres/blob/3cf5264557bee2ba848e5276beecc10571d468a6/contrib/pg_stash_advice/pg_stash_advice.control)
- [SQL API 与默认权限](https://github.com/postgres/postgres/blob/3cf5264557bee2ba848e5276beecc10571d468a6/contrib/pg_stash_advice/pg_stash_advice--1.0.sql)

`pg_stash_advice` 把 plan advice 字符串存入按 PostgreSQL query ID 索引、使用动态共享内存的命名 stash。会话通过 `pg_stash_advice.stash_name` 选择 stash；规划匹配查询时，保存的 advice 会被自动应用。

### 创建并选择 Stash

预加载是让每个会话都能应用 advice 并自动持久化的常规方式。

```conf
shared_preload_libraries = 'pg_stash_advice'
pg_stash_advice.persist = on
pg_stash_advice.persist_interval = 30
```

重启后，分别通过 `EXPLAIN (VERBOSE)` 和 `EXPLAIN (PLAN_ADVICE)` 取得 query ID 与 advice 字符串，再将其存储：

```sql
CREATE EXTENSION pg_stash_advice;

SELECT pg_create_advice_stash('critical_queries');
SELECT pg_set_stashed_advice(
  'critical_queries',
  :'query_id'::bigint,
  :'plan_advice'
);

SET pg_stash_advice.stash_name = 'critical_queries';
```

检查函数包括 `pg_get_advice_stashes()` 和 `pg_get_advice_stash_contents(text)`。`pg_drop_advice_stash(text)` 用于删除 stash。安装脚本撤销了公众对所有管理与检查函数的执行权限。

### 注意事项

- 本页描述的是 PostgreSQL 开发分支代码，官方页面也把该分支标记为不受支持。请确认确切的服务器构建包含版本 `1.0`，不要假定某个已发布大版本一定可用。
- Stash 会消耗动态共享内存。应用 advice 会增加规划开销，还可能阻止 planner 随数据分布变化而调整，因此必须限制 stash 的大小与作用范围。
- 持久化默认开启，由后台工作进程把数据写入数据目录中的 `pg_stash_advice.tsv`。自动启动工作进程要求预加载；`pg_start_stash_advice_worker()` 可以手动启动持久化，但不会加载已有文件。
- 安全模型有意保持简单。尽管管理函数默认不向公众开放，任何用户都可以选择已知的 stash 名称，这可能暴露其中的 advice 字符串。不要在 plan advice 中存储秘密。
- Query ID 依赖 `compute_query_id`；当该参数为 `auto` 时，加载模块会启用计算。查询或 schema 变化后，必须维护绑定到旧 query ID 的 advice。
