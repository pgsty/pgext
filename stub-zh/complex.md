## 用法

来源：

- [1.0 版本 SQL API](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex--1.0.sql)
- [C 实现](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.c)
- [扩展控制文件](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.control)

`complex` 定义了由两个 `float8` 分量组成的 16 字节复数类型。1.0 版本提供算术和比较操作符、从 `int4` 与 `float8` 的隐式转换、`sum`、`min`、`max`、`avg` 以及 `complex_pwrt(complex,int4)`。

```sql
CREATE EXTENSION complex;
SELECT '(1,2)'::complex + '(3,-4)'::complex;
SELECT avg(z) FROM measurements;
```

SQL API 为 `<`、`<=`、`>`、`>=`、`min` 和 `max` 定义了全序，但复数在数学上没有自然顺序。在用于业务逻辑或排序索引前，必须检查并测试实现选择的排序。浮点相等还继承 NaN、无穷、带符号零、舍入及平台问题。

仓库没有用户 README 或发行说明。应在精确构建上确定可接受文本语法、幂函数的多行返回语义、除零行为、溢出、二进制转储恢复及 PostgreSQL 大版本兼容性。当可移植性和明确数值契约重要时，应使用核心数值对或应用数学库。
