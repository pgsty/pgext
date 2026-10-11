## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/README.md)
- [pg_plan_guard.control](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/pg_plan_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/CHANGELOG.md)
- [pg_plan_guard--1.1.8--1.1.9.sql](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/pg_plan_guard--1.1.8--1.1.9.sql)

`pg_plan_guard` 1.1.9 保存认可的 PostgreSQL 19 计划 advice，在重新规划已有查询时检测变化。它要求加载 `pg_plan_advice` 模块，不支持 PostgreSQL 18。

### 核心用法

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

### 运行边界

`plan_guard.capture()` 保存查询、advice、作者和搜索路径。`plan_guard.verify()` 重新规划，`plan_guard.status` 提供状态；`plan_guard.drift_log` 记录状态转换并拒绝修改。规划在只读且回滚的子事务中以记录的作者身份执行，检查角色须有权切换到该身份。1.1.7 之前的基线没有作者记录，必须重新捕获，否则返回错误。安装会拒绝使用由无关普通角色预先创建的模式。规划可能调用函数，因此仍须只保存可信 SQL 并审查基线。`plan_guard.watch()` 可选使用 `pg_living_assertions`，普通捕获和验证不依赖它。`plan_guard.sync_stash()` 额外要求预加载 `pg_stash_advice`、重启并配置 stash。Apache-2.0 从 1.1.3 开始适用。

### 1.1.9 的角色隔离

1.1.9 在需要角色隔离时，通过作者拥有的临时 `SECURITY DEFINER` 函数规划已保存的 SQL。规划器求值的函数不能再通过重置角色或会话授权取回执行者身份。临时函数在只读规划子事务中创建并回滚。需要该执行环境时，作者须拥有数据库的 `TEMP` 权限，执行者须有权代理作者；权限不足会记录为基线错误。这修复了 1.1.7/1.1.8 中重置角色的逃逸路径，但保存的 SQL 仍须可信。
