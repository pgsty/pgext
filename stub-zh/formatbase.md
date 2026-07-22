## 用法

来源：

- [固定提交的上游 README](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/README.md)
- [固定提交的 2.0.0 版安装 SQL](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase--2.0.0.sql)
- [固定提交的 C 实现](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.c)
- [固定提交的扩展 control 文件](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.control)

`formatbase` `2.0.0` 可以把有符号整数转换为 2 到 64 进制的文本，也可以把该文本解析回 `smallint`、`integer` 或 `bigint`。它适合紧凑标识符和显式进制转换，但不提供任意精度算术。

### 核心流程

```sql
CREATE EXTENSION formatbase;

SELECT text(255::bigint, 16);       -- ff
SELECT text(-32602::bigint, 36);    -- -p5m
SELECT int4('ff', 16);              -- 255
SELECT int8('-p5m', 36);            -- -32602
SELECT int2(text(7::smallint, 2), 2);
```

2 到 36 进制使用数字与后续小写字母，解析字母时不区分大小写。37 到 64 进制区分大小写，并用大写字母、`_` 和 `` ` `` 扩展字符表。

### 重要对象

- `text(bigint, integer)`、`text(integer, integer)` 和 `text(smallint, integer)`：以指定进制编码有符号整数。
- `int8(text, integer)`、`int4(text, integer)` 和 `int2(text, integer)`：解码文本并返回函数名对应的整数类型。

### 运维说明

所有函数都是 `IMMUTABLE` 和 `STRICT`，因此可以用于生成表达式和索引；任一输入为 `NULL` 时返回 `NULL`。进制超出 2–64、空字符串、非法数字、单独的负号或超出返回类型范围的值都会报错。`text`、`int2`、`int4` 和 `int8` 这些函数名刻意保持简短，外观类似类型转换；search path 可能歧义时应带上 schema。编码文本本身不记录进制，因此必须另行保存所用 base。
