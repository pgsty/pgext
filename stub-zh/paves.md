## 用法

来源：

- [Official README](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/README.md)
- [Control file](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/extension/paves/paves.control)
- [Version 0.1 SQL](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/extension/paves/sql/paves--0.1.sql)
- [Configuration and limits](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/docs/07-reference.md)
- [Index maintenance](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/docs/02-index-build.md)

`paves` 是面向 PostgreSQL 14 的过滤向量检索研究扩展，将标量过滤条件与 L2 向量距离排序结合，在 CPU 和可选的 GPU 路径之间选择执行方式。索引是静态快照，表数据变化后需要重建。

### 核心流程

准备与 PostgreSQL 14 匹配的库及 `vector` 依赖，将 `paves` 加入预加载列表并重启 PostgreSQL，然后由管理员启用扩展：

```conf
shared_preload_libraries = 'paves'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION paves;

CREATE TABLE items (
    id bigint PRIMARY KEY,
    embedding vector(3),
    price integer,
    category integer
);
INSERT INTO items VALUES
    (1, '[0.1,0.2,0.3]', 50, 3),
    (2, '[0.4,0.5,0.6]', 80, 7),
    (3, '[0.8,0.9,1.0]', 150, 1);

SELECT * FROM paves_capacity('items');
SELECT paves_build('items');

EXPLAIN SELECT id FROM items
WHERE price < 100 AND category IN (3, 7)
ORDER BY embedding <-> '[0.1,0.2,0.3]'::vector
LIMIT 10;
```

索引使用表中的第一个 `vector` 列。向量维度必须一致；空向量会跳过，空数据构建会报错。仅支持使用 `<->` 的单项 L2 距离排序且带有 `LIMIT` 的查询。其他查询仍走 PostgreSQL 常规路径，应检查执行计划确认实际路径。

### 接口与调优

- `paves_build(regclass)` 扫描表，原子替换索引文件，返回已索引的行数。
- `paves_capacity(regclass)` 报告维度、行数及标量列上限、磁盘与内存估算和 GPU 可用性。
- `paves_feedback()` 提供路由反馈与 CPU/GPU 负载计数器。
- `paves.enable` 控制自定义路径；`paves.force_strategy` 可选 `auto`、`brute`、`hnsw`、`gpu_brute` 或 `gpu_hnsw`。
- `paves.ef_search` 控制图搜索宽度；`paves.threads` 控制 CPU 扫描线程数。
- `paves.build_m`、`paves.build_ef_construction` 和 `paves.build_threads` 控制索引构建。

CUDA 是可选依赖，无 CUDA 时可使用纯 CPU 构建。GPU 后台进程需要预加载并重启。`paves.gpu_enable` 约束自动路由，`paves.gpu_batch_max` 与 `paves.gpu_batch_wait_us` 调整后台进程的微批处理。尺寸超限或后台容量不可用时，可能回退到 CPU 路径。

### 维护与限制

写入后须重新执行 `paves_build()`。回表可见性检查不能维护索引，也不能保证变化数据的召回。索引文件位于 `$PGDATA/paves/`；`DROP TABLE` 后残留的文件需要手工清理。重建期间新旧文件会短暂共存，旧映射可能保留到会话结束。

该实现面向 PostgreSQL 14，每个索引常驻一张 GPU；不能据此推断支持其他主版本或多 GPU 分片。近似搜索路径会在召回与计算量之间取舍，应针对实际负载验证正确性和资源限制，不能将公开基准结果当作普遍性能保证。已核对的源码修订未声明许可证。
