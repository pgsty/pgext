## 用法

来源：

- [Pinned official README](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/README.md)
- [Pinned extension SQL](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/safecast--0.0.1.sql)

`safecast` 提供几个小型 SQL 函数：先用类似数值的正则表达式检查文本，对部分不匹配的值返回 `NULL`，再进行类型转换。尽管名称如此，0.0.1 版并不会捕获转换异常，不能安全处理任意不可信输入。

### 核心流程

```sql
CREATE EXTENSION safecast;

SELECT to_integer('42');
SELECT to_integer('not-a-number');
SELECT to_float('-12.5');
SELECT to_bigint('9000000000');
SELECT to_double_precision('123');
```

第二次调用返回 `NULL`；被接受的字符串会交给 PostgreSQL 的普通类型转换实现。

### 函数与接受形式

- `to_integer(text)`：只接受一个或多个 ASCII 数字，再转换为 `integer`。
- `to_bigint(text)`：只接受一个或多个 ASCII 数字，再转换为 `bigint`。
- `to_float(text)`：接受由 `-`、数字和 `.` 构成的任意非空组合，再转换为 `float`（`double precision`）。
- `to_double_precision(text)`：只接受一个或多个 ASCII 数字，再转换为 `double precision`。

### 注意事项

超出类型范围仍会报错。`to_float` 的正则表达式过于宽松，会接受 `.` 或 `--1` 等畸形值，随后在转换时出错。整数函数拒绝前导符号、空白、小数点和指数形式；`to_double_precision` 同样拒绝符号、小数和指数。也不接受区域相关数值格式以及 `NaN`/`Infinity`。

这些函数没有声明 `STRICT`、易变性或并行属性，而且这个已弃用项目的测试面很小。若必须可靠地把失败变成 `NULL`，应采用显式校验或捕获异常的函数。该扩展是纯 SQL 实现，无需预加载或重启。
