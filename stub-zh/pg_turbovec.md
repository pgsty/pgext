## 用法

来源：

- [v2.10.3 README](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/README.md)
- [v2.10.3 变更日志](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/CHANGELOG.md)
- [升级矩阵](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/docs/UPGRADING.md)
- [控制文件](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/pg_turbovec.control)
- [过滤指南](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/docs/FILTERING.md)

`pg_turbovec` 2.10.3 提供 `turbovec.vector` 类型和紧凑向量索引，并用堆中的原始向量对候选项重新排序。平坦搜索扫描量化编码，IVF 则进一步只搜索选定的单元；除非候选集覆盖所有真实近邻，两者都属于近似搜索。**任何 1.x 索引升级到 2.x 都必须重建。**

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

### 从 1.x 升级到 2.10.3

安排维护窗口，保留堆中的向量以及经过验证的备份。安装匹配的新二进制文件、更新 SQL 扩展，然后重启 PostgreSQL，确保所有后端使用新库：

```sql
ALTER EXTENSION pg_turbovec UPDATE TO '2.10.3';
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
