## 用法

来源：

- [上游 README](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/README.md)
- [扩展 control 文件](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/vectorize_engine.control)
- [SQL 安装脚本](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/vectorize_engine--1.0.sql)
- [规划器钩子实现](https://github.com/zhangh43/vectorize_engine/blob/e886c5d2b3e61d1c1e8b34cfa46664a454a7ef05/plan.c)

`vectorize_engine` `1.0` 版是面向 PostgreSQL 9.6 的原型向量化执行器。后规划器钩子把受支持的顺序扫描和聚合节点替换为自定义扫描节点；功能由默认关闭的 GUC 控制。

### 示例

使用合适 SIMD 参数构建 PostgreSQL，把 `vectorize_engine` 加入 `shared_preload_libraries`，重启后执行：

```sql
CREATE EXTENSION vectorize_engine;
SET enable_vectorize_engine = on;
EXPLAIN (ANALYZE, BUFFERS) SELECT category, sum(amount)
FROM benchmark_rows GROUP BY category;
```

源码复制并修改 PostgreSQL 执行器代码，且明确基于 PostgreSQL 9.6；仓库自 2020 年后没有变化。只有有限的节点/类型/表达式子集会向量化，不受支持的计划预期回退。若不进行大量移植，它无法兼容现代服务器内部接口。只能在匹配版本的可丢弃基准实例中使用，并应把每个结果和计划与关闭该功能时对比。
