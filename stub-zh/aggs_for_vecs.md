

## 用法

> [aggs_for_vecs: 数组的聚合函数（向量/逐行模式）](https://github.com/pjungwir/aggs_for_vecs)

提供将数组视为向量的聚合函数，跨多行计算按位置的统计量。支持 `SMALLINT`、`INTEGER`、`BIGINT`、`REAL` 和 `DOUBLE PRECISION` 类型。

```sql
CREATE EXTENSION aggs_for_vecs;
```

### 聚合函数

| 函数 | 说明 |
|---|---|
| `vec_to_count(anyarray)` | 每个位置的非空计数 |
| `vec_to_sum(anyarray)` | 每个位置的求和 |
| `vec_to_min(anyarray)` | 每个位置的最小值 |
| `vec_to_max(anyarray)` | 每个位置的最大值 |
| `vec_to_mean(anyarray)` | 每个位置的平均值（返回 `FLOAT[]`） |
| `vec_to_weighted_mean(values, weights)` | 每个位置的加权平均值 |
| `vec_to_var_samp(anyarray)` | 每个位置的样本方差 |
| `vec_to_first(anyarray)` | 每个位置的第一个非空值（配合 ORDER BY 使用） |
| `vec_to_last(anyarray)` | 每个位置的最后一个非空值（配合 ORDER BY 使用） |
| `hist_2d(x, y, ...)` | 二维直方图 |
| `hist_md(vals, indexes, ...)` | N 维直方图 |

### 非聚合函数

| 函数 | 说明 |
|---|---|
| `vec_add(l, r)` | 逐元素加法 |
| `vec_sub(l, r)` | 逐元素减法 |
| `vec_mul(l, r)` | 逐元素乘法 |
| `vec_div(l, r)` | 逐元素除法 |
| `vec_elements(array, indexes)` | 按指定索引选取元素 |
| `pad_vec(array, length)` | 用 NULL 扩展数组至指定长度 |
| `vec_coalesce(array, default)` | 将 NULL 替换为默认值 |
| `vec_trim_scale(numeric[])` | 去除 NUMERIC 元素的尾部零 |
| `vec_without_outliers(vals, mins, maxs)` | 将异常值替换为 NULL |

### 示例

```sql
-- 跨行按位置取最小值
SELECT vec_to_min(vals) FROM (VALUES
  (ARRAY[1,2,3,4]),
  (ARRAY[5,0,-5,0]),
  (ARRAY[3,6,0,9])
) AS t(vals);
-- {1,0,-5,0}

-- 按位置求平均值
SELECT vec_to_mean(vals) FROM my_table;
```
