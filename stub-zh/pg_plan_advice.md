## 用法

来源：

- [PostgreSQL 19 documentation](https://www.postgresql.org/docs/19/pgplanadvice.html)
- [Module definition](https://github.com/postgres/postgres/blob/REL_19_BETA3/contrib/pg_plan_advice/Makefile)
- [Module implementation](https://github.com/postgres/postgres/blob/REL_19_BETA3/contrib/pg_plan_advice/pg_plan_advice.c)
- [Versioned manual](https://github.com/postgres/postgres/blob/REL_19_BETA3/doc/src/sgml/pgplanadvice.sgml)

`pg_plan_advice` 是 PostgreSQL 19 附带的模块，将规划器选择记录为建议，并据此约束后续规划。本文以 PostgreSQL 19 beta3 为版本边界，该模块没有独立的 SQL 扩展版本。

### 会话工作流

```sql
LOAD 'pg_plan_advice';
CREATE TEMP TABLE advice_demo (id integer PRIMARY KEY, body text);
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
SET pg_plan_advice.advice = 'SEQ_SCAN(d)';
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
RESET pg_plan_advice.advice;
```

### 启用与反馈

可在拥有足够权限时用 `LOAD` 加载到单个会话，或通过 `session_preload_libraries` 为新会话加载，也可配置 `shared_preload_libraries` 后重启服务器。没有扩展 DDL 步骤。`EXPLAIN (PLAN_ADVICE)` 使用关系别名作为目标生成建议，`pg_plan_advice.advice` 则约束扫描、连接顺序、连接方法或并行方式。

应检查返回的反馈，识别未匹配、冲突、不适用或失败的建议。`pg_plan_advice.feedback_warnings` 可发出警告，`pg_plan_advice.always_store_advice_details` 以额外规划开销保留解释预备语句时有用的细节。

### 限制

建议只能约束核心规划器认为可行的计划，不能强制语义无效的计划，也不能控制聚合策略或集合操作。固定建议可能随数据变化而失效。`pg_stash_advice` 可按查询 ID 保存建议，`pg_plan_guard` 则监控相对认可基线的漂移。
