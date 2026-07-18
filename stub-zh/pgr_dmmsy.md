## 用法

来源：

- [锁定提交的上游 README](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/README.md)
- [0.1.0 版 SQL API](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/sql/pgr_dmmsy--0.1.0.sql)
- [锁定提交的 PostgreSQL C 包装器](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/src/dmmsy.c)
- [锁定提交的算法实现](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/src/dmmsy_algorithm.c)
- [锁定提交的参数测试](https://github.com/offline-agency/pgr_dmmsy/blob/f8de14341094b1e2d74ed48d4d76a19e66c4cf36/test/sql/test_parameters.sql)

pgr_dmmsy 0.1.0 实现了一个带 pgRouting 风格边查询的有向或无向非负单源最短路径函数。它需要 PostgreSQL 14 或更高版本和 pgRouting 3.x。该函数当前可作为点到点路径查询；文档中的全目标形式并未正确实现。

### 点到点路径

边 SQL 必须严格按该顺序返回 id BIGINT、source BIGINT、target BIGINT 和 cost FLOAT8：

```sql
CREATE EXTENSION pgrouting;
CREATE EXTENSION pgr_dmmsy;

CREATE TABLE edges (
  id bigint,
  source bigint,
  target bigint,
  cost double precision
);

INSERT INTO edges VALUES
  (1, 1, 2, 1.5),
  (2, 2, 3, 2.0),
  (3, 3, 4, 0.5);

SELECT seq, path_seq, node, edge, cost, agg_cost
FROM pgr_dmmsy(
  'SELECT id, source, target, cost FROM edges',
  1,
  4
)
ORDER BY path_seq;
```

NULL 边字段和负成本会报错。directed 为 false 时，会为每个输入行添加一条反向边。未知或不可达目标不返回行。

### 当前实现边界

README 声称 NULL 目标会返回所有可达顶点的距离图，但 C 包装器只在 target 非 NULL 且非负时构造输出。锁定回归测试明确预期 NULL 目标返回空集。在包装器拥有通过测试的全目标输出路径之前，应始终提供具体目标。

output_predecessors 和 constant_degree 被接受，但当前源码并未使用它们。max_levels 可能在目标确定前停止阶段处理，随后不返回路径。仓库的实现说明与当前算法注释也存在分歧：它究竟是简化的 Dijkstra 实现，还是宣称的渐进 DMMSY 算法。应将复杂度声明视为未验证，并在自有图上与 pgr_dijkstra 对比正确性和性能。

包装器在计算前会把完整边查询读入后端内存。它只检查成本是否小于零，所以调用者应显式拒绝 NaN 和无穷大。任意 edges_sql 使用调用者权限执行；应只允许可信 SQL 模板，对每个对象使用限定名，且不得拼接用户输入。

该项目处于活跃状态，但版本仍为 0.1.0，没有声明许可证，CI 也只覆盖所声称支持大版本的一个子集。投入生产前，应验证等成本平局、平行边、自环、零权重、巨大或稀疏顶点 ID、不连通图、取消、内存压力、directed false、每个调优参数，并与可信实现比较。
