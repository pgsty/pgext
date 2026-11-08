## 用法

来源：

- [README v1.4.1](https://github.com/GyuyeongKim/Tessera-public/blob/v1.4.1/README.md)
- [v1.4.1 release](https://github.com/GyuyeongKim/Tessera-public/releases/tag/v1.4.1)
- [tessera_pg.control](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.4.1/tessera_pg.control)
- [tessera_pg--1.0.sql](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.4.1/tessera_pg--1.0.sql)
- [Makefile](https://api.pgxn.org/src/tessera_pg/tessera_pg-1.4.1/Makefile)

`tessera_pg` 通过稀疏矩阵乘法替换受支持的连接聚合计划，并提供条目相似度函数。分发版本 1.4.1 安装的 SQL 版本为 1.0，目标是 PostgreSQL 18；原生引擎需要 OpenMP，x86-64 平台还需要 AVX2 或更新的指令集。

### 核心流程

先在会话中启用规划器钩子，再安装 SQL 函数：

```sql
LOAD 'tessera_pg';
CREATE EXTENSION tessera_pg;
CREATE TABLE edges (k bigint NOT NULL, x bigint NOT NULL);
INSERT INTO edges VALUES (1,10),(1,20),(2,10),(2,30);
EXPLAIN SELECT a.x, b.x, COUNT(*)
FROM edges a JOIN edges b ON a.k = b.k
GROUP BY a.x, b.x;
SELECT * FROM tessera_sim('edges', 'k', 'x', 'jaccard', 0.5);
```

出现 `Custom Scan (Tessera)` 表示计划已被接管。1.4.1 版本新增对可空列求和与计数、使用 `RANK()` 或 `ROW_NUMBER()` 的组内 top-k 裁剪，以及一次计算多个浮点求和项。只有所有因子都非 NULL 的连接行对才参与求和；没有有效贡献行对的分组返回 NULL。扩展仍支持 `COUNT(DISTINCT a.k)`、`AVG`、单侧 `SUM`、表达式连接键和分组列、输入过滤、HAVING 与 top-N 裁剪。

PostgreSQL 对不支持的情况仍使用常规计划，包括可空 decimal 权重、可空 `float8` 平均值、整数与 `float8` 的乘积、`real` 乘积、`MIN`、`MAX`、外连接及非确定性排序规则。多个求和项与可空求和输入对于分区表、继承表或自带计划的输入还有额外限制。应查阅对应版本的支持矩阵；适配器仅识别 PostgreSQL 内置运算符和聚合函数。

### 相似度函数与设置

`tessera_sim` 返回 bigint 条目对及其分数，`tessera_sim_text` 返回文本条目对。参数指定关系、键列、条目列、相似度度量、阈值、可选权重列、条目对模式和线程数。两个 SQL 函数均标记为 volatile 与 parallel-unsafe。

`tessera.enabled` 默认开启，`tessera.threads` 默认为 0，`tessera.frontier` 默认为 -1，表示由成本模型选择。另有 `tessera.min_expansion`、`tessera.min_expansion_cross` 和 `tessera.trace`。默认线程数量会考虑主执行进程与查询的 `max_parallel_workers_per_gather` 工作进程。

### 运行边界

通过 `LOAD` 或 `session_preload_libraries` 激活钩子即可，不强制要求共享预加载或重启。`CREATE EXTENSION` 安装可迁移的 SQL 对象。计划内存可能达到 `work_mem * hash_mem_multiplier`；溢写使用 `temp_tablespaces`，并受 `temp_file_limit` 限制。应针对实际负载评估执行计划、临时磁盘和并发量，源码基准测试不代表生产环境的加速保证。
