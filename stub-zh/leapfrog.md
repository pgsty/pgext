## 用法

来源：

- [官方 README](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/README.md)
- [控制文件](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/leapfrog.control)
- [扩展 SQL](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/leapfrog--0.0.3.sql)

`leapfrog` 0.0.3 是面向 PostgreSQL 18 的实验性规划器扩展。它将普通 B-tree 索引组合为有序交集，使符合条件、带排序与条数限制的过滤查询可以减少专用索引。其他 PostgreSQL 大版本不受支持。

### 启用与检查

由超级用户安装 SQL 辅助函数，并在规划查询的会话中加载动态库：

```sql
CREATE EXTENSION leapfrog;
LOAD 'leapfrog';
SELECT leapfrog_version(), leapfrog_is_loaded();

CREATE TABLE orders (
    id bigint PRIMARY KEY,
    customer_id bigint NOT NULL,
    product_id bigint NOT NULL
);
CREATE INDEX ON orders (customer_id, id);
CREATE INDEX ON orders (product_id, id);
ANALYZE orders;
EXPLAIN
SELECT * FROM orders
WHERE customer_id = 100 AND product_id = 200
ORDER BY id DESC LIMIT 100;
```

当估算成本有利时，规划器可能选择 `Custom Scan (Index Merge Intersect)`。空的演示表通常仍使用原生计划。外侧索引以排序列结尾，索引前导列必须具备等值条件。每个附加索引都必须提供独立的等值约束或部分索引谓词。

### 加载与设置

`CREATE EXTENSION` 注册 SQL 辅助函数；之后每个新会话仍需加载动态库。`LOAD 'leapfrog'` 只作用于当前会话。管理员可以通过 `session_preload_libraries` 为新会话预加载，或通过 `shared_preload_libraries` 为所有后端预加载，后者需要重启服务器。

| 设置 | 默认值 | 用途 |
| --- | --- | --- |
| `leapfrog.enabled` | on | 提供自定义执行路径 |
| `leapfrog.force` | off | 为诊断优先选择符合条件的自定义路径 |
| `leapfrog.max_indexes` | 4 | 每次交集的最大索引流数量，范围为 2 至 16 |
| `leapfrog.min_indexes` | 2 | 开发者使用的最小索引流数量控制 |
| `leapfrog.skip_after` | 64 | 初始前跳阈值，0 禁用前跳 |

`leapfrog.force` 可能选出更慢的计划。修改默认值前应检查实际执行计划与耗时。普通规划器仍保留其他执行路径。

### 限制与维护

排序列必须是非空整数：`smallint`、`integer` 或 `bigint`。索引列必须采用升序，索引格式须为 nbtree 版本 4；查询的降序通过反向扫描实现。不支持表达式索引列，也不支持前导等值列上的范围或列表条件。从旧版 PostgreSQL 迁移而来的索引可能需要重建，才能满足 nbtree 格式要求。

重复键的 TID 分组受 `work_mem` 限制，超出限制的分组退回逐行条件检查。未结束的扫描可能持续固定索引页，从而延迟清理。该扩展不创建应用数据表，也不会替换现有索引。

`DROP EXTENSION` 不会卸载已进入后端的动态库。可通过 `SET leapfrog.enabled = off` 禁用其规划器行为，并在替换动态库后重新连接会话。应使用正规安装或软件包替换流程；原地覆盖已加载的动态库文件可能导致使用它的进程崩溃。
