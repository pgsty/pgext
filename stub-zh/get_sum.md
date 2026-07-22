## 用法

来源：

- [扩展控制文件](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.control)
- [0.0.1 版扩展 SQL](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum--0.0.1.sql)
- [C 实现](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.c)
- [上游回归测试预期](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/expected/get_sum_test.out)

`get_sum` 0.0.1 是一个最小化 C 扩展示例，用于相加两个 PostgreSQL `integer` 值。它是教学样例，并不是内置 `+` 运算符更安全或功能更强的替代品。

### 核心流程

```sql
CREATE EXTENSION get_sum;

SELECT get_sum(101, 202);
SELECT get_sum(NULL, 11);
```

第一次调用返回 `303`。SQL 声明带有 `STRICT`，因此任一参数为 NULL 时会直接返回 NULL，不会进入 C 函数。

### 接口

`get_sum(integer, integer) RETURNS integer` 声明为 `IMMUTABLE STRICT`。它只接受两个 `int4` 值并返回其 C `int32` 和值，没有 `bigint`、`numeric`、聚合、可变参数或溢出检查重载。

### 溢出注意事项

实现直接执行有符号 32 位 C 加法，没有调用 PostgreSQL 的整数溢出检查。上游回归输出记录了 `get_sum(2147483647, 2147483647)` 返回 `-2`，而不是抛出 `integer out of range`。C 有符号溢出并不是稳定的应用契约，其行为可能受编译器假设影响。

普通查询应使用 PostgreSQL 内置算术。即使开放此样例，也应仅用于已证明输入和结果都处于 `integer` 范围内的演示。
