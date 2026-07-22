## 用法

来源：

- [官方 README](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/README.md)
- [官方控制文件](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/geqo_extension.control)
- [官方模块入口](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/src/geqo_main.c)
- [官方 SQL 测试](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/test/sql/01_geqo_set_parameters.sql)

`geqo_extension` 0.0.1 是 2015 年的实验项目，它把 PostgreSQL Genetic Query Optimizer 抽取成可加载的 planner-hook 库。它没有扩展 SQL 对象，应使用 `LOAD` 而不是 `CREATE EXTENSION`；其用途是规划器研究，不是服务器内置 GEQO 的可移植替代品。

### 会话流程

在一次性会话中加载库并设置规划器参数，再比较外部 hook 启用前后的计划。

```sql
BEGIN;

LOAD 'geqo_extension';
SET geqo_extension = on;
SET geqo_threshold = 2;
SET geqo_effort = 5;

EXPLAIN (COSTS ON)
SELECT *
FROM a
JOIN b USING (id)
JOIN c USING (id);

ROLLBACK;
```

`LOAD` 会安装进程本地的 `join_search_hook`；后端结束后即移除。用户可设置的 GUC `geqo_extension` 在启用时选择抽取出的优化器。

### 规划器行为

模块复制了 PostgreSQL 优化器源码，并使用 `geqo_effort`、`geqo_pool_size`、`geqo_generations`、`geqo_selection_bias` 和 `geqo_seed` 等内置 GEQO 参数。由于搜索是遗传式且对参数敏感，计划质量与规划时间可能变化。

在固定源码中，只要 `geqo_extension` 为 true，hook 的第一个分支就会调用外部优化器，因此后面的 `geqo_threshold` 分支在该状态下不可达；不要假设阈值会保护小型 join。必须通过 `EXPLAIN` 确认每种测试查询形态的行为。

### 兼容性与运维边界

Planner hook 和复制的优化器内部代码与 PostgreSQL 源码版本紧密耦合。这些代码早于现代 PostgreSQL 版本，也没有维护中的兼容矩阵。只能使用匹配的服务器头文件编译，并在加载前运行规划器和回归测试。

系统只有一条 `join_search_hook` 链，与其他规划器扩展的加载顺序会产生影响。错误计划或模块缺陷会影响加载该库的后端查询。应使用隔离基准测试、保留对照会话，且不能从小型合成 join 推断生产收益。
