## 用法

来源：

- [1.1.1 README](https://github.com/supervc-stack/VectorChord/blob/1.1.1/README.md)
- [Control and dependency](https://github.com/supervc-stack/VectorChord/blob/1.1.1/vchord.control)
- [Preload requirement](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/lib.rs)
- [1.1.1 SQL objects](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/install/vchord--1.1.1.sql)
- [Query settings](https://github.com/supervc-stack/VectorChord/blob/1.1.1/src/index/gucs.rs)
- [1.1.1 migration](https://github.com/supervc-stack/VectorChord/blob/1.1.1/sql/upgrade/vchord--1.1.0--1.1.1.sql)
- [1.1.1 release notes](https://github.com/supervc-stack/VectorChord/releases/tag/1.1.1)

`vchord` 使用 pgvector 的类型，为 PostgreSQL 增加近似向量索引，提供基于分区的 `vchordrq` 和基于图的 `vchordg` 访问方法。扩展依赖 `vector`，要求共享预加载，创建时需要超级用户权限。

### 创建与查询索引

将库加入已有预加载列表，保留其他条目，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'vchord'
```

```sql
CREATE EXTENSION vchord CASCADE;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
INSERT INTO items(embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
CREATE INDEX items_embedding_idx ON items
USING vchordrq (embedding vector_l2_ops);

SELECT id FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
SELECT vchordrq_prewarm('items_embedding_idx'::regclass);
```

`vector_l2_ops` 对应 `<->`，`vector_ip_ops` 对应 `<#>`，`vector_cosine_ops` 对应 `<=>`。内积运算符返回负值，以适配索引升序扫描。图访问方法可使用同样的运算符类，应按工作负载选择索引方案：

```sql
CREATE INDEX items_embedding_graph_idx ON items
USING vchordg (embedding vector_l2_ops);
```

### 范围查询与调优

扩展为范围检索提供显式球体谓词：

```sql
SELECT id FROM items
WHERE embedding <<->> sphere('[1,2,3]'::vector, 0.5);

SET vchordrq.probes = '100';
SET vchordrq.epsilon = 1.9;
SET vchordg.ef_search = 64;
```

`<<->>`、`<<#>>` 和 `<<=>>` 分别是 L2、内积和余弦度量的球体谓词。探测数量取决于分区布局，应使用代表性数据调优。epsilon 设置控制重排序的权衡，图检索设置控制候选搜索范围。两种索引方法都是近似检索，选择参数前应检查召回率、过滤条件和查询计划。

### 1.1.1 的量化接口

`rabitq8` 和 `rabitq4` 保存量化向量。`quantize_to_rabitq8` 与 `quantize_to_rabitq4` 接受 `vector` 或 `halfvec`。1.1.1 为这两种量化类型新增 `dequantize_to_vector` 和 `dequantize_to_halfvec` 重载：

```sql
SELECT dequantize_to_vector(quantize_to_rabitq8('[1,2,3]'::vector));
SELECT dequantize_to_halfvec(quantize_to_rabitq4('[1,2,3]'::halfvec));
```

量化会损失精度，反量化返回近似结果。本次发布还替换了量化实现。安装匹配的库和 SQL 文件后，先为预加载库重启服务，再更新每个数据库中的扩展：

```sql
ALTER EXTENSION vchord UPDATE TO '1.1.1';
```

1.1.0 到 1.1.1 的脚本新增这四个转换重载，未声明索引格式迁移。更早版本的升级要求取决于起始版本。索引构建和预热会消耗资源，应按数据集大小安排。`vchordg_prewarm` 是对应的图索引辅助函数。control 声明扩展可重定位，在非默认模式中安装时，应限定扩展对象的模式或将安装模式加入搜索路径。
