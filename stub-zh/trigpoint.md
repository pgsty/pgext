## 用法

来源：

- [发布元数据](https://api.pgxn.org/dist/trigpoint/0.5.0/META.json)
- [运维指南](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/README.md)
- [控制文件](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/trigpoint.control)
- [安装 SQL](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/trigpoint--0.5.sql)
- [更新日志](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/CHANGELOG.md)

`trigpoint` 为 PostgreSQL 16–18 上的 pgvector 提供向量索引访问方法。它在紧凑向量摘要上建立 HNSW 路由，并在读取堆表行之前使用校准下界或精确下界剪枝。发行版本 0.5.0 安装的 SQL 版本为 0.5，需要 pgvector 0.8.0 或更高版本。

### 核心用法

由管理员安装原生库，然后创建两个扩展：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION trigpoint;
CREATE TABLE items (id bigint, embedding vector(3));
CREATE INDEX items_embedding_idx
    ON items USING trigpoint (embedding vector_l2_ops);

SELECT id, embedding <-> '[1,2,3]'::vector AS distance
FROM items
ORDER BY embedding <-> '[1,2,3]'::vector
LIMIT 10;
```

`vector_l2_ops` 操作符类支持 L2 距离；`vector_cosine_ops` 通过 `<=>` 支持余弦距离。每个索引只能包含一个向量列，不支持内积、反向扫描、覆盖列和并行查询扫描。

### 搜索控制

默认的 `trigpoint.iterative_scan = relaxed_order` 分批扩展带过滤条件的扫描，后续批次可能返回距离更近的行。`strict_order` 丢弃不符合距离顺序的候选，以保持输出有序。需要精确的结果前缀时，还应启用 `trigpoint.exact`：

```sql
BEGIN;
SET LOCAL trigpoint.iterative_scan = strict_order;
SET LOCAL trigpoint.exact = on;
SELECT id
FROM items
ORDER BY embedding <-> '[1,2,3]'::vector
LIMIT 10;
COMMIT;
```

`trigpoint.max_scan_tuples` 默认为 20000；`trigpoint.max_probes` 默认为 0，允许访问全部分区。达到上限时可能返回不足指定数量的行，且不会报错。其他列的过滤条件在索引之上执行；不支持直接针对向量列的索引过滤条件。近似搜索的召回率取决于路由与校准设置。

### 内存与预加载

构建索引时，`maintenance_work_mem` 需要约为行数乘向量维度乘六字节，并需要足够的动态共享内存。PostgreSQL 17–18 的共享图像注册表不需要预加载。PostgreSQL 16 若要共享图像，需要将库加入现有预加载列表并重启：

```ini
shared_preload_libraries = 'trigpoint'
```

PostgreSQL 16 未预加载时仍可使用扩展，但每个后端会保留独立索引图像。`trigpoint_shared_available()` 可检查共享能力。超级用户参数 `trigpoint.max_shared_mb` 限制共享图像总量；`trigpoint.on_shared_cap = private` 允许达到上限后退回私有副本，而不是报错。

### 维护与访问权限

```sql
SELECT trigpoint_version();
SELECT trigpoint_info('items_embedding_idx'::regclass);
SELECT * FROM trigpoint_check('items_embedding_idx'::regclass, false);
SELECT * FROM trigpoint_shared_total();
VACUUM items;
```

`VACUUM` 记录删除，并在达到阈值时合并索引图像；`REINDEX` 重建索引。目录升级使用 `ALTER EXTENSION trigpoint UPDATE`，导出后恢复会重新构建索引。该版本说明了页格式 3，并测试了 PostgreSQL 17→18 升级，不能据此保证未来任意格式均兼容。

接受索引参数的辅助函数要求调用者能读取对应表；深度检查要求表所有权。`trigpoint_image()` 和 `trigpoint_route()` 会暴露向量或邻居信息，已撤销 `PUBLIC` 的执行权限。发布构建不包含修改页面的测试钩子。可额外启用 PostgreSQL 数据校验和检测损坏；结构检查器不能证明所有模型数值均正确。
