

## 用法

> [arraymath: PostgreSQL 的逐元素数组运算](https://github.com/pramsey/pgsql-arraymath)

提供整数、浮点数和数值数组的逐元素运算符和工具函数。

```sql
CREATE EXTENSION arraymath;
```

### 运算符

所有运算符以 `@` 为前缀，表示逐元素操作。支持数组对数组（相同长度或循环扩展）以及数组对标量。

| 运算符 | 说明 | 返回类型 |
|---|---|---|
| `@=` | 逐元素相等 | `boolean[]` |
| `@<` | 逐元素小于 | `boolean[]` |
| `@>` | 逐元素大于 | `boolean[]` |
| `@<=` | 逐元素小于等于 | `boolean[]` |
| `@>=` | 逐元素大于等于 | `boolean[]` |
| `@+` | 逐元素加法 | 同类型 |
| `@-` | 逐元素减法 | 同类型 |
| `@*` | 逐元素乘法 | 同类型 |
| `@/` | 逐元素除法 | 同类型 |

### 函数

| 函数 | 说明 |
|---|---|
| `array_sum(anyarray)` | 所有元素之和 |
| `array_avg(anyarray)` | 所有元素的平均值 |
| `array_min(anyarray)` | 最小元素 |
| `array_max(anyarray)` | 最大元素 |
| `array_median(anyarray)` | 中位数元素 |
| `array_sort(anyarray)` | 升序排序 |
| `array_rsort(anyarray)` | 降序排序 |

### 示例

```sql
-- 数组与标量
SELECT ARRAY[1,2,3,4] @< 4;             -- {t,t,t,f}
SELECT ARRAY[3.4,5.6,7.6] @* 8.1;       -- {27.54,45.36,61.56}

-- 数组与数组（较短数组循环扩展）
SELECT ARRAY[1,2,3,4,5,6] @* ARRAY[1,2]; -- {1,4,3,8,5,12}
SELECT ARRAY[1,2,3] @= ARRAY[3,2,1];     -- {f,t,f}

-- 工具函数
SELECT array_sort(ARRAY[9,1,8,2,7]);     -- {1,2,7,8,9}
SELECT array_sum(ARRAY[1,2,3,4,5]);      -- 15
SELECT array_median(ARRAY[1,2,3,4,5]);   -- 3
```
