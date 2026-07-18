## 用法

来源：

- [上游 README](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/README.md)
- [扩展 control 文件](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/numpgsql.control)
- [数组数学 SQL](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/src/math/math.sql)
- [随机数组 SQL](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/src/random/random.sql)

`numpgsql` `0.0.1` 版为 PostgreSQL 数值数组提供受 NumPy 启发的运算。其 C++/Boost 实现包含逐元素算术和超越函数、向量归约与聚合、索引/布尔切片、排序/反转，以及可设种子的随机数组和分布。

### 示例

```sql
CREATE EXTENSION numpgsql;
SELECT cos(ARRAY[5, 1, 6, 4]::real[]);
SELECT ARRAY[10, 20, 30, 40] @ ARRAY[2, 4];
SELECT random_seed(42);
SELECT random_randn(2, 2);
```

项目把广播、类型转换、按轴聚合和文档列为未完成项，且自 2019 年后没有更新。许多函数会重载通用名称和运算符，结果类型继承 PostgreSQL 数组元素语义，包括整数截断/溢出行为。若未针对形状、类型、空值、溢出和可复现性要求测试，不应把它视为兼容 NumPy。
