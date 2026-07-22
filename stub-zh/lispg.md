## 用法

来源：

- [扩展 control 文件](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg.control)
- [1.0 扩展 SQL](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg--1.0.sql)
- [求值器实现](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg_operations.c)

`lispg` 1.0 是 PostgreSQL 的实验性类 Lisp 值类型与求值器。已实现的求值器只是一个非常小的数值计算器，不能把它当作通用 Lisp 运行时或过程语言。

### 核心流程

```sql
CREATE EXTENSION lispg;

SELECT '(+ 2 3)'::lispg;
SELECT lispg_debug('(+ 2 3)');
```

自定义类型负责解析和序列化带括号的表达式。调试求值函数解析 C 字符串、执行表达式，并返回其文本表示。

### SQL 对象

- `lispg`：采用 extended 存储的变长类型。
- `lispg_in(cstring)` 与 `lispg_out(lispg)`：该类型的输入和输出函数。
- `lispg_debug(cstring) RETURNS text`：执行表达式以供检查。

固定提交中的实现只注册了二元数值运算 `+` 和 `-`，且没有提供持久化、沙箱、SQL 访问或稳定语言 API 的文档。它只适合针对实际安装的精确构建进行实验。
