## 用法

来源：

- [Threads per query](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/PRODUCTION.md)
- [v2.14.1 README](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/README.md)
- [v2.14.1 变更日志](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/CHANGELOG.md)
- [升级矩阵](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/UPGRADING.md)
- [控制文件](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/pg_turbovec.control)
- [过滤指南](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/FILTERING.md)

`pg_turbovec` 2.14.1 提供 `turbovec.vector` 类型和紧凑向量索引，并用堆中的原始向量对候选项重新排序。平坦搜索扫描量化编码，IVF 则进一步只搜索选定的单元；除非候选集覆盖所有真实近邻，两者都属于近似搜索。**任何 1.x 索引升级到 2.x 都必须重建。**

### 创建和查询向量

```sql
CREATE EXTENSION pg_turbovec;
SET search_path = public, turbovec;

CREATE TABLE items (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding turbovec.vector CHECK (turbovec.vector_dims(embedding) = 8)
);
INSERT INTO items (embedding) VALUES
  ('[1,2,3,4,5,6,7,8]'),
  ('[2,1,3,5,4,7,6,8]'),
  ('[8,7,6,5,4,3,2,1]');
CREATE INDEX items_embedding_idx ON items
USING turbovec (embedding turbovec.vec_cosine_ops)
WITH (bit_width = 4);

SELECT id, embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector AS distance
FROM items
ORDER BY embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector
LIMIT 3;
```

索引中的向量必须具有一致的维度，且维度为 8 的倍数；该类型最多接受 16,000 个坐标。示例使用八维向量，因此能够建立索引。将 `turbovec.vec_cosine_ops` 与 `<=>` 配合使用，将 `turbovec.vec_ip_ops` 与 `<#>` 配合使用；`<->` 和 `<+>` 还提供精确距离运算。

### 选择索引并调优查询

- `lists = 0` 是默认的平坦量化扫描；候选项重新排序并不保证任意数据集都达到精确召回率。
- `WITH (lists = N)` 启用 IVF。应使用足够的代表性数据训练，根据语料选择单元数量，并与精确搜索基线比较召回率；增加单元数量本身不保证查询更快。
- 默认 `bit_width = 4`，也支持二位、三位量化。`bit_width = 1` 使用经过均值中心化的符号二值量化，支持 IVF，并从堆中重新排序；它与 TurboQuant 是不同的方案。
- `WITH (graph = true)` 已弃用。新索引优先选择平坦或 IVF；已有图索引应查阅上游迁移指南。

无需预加载也能建立索引和查询，但上游建议将库加入 `shared_preload_libraries`，使调优 GUC 正确注册。应与已有条目合并并重启 PostgreSQL，不要替换其他必需的库：

```conf
shared_preload_libraries = 'pg_turbovec'
```

```sql
SELECT count(*) FROM pg_settings WHERE name LIKE 'turbovec.%';
SET turbovec.probes = 16;
SET turbovec.search_k = 64;
```

设置查询必须返回非零数量；SET 接受带点号的占位参数，并不证明扩展的 GUC 已注册。重要参数包括 `turbovec.probes`、`turbovec.search_k`、`turbovec.oversample`、`turbovec.iterative_scan` 和 `turbovec.cache_size_mb`。稳定的过滤条件可使用部分索引，也可使用文档所述的允许列表或迭代扫描路径；应将过滤结果与精确基线比较。原生分区的索引需要分别维护。

### 从 1.x 升级到 2.12.1

安排维护窗口，保留堆中的向量以及经过验证的备份。安装匹配的新二进制文件、更新 SQL 扩展，然后重启 PostgreSQL，确保所有后端使用新库：

```sql
ALTER EXTENSION pg_turbovec UPDATE TO '2.12.1';
```

2.0.0 将索引格式从 7 改为 8。旧索引无法原地读取：**恢复查询前，必须从堆数据重建每一个 TurboVec 索引**，包括每个分区上的索引。在 psql 中生成并执行重建语句：

```sql
SELECT format('REINDEX INDEX %I.%I;', n.nspname, c.relname)
FROM pg_class c
JOIN pg_am a ON a.oid = c.relam
JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE a.amname = 'turbovec';
\gexec
```

这里使用普通的阻塞式 REINDEX。如果选择并发重建，应在事务块和 DO 函数之外逐条执行顶层命令，并考虑新索引可用前旧格式扫描仍会失败。不能把 SQL 扩展更新成功当作完成了 1.x 迁移。

2.10.2 到 2.10.3 的补丁升级保留格式 8，本身不要求重建。更早版本之间的升级仍应遵循上游升级矩阵。即使补丁保持格式兼容，已经损坏的索引仍须修复；`turbovec.turbovec_check(regclass)` 会报告检测到的损坏及原因。

### 兼容性与维护

