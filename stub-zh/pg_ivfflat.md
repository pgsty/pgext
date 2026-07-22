## 用法

来源：

- [官方仓库 README](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/README.md)
- [PostgreSQL 扩展实现](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/lib.rs)
- [向量类型实现](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/vector_type.rs)
- [距离操作符实现](https://github.com/SteveLauC/pg_ivfflat/blob/691139bc07f57bbff20a40144b913d3cfd84a05d/pg_extension/src/operator.rs)

`pg_ivfflat` 0.1.0 是一个用于探索 IVFFlat 索引访问方法的实验性 pgrx 项目。固定版本的 PostgreSQL 扩展实际上没有注册任何索引访问方法：它只提供自定义 `vector` 类型和余弦距离操作符。不要把这个 API 与 `pgvector` 混淆，也不要因为软件包名就假定存在 IVFFlat 索引。

### 可用流程

```sql
CREATE EXTENSION pg_ivfflat;

CREATE TABLE items (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    embedding vector(3)
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id, embedding <=> '[3,1,2]'::vector(3) AS cosine_distance
FROM items
ORDER BY cosine_distance;
```

`vector(n)` 类型把 JSON 数组输入存储为 `float8` 值；存在类型修饰符时，维度范围限制为 1–65,535。其文本输出为紧凑 JSON 数组。`<=>` 接受两个向量，以 `double precision` 返回余弦距离。

### 实验边界

此源码快照没有 `CREATE INDEX ... USING ivfflat` 流程、操作符类、`lists` 选项、probes 设置或近似最近邻路径。上面的查询只是逐行计算距离后进行普通排序；应通过 `EXPLAIN` 检查计划，并预期全表工作。

Cargo 清单只启用 PostgreSQL 17，并使用 pgrx 0.12.9。通用名称 `vector` 和 `<=>` 与其他向量扩展的 API 重叠，在同一模式中安装可能发生冲突。仓库里的独立 `ivfflat` Rust crate 是内存实验，并未在这里暴露为 PostgreSQL 索引。应把此版本视为开发代码，而不是生产向量搜索实现。
