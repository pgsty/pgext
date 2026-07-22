## 用法

来源：

- [扩展控制文件](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils.control)
- [扩展 SQL API](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils--0.0.1.sql)
- [C 实现](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils.c)

`utils` 0.0.1 是 `l-learn` 单仓库中的双函数 C 学习示例。它对两个 PostgreSQL `integer` 值做加法或乘法。尽管名称很宽泛，但并不包含通用工具集合。

### 可用 API

```sql
CREATE EXTENSION utils;

SELECT addme(20, 22);
SELECT multiply(6, 7);
```

`addme(integer, integer) RETURNS integer` 执行 C `int32` 加法。`multiply(integer, integer) RETURNS integer` 执行 C `int32` 乘法。两个函数都是 `STRICT`，所以任一参数为 NULL 时不会调用 C 代码，而是返回 NULL。

### 限制

SQL 声明没有把函数标记为 `IMMUTABLE` 或 `PARALLEL SAFE`，尽管这些算术操作本意上没有数据库副作用。更重要的是，实现直接使用有符号 C 算术，没有 PostgreSQL 通常的整数溢出检查；超范围输入可能触发未定义 C 行为，而不是抛出标准 `integer out of range` 错误。

实际工作应使用 PostgreSQL 内置 `+` 和 `*` 操作符。通用扩展名与函数名可能和无关对象冲突，单仓库也没有提供兼容矩阵或扩展专用用户文档。只能在测试精确服务器构建后，把它当作最小 PGXS/C 示例。
