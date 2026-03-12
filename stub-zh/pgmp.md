

## 用法

> [pgmp: PostgreSQL 多精度算术（GMP）](https://github.com/dvarrazzo/pgmp/)

`pgmp` 扩展将 GNU 多精度算术库（GMP）集成到 PostgreSQL 中，提供任意精度的整数和有理数类型。

```sql
CREATE EXTENSION pgmp;
```

### 数据类型

- **`mpz`**：任意大小的整数，仅受 PostgreSQL 的 1GB varlena 上限限制
- **`mpq`**：任意精度的有理数，用于精确分数运算

### 基本用法

```sql
-- 任意精度整数
SELECT '123456789012345678901234567890'::mpz * 2;

-- 精确有理数运算（无舍入）
SELECT '1'::mpq / '3'::mpq;  -- 1/3

-- 与 PostgreSQL 原生类型混合运算
SELECT 42::mpz + '100'::mpz;
```

### 运算符

`mpz` 和 `mpq` 类型均支持标准算术运算符（`+`、`-`、`*`、`/`、`%`）和比较运算符（`=`、`<>`、`<`、`>`、`<=`、`>=`）。

### 函数

该扩展暴露了 GMP 库中这些类型的所有函数，包括：

- 素数函数
- 随机数生成
- 因式分解
- 最大公约数、最小公倍数及其他数论函数

### 索引支持

`mpz` 和 `mpq` 均支持 Btree 和 Hash 索引，可进行高效查询和排序。

完整文档：https://www.varrazzo.com/pgmp/
