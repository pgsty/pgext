## 用法

来源：

- [PGXN 1.1.0 README](https://pgxn.org/dist/pg_plan_guard/1.1.0/README.html)
- [pg_plan_guard 控制文件](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.0/pg_plan_guard.control)
- [pg_plan_guard 1.1 SQL 定义](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.0/pg_plan_guard--1.1.sql)
- [PostgreSQL 许可证](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.0/LICENSE)

`pg_plan_guard` 把认可的 PostgreSQL 19 计划形态保存为 plan-advice 文本，并在重新规划产生不同 advice 时报告漂移。它是关键查询的监控层，不是优化器，也不会自动选择基线。

### 核心流程

```sql
LOAD 'pg_plan_advice';
CREATE EXTENSION pg_plan_guard;

SELECT plan_guard.capture(
    'semantic_search',
    'SELECT id FROM docs ORDER BY embedding <=> ''[1,2,3]'' LIMIT 10',
    'must use the vector index'
);

SELECT * FROM plan_guard.verify();
SELECT * FROM plan_guard.status WHERE state <> 'ok';
```

`plan_guard.capture(...)` 使用 `EXPLAIN (PLAN_ADVICE)`，不会执行查询。`plan_guard.verify(...)` 重新规划已存 SQL，比较生成的 advice，记录漂移或错误状态转换，并在某个基线无法规划时继续处理其他项。

### 对象与可选 Stash

- `plan_guard.baselines` 保存认可的查询文本、advice、可读计划与当前状态。
- `plan_guard.drift_log` 是只追加的状态转换历史。
- `plan_guard.status` 是监控视图。
- `plan_guard.advice_for(...)` 与 `plan_guard.query_id_for(...)` 暴露底层规划信息。
- `plan_guard.sync_stash(...)` 把认可 advice 复制到 `pg_stash_advice` stash。

扩展撤销了会规划已存 SQL 的函数对 public 的执行权限。只有能够规划所有已存查询并更新基线表的角色，才应获得捕获与验证权限。

### 要求与边界

`pg_plan_guard` 要求 PostgreSQL 19 与 `pg_plan_advice` 模块，不能在 PostgreSQL 18 上使用。通过 `plan_guard.sync_stash(...)` 应用 advice 还需要把 `pg_stash_advice` 加入 `shared_preload_libraries`，并设置非空的 `pg_stash_advice.stash_name`。验证只检测漂移；除非配置这条可选 stash 路径，否则不会强制执行计划。

查询以可执行 SQL 文本保存并重新规划，因此参数化应用查询需要具有代表性的字面量。基线也可能认可劣质计划，批准前应审查捕获的 advice。虽然 `EXPLAIN` 不执行查询，规划仍可能获取目录锁、调用规划钩子，并在模式或权限变更后失败。

### 持续断言集成

控制版本 `1.1` 新增 `plan_guard.watch(name, query_sql, note)`，通过 `pg_living_assertions` 注册计划保持不变的断言。只有 `plan_guard.watch` 需要该配套扩展，普通捕获和验证仍可独立使用。调用可选函数前应显式安装配套扩展。PGXN 发行版本为 1.1.0。
