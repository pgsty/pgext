## 用法

来源：

- [pgmp 1.0.6 README](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/README.rst)
- [pgmp 1.0.6 发行说明](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/NEWS.rst)
- [pgmp 1.0.6 元数据](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/META.json)
- [pgmp 控制文件](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/pgmp.control)
- [pgmp 官方文档](https://dvarrazzo.github.io/pgmp/)

`pgmp` 在 PostgreSQL 内提供 GNU MP 算术。它通过 `mpz` 增加任意大小整数，通过 `mpq` 增加精确有理数，并提供类型转换、算术、比较、聚合、数论、位运算与随机数函数。

### 核心流程

```sql
CREATE EXTENSION pgmp;

SELECT '123456789012345678901234567890'::mpz * 2;
SELECT mpq(1::mpz, 3::mpz) + mpq(1::mpz, 6::mpz);
SELECT gcd(48::mpz, 18::mpz);
SELECT nextprime(100000000000000000000::mpz);
```

`mpz` 是任意大小整数类型，但仍受 PostgreSQL 单值大小限制。`mpq` 保存规范化的分子和分母，因此在显式转换为近似类型之前，分数运算保持精确。

### 重要对象

- `mpz(text)` 与类型转换可从十进制或受支持的带基数前缀形式构造整数。
- `mpq(text)` 与 `mpq(mpz, mpz)` 用于构造有理数。
- 两种类型均支持普通比较以及 btree 或 hash 索引。
- 整数辅助函数包括指定舍入方向的除法、幂、开方、素性测试、`gcd`、`lcm`、阶乘、Fibonacci 与 Lucas 数、位运算和随机状态函数。
- 有理数辅助函数包括分子分母访问、求逆、分母限制、算术、比较与聚合。
- `gmp_version()` 与 `gmp_max_bitcnt()` 提供库信息。

需要精确十进制或有理数语义时，不要使用浮点输入；应通过文本、整数，或显式分子和分母构造数值。

### 1.0.6 版本说明

1.0.6 发行版增加 PostgreSQL 19 构建兼容性，在元数据中把 PostgreSQL 14 设为支持下限，并为幂、Fibonacci 与 Lucas 数路径补充缺失的 unsigned-long 范围检查。

上游发行版版本为 1.0.6，但该标签下的 `pgmp.control` 当前声明 SQL 扩展版本为 `1.1`。创建扩展时不要强制版本，并在设计升级前检查数据库实际报告的值：

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmp';
```

pgmp 依赖 GMP 共享库。GMP 4.1 缺少上游文档列出的少数开方、位运算和随机状态函数；需要这些对象时应使用较新的 GMP。大操作数可能消耗大量后端内存和 CPU，因此面向不可信算术负载时应设置语句超时与输入限制。
