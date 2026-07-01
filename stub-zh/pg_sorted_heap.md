


## 用法

- 来源：[pg_sorted_heap README](https://github.com/skuznetsov/pg_sorted_heap)，[stable API](https://github.com/skuznetsov/pg_sorted_heap/blob/main/docs/api-stable.md)，[SQL API](https://github.com/skuznetsov/pg_sorted_heap/blob/main/docs/api.md)，[control file](https://github.com/skuznetsov/pg_sorted_heap/blob/main/pg_sorted_heap.control)

`pg_sorted_heap` 增加了 `sorted_heap` 表访问方法、按页 zone-map 裁剪、维护辅助函数、内置 `svec`/`hsvec` 向量类型、集成进 planner 的 `sorted_hnsw` 索引访问方法，以及稳定版 GraphRAG 封装。上游文档说明当前发布接口支持 PostgreSQL 16、17 和 18。

### Sorted Heap 表

在带有主键的表上使用 `USING sorted_heap`。通过 COPY 路径批量导入时会按主键排序；compaction 会对已有行做全局排序，并重建 zone map：

```sql
CREATE EXTENSION pg_sorted_heap;

CREATE TABLE events (
  ts timestamptz,
  src text,
  data jsonb,
  PRIMARY KEY (ts, src)
) USING sorted_heap;

COPY events FROM '/path/to/events.csv';

SELECT sorted_heap_compact('events'::regclass);

EXPLAIN (ANALYZE, BUFFERS)
SELECT *
FROM events
WHERE ts BETWEEN '2026-01-01' AND '2026-01-02'
  AND src = 'sensor-42';
```

README 描述了针对主键谓词由 planner 注入的 `SortedHeapScan` 路径，以及在 heap block 粒度上的 zone-map 裁剪。

### 维护与观测

稳定维护函数包括：

```sql
SELECT sorted_heap_compact('events'::regclass);
CALL sorted_heap_compact_online('events'::regclass);

SELECT sorted_heap_merge('events'::regclass);
CALL sorted_heap_merge_online('events'::regclass);

SELECT sorted_heap_rebuild_zonemap('events'::regclass);
SELECT sorted_heap_zonemap_stats('events'::regclass);
```

分区辅助函数作用于父表下具体的 sorted-heap 叶子分区：

```sql
SELECT * FROM sorted_heap_partition_status('events_parent'::regclass);
SELECT * FROM sorted_heap_partition_maintenance_plan('events_parent'::regclass, 'compact');
SELECT * FROM sorted_heap_compact_partitions('events_parent'::regclass);
```

### 向量搜索

稳定向量 API 包括 float32 向量 `svec(dim)`、float16 向量 `hsvec(dim)`，以及 `sorted_hnsw` 索引访问方法：

```sql
CREATE TABLE documents (
  id bigserial PRIMARY KEY,
  embedding svec(384),
  content text
);

CREATE INDEX documents_embedding_idx
ON documents USING sorted_hnsw (embedding)
WITH (m = 16, ef_construction = 200);

SET sorted_hnsw.ef_search = 96;

SELECT id, content
FROM documents
ORDER BY embedding <=> '[0.1,0.2,0.3]'::svec
LIMIT 10;
```

如果希望基表存储更紧凑，可以使用 `hsvec` 和对应的 operator class：

```sql
CREATE TABLE documents_compact (
  id bigserial PRIMARY KEY,
  embedding hsvec(384),
  content text
);

CREATE INDEX documents_compact_embedding_idx
ON documents_compact USING sorted_hnsw (embedding hsvec_cosine_ops)
WITH (m = 16, ef_construction = 200);
```

共享解码图缓存由 `sorted_hnsw.shared_cache` 控制。上游示例特别说明，使用它需要预加载该扩展：

```conf
shared_preload_libraries = 'pg_sorted_heap'
```

```sql
SET sorted_hnsw.shared_cache = on;
```

### GraphRAG

稳定版 fact-shaped GraphRAG 入口要求事实数据按 `(entity_id, relation_id, target_id)` 聚簇，或先注册一组别名映射：

```sql
CREATE TABLE facts (
  entity_id int4,
  relation_id int2,
  target_id int4,
  embedding svec(384),
  payload text,
  PRIMARY KEY (entity_id, relation_id, target_id)
) USING sorted_heap;

CREATE INDEX facts_embedding_idx
ON facts USING sorted_hnsw (embedding)
WITH (m = 24, ef_construction = 200);

SET sorted_hnsw.ef_search = 128;

SELECT *
FROM sorted_heap_graph_rag(
  'facts'::regclass,
  '[0.1,0.2,0.3]'::svec,
  relation_path := ARRAY[1, 2],
  ann_k := 64,
  top_k := 10,
  score_mode := 'path'
);
```

如果事实表使用了不同的列名，可以注册一次替代映射：

```sql
SELECT sorted_heap_graph_register(
  'facts_alias'::regclass,
  entity_column := 'src_id',
  relation_column := 'edge_type',
  target_column := 'dst_id',
  embedding_column := 'vec',
  payload_column := 'body'
);
```

对于按路由或租户分片的事实表，可使用 `sorted_heap_graph_route(...)`，并通过 `sorted_heap_graph_route_plan(...)` 查看路由计划。

### 稳定 GUC

- `sorted_heap.enable_scan_pruning`：启用 sorted-heap custom scan 裁剪；默认 `on`。
- `sorted_heap.vacuum_rebuild_zonemap`：在 `VACUUM` 期间重建 zone map；默认 `off`。
- `sorted_heap.lazy_update`：推迟即时 zone-map 更新维护；默认 `off`。
- `sorted_hnsw.ef_search`：运行时 HNSW 搜索宽度；默认 `64`。
- `sorted_hnsw.shared_cache`：预加载后使用的共享解码图缓存；默认 `on`。
- `sorted_hnsw.sq8`：SQ8 解码缓存表示；默认 `on`。
- `sorted_hnsw.build_sq8`：低内存索引构建模式；默认 `off`。

### 注意事项

- `sorted_heap.lazy_update = on` 会牺牲 scan pruning，换取 update-heavy 工作负载下更快的更新，直到 compaction 或 merge 恢复裁剪效果。
- `sorted_hnsw.shared_cache` 应与 `shared_preload_libraries = 'pg_sorted_heap'` 配合使用。
- planner 集成的 `sorted_hnsw` 有序扫描需要 `LIMIT`；SQL API 说明，在没有 limit 或 `LIMIT > sorted_hnsw.ef_search` 时不会选择它们。
- 底层 GraphRAG 与旧式/手工 ANN 辅助函数仍有文档，但稳定的应用侧 API 是 `docs/api-stable.md` 中的精简接口。
