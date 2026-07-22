## 用法

来源：

- [官方 README](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/README.md)
- [0.0.1 版 SQL 声明](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array--0.0.1.sql)
- [C 实现](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array.c)

`sum_array` 增加一个 C 函数，用于合计 PostgreSQL 数值数组，并始终返回 `double precision`。它接受元素类型为 `smallint`、`integer`、`bigint`、`real` 或 `double precision` 的数组；它不是跨行聚合函数。

### 核心流程

```sql
CREATE EXTENSION sum_array;

SELECT sum_array(ARRAY[1, 2, 3, 4, 5]::integer[]);
SELECT sum_array(ARRAY[1.1, 2.2, 3.3]::double precision[]);
```

SQL 函数被声明为 `IMMUTABLE` 与 `STRICT`。因此 NULL 数组返回 NULL，而非 NULL 的受支持数组会逐元素转换为 `float8`，并从零开始累加。

### 数值与 NULL 注意事项

实现没有检查 PostgreSQL 返回的逐元素 NULL 标志。不要传入包含 NULL 元素的数组，也不要依赖其当前偶然结果；应先显式移除或替换这些元素。不受支持的元素类型（包括 `numeric`）会报错。

所有整数输入都会转换为 `double precision`，因此足够大的 `bigint` 会失去精确性，而浮点合计保留通常的舍入、溢出、无穷和 NaN 行为。项目已停留在废弃的 `0.0.1` 版；PostgreSQL 内建的 `unnest()` 配合 `sum()` 维护更完善，也允许调用方显式选择 NULL 与结果类型语义。
