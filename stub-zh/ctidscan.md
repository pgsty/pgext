## 用法

来源：

- [官方 README](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/README.md)
- [模块控制文件](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/ctidscan.control)
- [自定义扫描实现](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/ctidscan.c)

`ctidscan` 是一个无扩展 DDL 的自定义扫描演示模块，它会为带有 `ctid` 不等式或范围约束的堆扫描增加规划路径。该模块没有扩展 SQL，只应在研究物理元组范围扫描的会话中加载。

### 核心流程

```sql
LOAD 'ctidscan';

CREATE TABLE ctid_demo AS
SELECT g AS id, repeat('x', 100) AS payload
FROM generate_series(1, 10000) AS g;

SET enable_ctidscan = on;

EXPLAIN (COSTS OFF)
SELECT ctid, id
FROM ctid_demo
WHERE ctid BETWEEN '(2,1)'::tid AND '(3,20)'::tid;
```

规划器提取 `ctid` 的上下界，并在相应数据块附近开始或停止堆扫描。评估自定义路径时，应使用 `SET enable_ctidscan = off` 对照执行计划和结果。

### 对象与配置

- `LOAD 'ctidscan'` 为当前后端初始化模块。
- `enable_ctidscan` 启用或禁用它的规划路径。
- 执行自定义扫描时，模块会输出开始和结束信息。

README 还展示了 `shared_preload_libraries = 'ctidscan'`，但建议局部加载，避免原型影响无关用户。修改 `shared_preload_libraries` 后需要重启服务器。

### 安全与兼容性

`ctid` 是物理元组地址，不是持久行标识。它可能在 `UPDATE`、`VACUUM FULL`、`CLUSTER` 或任何表重写后变化；绝不能把它保存为应用主键，也不能用旧值授权或修改某一行。

实现复制了执行器结构，并使用 PostgreSQL 9.x 时代的堆访问 API。它没有当前兼容矩阵，跨大版本可能失效。应把 `ctidscan` 视为规划器扩展示例代码，用普通扫描核对正确性，且不要在生产集群上预加载。
