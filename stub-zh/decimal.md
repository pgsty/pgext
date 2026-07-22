## 用法

来源：

- [官方 README](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/README.md)
- [官方扩展 SQL](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal--1.0.sql)
- [官方 C 实现](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.c)
- [官方 control 文件](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.control)

`decimal` 是 2013 年的实验，把编译器 `_Decimal32` 和 `_Decimal64` 值暴露为固定长度 PostgreSQL 类型 `decimal32` 与 `decimal64`。其目的是比较固定十进制运算和 PostgreSQL `numeric`，并不是完整的生产级数值实现。

### 核心流程

```sql
CREATE EXTENSION decimal;

SELECT '3.14'::decimal32 + '2.00'::decimal32;
SELECT '12.3456'::decimal64 * '2'::decimal64;
SELECT round('12.3456'::decimal64, 2);

SELECT sum(v)
FROM (VALUES ('1.25'::decimal64), ('2.75'::decimal64)) AS x(v);
```

两个 `decimal32` 值的算术结果是 `decimal64`。两个 `decimal64` 值的算术结果也是 `decimal64`；唯一聚合为 `sum(decimal64)`。

### 转换与缺失接口

隐式转换接受常见整数类型、从 `float4` 到 `decimal32`、从 `float8` 到 `decimal64`、从 `numeric` 到任一十进制类型，以及从 `decimal32` 到 `decimal64`。`decimal64` 可转回 `numeric`；缩窄为 `decimal32` 需要显式转换。

扩展提供 `+`、`-`、`*`、`/` 与双参数 `round`，但没有比较/相等操作符、哈希支持、B-tree 操作符类、`avg`、typmod 精度/标度或文档化的 NaN/infinity 策略。它不能替代索引键或普通财务约束中的 `numeric`。

### 已知实现风险

核验的输入函数调用 `strtod`/`strtold`，但没有确认整个字符串都被消费。带尾随垃圾的输入，甚至可能包括空输入，会把数值前缀当作合法值，而不是拒绝。转换前必须校验输入。

`decimal64` 的乘法溢出检查会用结果除以左操作数。当该操作数为零时，有效乘法也可能因为零除零检查而失败。算术检查既不完整也不一致，不能依赖它保证正确性或防止溢出。

`decimal32` 舍入实现允许负 scale，而 `decimal64` 版本拒绝；两者还会经过二进制浮点 `pow`/`round`。应测试所有需要的精度与边界情况。

版本 `1.0` 只有一个 2013 年源码修订，并依赖不可移植的编译器十进制扩展及其 ABI。维护中的精确运算应使用 PostgreSQL `numeric`；该扩展只适合在针对确切编译器/架构/PostgreSQL 组合完成构建测试后用于隔离研究。
