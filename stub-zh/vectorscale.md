## 用法

来源：

- [0.9.1 README](https://github.com/timescale/pgvectorscale/blob/0.9.1/README.md)
- [Control and dependency](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/vectorscale.control)
- [0.9.1 migration SQL](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/sql/vectorscale--0.9.0--0.9.1.sql)
- [0.9.1 security and upgrade notes](https://github.com/timescale/pgvectorscale/releases/tag/0.9.1)

`vectorscale` 为 pgvector 增加 StreamingDiskANN 近似向量索引，`diskann` 访问方法支持 L2、内积、余弦距离和标签过滤。扩展依赖 `vector`，创建时需要超级用户权限。0.9.1 校验向量类型、维度及存储值布局，修复可能导致崩溃、内存泄露和越界写入的问题。

### 核心流程

声明明确的向量维度，并选择与查询距离相匹配的运算符类：

```sql
CREATE EXTENSION vectorscale CASCADE;
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    contents text,
    embedding vector(3),
    labels smallint[]
);
INSERT INTO documents(contents, embedding, labels)
VALUES ('PostgreSQL search', '[1,2,3]', ARRAY[1,3]::smallint[]);
CREATE INDEX documents_diskann ON documents
USING diskann (embedding vector_cosine_ops, labels);

SELECT id, contents FROM documents
WHERE labels && ARRAY[1]::smallint[]
ORDER BY embedding <=> '[1,2,3]'::vector
LIMIT 10;
```

`vector_l2_ops` 对应 `<->`，`vector_ip_ops` 对应 `<#>`，`vector_cosine_ops` 对应 `<=>`。标签使用 `smallint[]`，`&&` 表示与任一请求标签重叠。也支持普通 WHERE 条件，但严格过滤可能减少返回数量，应结合实际工作负载验证。

### 调优与排序

`diskann.query_search_list_size` 控制图检索额外候选数量，默认 100；`diskann.query_rescore` 控制精确重评分数量，默认 50，0 表示禁用：

```sql
SET diskann.query_search_list_size = 200;
SET diskann.query_rescore = 100;
```

构建选项包括 `storage_layout`、`num_neighbors`、`search_list_size` 和 `num_dimensions`，默认使用压缩的内存优化存储。增加维护内存前，应考虑并发构建及数据集大小。此版本的并行构建要求支持的压缩布局，且不支持标签列。

DiskANN 采用宽松的距离排序。需要严格排序时，可对物化结果集再次排序；这只能重排已检索的候选，不能将近似检索变成全量精确搜索。空向量不建立索引，空标签视为空数组，标签数组中的空元素被忽略。不支持在 UNLOGGED 表上创建索引。

### 升级到 0.9.1

安装匹配的扩展文件后，更新每个数据库中的扩展：

```sql
ALTER EXTENSION vectorscale UPDATE TO '0.9.1';
```

升级会将运算符类绑定到 pgvector 的实际安装模式，并检查已有绑定。若已有运算符类指向错误的类型或运算符，升级会中止；应依照发布说明删除受影响的运算符类并重新创建扩展对象，同时处理依赖索引。

**DiskANN 现在要求列类型具有有效的 `vector(N)` 维度。** 没有维度约束的向量列无法建立索引。持久化维度无效的旧索引会在扫描、插入和清理时报告错误。修正列类型后，必须**删除并重新创建受影响的索引**，`REINDEX` 无法修复这一问题。有效索引无需仅因安装 0.9.1 就全面重建。扩展创建后，SQL 对象不可重定位。
