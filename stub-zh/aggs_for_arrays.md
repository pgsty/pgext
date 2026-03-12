

## 用法

> [aggs_for_arrays: 单数组的类聚合函数（逐列模式）](https://github.com/pjungwir/aggs_for_arrays)

提供对单个数组输入计算统计量的函数。支持 `SMALLINT`、`INTEGER`、`BIGINT`、`REAL` 和 `DOUBLE PRECISION` 类型。

```sql
CREATE EXTENSION aggs_for_arrays;
```

### 函数

| 函数 | 说明 |
|---|---|
| `array_to_hist(values T[], start T, width T, count INT)` | 计算直方图分桶计数 |
| `array_to_hist_2d(x[], y[], ...)` | 二维直方图 |
| `array_to_mean(values T[])` | 数组元素的均值 |
| `array_to_median(values T[])` | 中位数（输入无需排序） |
| `sorted_array_to_median(values T[])` | 中位数（输入已排序） |
| `array_to_mode(values T[])` | 数组元素的众数 |
| `sorted_array_to_mode(values T[])` | 众数（输入已排序） |
| `array_to_percentile(values T[], pct FLOAT)` | 百分位数（0 到 1） |
| `sorted_array_to_percentile(values T[], pct FLOAT)` | 百分位数（输入已排序） |
| `array_to_percentiles(values T[], pcts FLOAT[])` | 多个百分位数 |
| `sorted_array_to_percentiles(values T[], pcts FLOAT[])` | 多个百分位数（输入已排序） |
| `array_to_max(values T[])` | 最大元素 |
| `array_to_min(values T[])` | 最小元素 |
| `array_to_min_max(values T[])` | `{min, max}` 元组 |
| `array_to_skewness(values T[])` | 偏度 |
| `array_to_kurtosis(values T[])` | 峰度 |

### 示例

```sql
-- 10 个分桶的直方图，起始值 0，桶宽 10
SELECT array_to_hist(ARRAY[5,15,25,35,45], 0, 10, 5);
-- {1,1,1,1,1}

-- 均值和中位数
SELECT array_to_mean(ARRAY[1,2,3,4,5]);   -- 3.0
SELECT array_to_median(ARRAY[1,2,3,4,5]); -- 3.0

-- 百分位数
SELECT array_to_percentile(ARRAY[1,2,3,4,5,6,7,8,9,10], 0.95);

-- 一次计算多个百分位数
SELECT array_to_percentiles(ARRAY[1,2,3,4,5], ARRAY[0.25, 0.5, 0.75]);
```