控制文件将对象固定在 `turbovec` 模式中，设置 `superuser = false`，且不支持重定位。上游覆盖 PostgreSQL 13–18，并在本版本将 PostgreSQL 19 标记为实验性；当前 Pigsty 包覆盖 14–18。替换二进制文件、预加载调优和跨索引格式迁移分别涉及重启与重建要求。2.10.3 并行执行冷后端的编码重排，不改变已有索引字节。应按实际语料规划索引构建内存、临时空间、WAL 和 vacuum 工作量，不要将上游测试数字当作性能保证。

### 升级到 2.11.0

2.11.0 在适用硬件和较大索引上采用分阶段的 2/4 位搜索，近似候选集合可能变化。同时修复了 2.10.3 及更早版本中的并发写入索引损坏和逐次扫描内存泄漏。对于来自 2.x 的健康格式 8 索引，安装匹配文件、更新扩展并重启即可；本次没有 SQL 接口或线格式变化，不要求例行 REINDEX。已发生的索引损坏仍需修复，1.x 索引必须重建的要求也没有改变。

### 2.12.0 的距离舍入变化

向量化精确距离内核改变八维及以上 vector 和 halfvec 输入的舍入结果。距离极近的邻居可能改变排序或跨越距离阈值，sparsevec 运算符不变。格式 8 和已有索引字节保持不变：更新 SQL 扩展并重新连接或重启以加载新库，无需例行重建索引。

表达式索引、部分索引谓词、存储生成列和物化视图可能保留旧的舍入结果。如果要求与新计算严格一致，应重建相关索引、刷新视图并重新计算存储的派生值。既有损坏仍需修复，所有 1.x 索引仍须重建。在空表上创建或被截断的 IVF 索引不会训练单元，应先装载数据，再按上游流程重建。

### 2.12.1 召回修复

2.12.1 将 turbovec 更新到 1.1.2，修复向量具有明显共同方向时分阶段搜索的召回损失。适用硬件上的候选集合可能变化，精确重排分数不变。本版还修复 Nix 构建，上游 pgrx 升至 0.19.3，CI 覆盖 PostgreSQL 19beta4。SQL 对象、GUC 和格式 8 均不变。更新扩展后重新连接以加载新库；健康的 2.x 索引无需 REINDEX。上述 1.x 重建与既有损坏修复要求仍适用。

### 2.13.0 限制扫描线程

`turbovec.scan_parallelism` 现用于限制 flat、IVF、ColBERT ANN 和 1-bit 扫描、`turbovec.knn()` 以及首次打开索引时的编码重排；已弃用的 graph 路径保持原状。在 2.13.0 中，默认值 0 最多使用四个核心，调低设置后在当前会话生效。调优时应权衡并发吞吐量与单条查询的延迟。

原本不生效的 `turbovec.search_concurrency` 设置已移除。对这一带点名称执行 SET 成功，可能只是 PostgreSQL 接受了占位参数，并不表示参数有效。2.13.0 还修复了外存 IVF 码本缓存的竞争问题。安装匹配文件、执行 `ALTER EXTENSION pg_turbovec UPDATE TO '2.13.0'`，再重新连接或重启。健康的格式 8 索引无需 REINDEX；上述 1.x 索引重建要求仍适用。

### 2.14.0 并发默认值与候选重排

`turbovec.scan_parallelism = 0` 现在默认使用单线程，以减少高并发下的 CPU 过度竞争；它并未消除各后端私有解码缓存的内存开销。低并发场景可以显式增大线程数，但应衡量查询延迟与整体吞吐量的取舍。

新增的 `turbovec.rerank_multiplier` 默认 16.0，范围为 1.0–1024.0；精确重排的候选窗口由请求的数量和 oversample 决定，不再按向量维数设置下限。候选集合可能变化，应重新验证召回率。安装匹配文件后执行 `ALTER EXTENSION pg_turbovec UPDATE TO '2.14.0'` 并重新连接以加载库及新默认值。格式仍为 8，健康的 2.x 索引无须例行 REINDEX；此前的 1.x 迁移与损坏修复要求不变。

### 2.14.1 的写入与 VACUUM 正确性修复

2.14.1 修复并发提交后的后端缓存过期、插入刷盘时丢失 IVF 墓碑、写事务未结束时 VACUUM 清空索引后旧行重新出现，以及退化 IVF 扫描忽略墓碑的问题。精确评分与格式 8 均未改变。安装匹配文件，执行 `ALTER EXTENSION pg_turbovec UPDATE TO '2.14.1'`，并回收长期保持的数据库连接，使其重新加载库与缓存。

两种历史状态需要执行一次 REINDEX：旧版中在 VACUUM 后又接受插入的 IVF 索引，以及写事务尚未结束时被 VACUUM 清空的任何索引。可在停止写入的表上执行 `VACUUM (INDEX_CLEANUP ON)`，再用 `turbovec.turbovec_check(regclass)` 比较估算的存活索引条目数与实际覆盖的堆表行数。计数须包含向量非空条件和部分索引谓词；索引条目多于覆盖行数时需要 REINDEX。其他健康的 2.x 索引无须例行重建；1.x 升级重建和已有损坏修复的要求仍然适用。
