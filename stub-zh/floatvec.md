

## 用法

> [floatvec: PostgreSQL 数组的逐元素算术运算](https://github.com/pjungwir/floatvec)

提供将数组视为向量进行基本算术运算的函数。支持 `SMALLINT`、`INTEGER`、`BIGINT`、`REAL` 和 `DOUBLE PRECISION` 类型。

```sql
CREATE EXTENSION floatvec;
```

### 函数

每个函数接受两个相同长度的数组，或一个数组和一个标量，或一个标量和一个数组。两个参数必须是相同类型。

| 函数 | 说明 |
|---|---|
| `vec_add(a, b)` | 逐元素加法 |
| `vec_sub(a, b)` | 逐元素减法 |
| `vec_mul(a, b)` | 逐元素乘法 |
| `vec_div(a, b)` | 逐元素除法 |
| `vec_pow(a, b)` | 逐元素乘方 |

### 示例

```sql
-- 数组 + 数组
SELECT vec_add(ARRAY[1,2,3], ARRAY[10,20,30]);  -- {11,22,33}

-- 数组 * 标量
SELECT vec_mul(ARRAY[1.0, 2.0, 3.0], 2.0);     -- {2.0, 4.0, 6.0}

-- 标量 - 数组
SELECT vec_sub(10, ARRAY[1,2,3]);               -- {9,8,7}
```
