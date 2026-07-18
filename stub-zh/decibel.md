## 用法

来源：

- [项目 README](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/README.md)
- [0.0.1 版安装 SQL](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/decibel--0.0.1.sql)
- [C 语言转换实现](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/decibel.c)

`decibel` 0.0.1 定义了一种 PostgreSQL 基础类型：输入和显示使用分贝值，内部则保存双精度线性值。输入路径执行 10^(dB/10)，输出路径使用 10*log10(value) 将存储值转换回分贝。

### 创建并转换数值

```sql
CREATE EXTENSION decibel;

SELECT 65.23::decibel AS level_db,
       pascals(65.23::decibel) AS linear_value;
```

该扩展提供 `decibelpascal(float8)`、`pascaldecibel(float8)` 和 `pascals(decibel)`，并支持 `decibel` 与常见数值类型之间的类型转换。它还为该类型定义了比较、算术操作符以及 `sum`、`avg`、`min`、`max` 聚合。

### 语义与兼容性

两个 `decibel` 值之间的加法、减法和平均值计算基于内部线性表示；与 `float8` 相加或相减则调整显示的分贝值，解释结果时必须区分这两种语义。API 将线性量称为帕斯卡，但实现中没有参考声压常量，只执行上述功率比转换。该项目是 C 扩展，仅有一个较早的 0.0.1 源码快照，也没有当前兼容矩阵，因此应在目标 PostgreSQL 版本上验证构建和数值行为。
