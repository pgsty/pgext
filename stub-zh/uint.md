

## 用法

> [uint: PostgreSQL 无符号整数类型](https://github.com/petere/pguint)

`uint` 扩展为 PostgreSQL 添加了无符号整数和小整数类型。

```sql
CREATE EXTENSION uint;
```

### 数据类型

| 类型 | 说明 | 范围 |
|------|-------------|-------|
| `int1` | 有符号 8 位整数 | -128 到 127 |
| `uint1` | 无符号 8 位整数 | 0 到 255 |
| `uint2` | 无符号 16 位整数 | 0 到 65535 |
| `uint4` | 无符号 32 位整数 | 0 到 4294967295 |
| `uint8` | 无符号 64 位整数 | 0 到 18446744073709551615 |

### 用法

```sql
CREATE TABLE foo (
    a uint4,
    b text
);

SELECT * FROM foo WHERE a > 4;
SELECT avg(a) FROM foo;
```

### 运算符和函数

所有类型包含完整的算术运算符（`+`、`-`、`*`、`/`、`%`）、比较运算符（`=`、`<>`、`<`、`>`、`<=`、`>=`）以及各类型组合之间的运算符。同时提供标准聚合函数和索引支持（Btree、Hash）。
