## 用法

来源：

- [Official PGXN documentation](https://pgxn.org/dist/bignum/doc/bignum.html)
- [Pinned extension SQL](https://github.com/beargiles/pg-bignum/blob/79950064ca96595dd0a26c81915ea2d7396e6986/sql/bignum.sql)

`bignum` 增加了一个基于 OpenSSL `BIGNUM` 实现的任意精度整数类型。当数值需要超出 PostgreSQL 内置 `bigint` 的范围且只需整数运算时可使用它；它不能替代 `numeric` 或 `decimal`。

### 核心流程

创建扩展，将整数文字值转换或写成 `bignum`，然后使用常规算术与比较运算符：

```sql
CREATE EXTENSION bignum;

SELECT '170141183460469231731687303715884105727'::bignum + 1;
SELECT 12::bignum * 12::bigint;
SELECT gcd(48::bignum, 18::bignum);
```

扩展提供从 `int4` 和 `int8` 到该类型的赋值转换，并为算术、比较运算符定义了 `bignum`/`int8` 混合形式。

### 重要对象

- 类型：`bignum`。
- 算术：一元 `-`，以及 `+`、`-`、`*`、`/`、`%`。
- 比较：`=`、`<>`、`<`、`<=`、`>=`、`>`。
- 函数：`abs(bignum)`、`gcd(bignum, bignum)` 及混合类型的 `gcd` 重载。
- 索引：默认 `bignum_ops` B-tree 运算符类支持排序与 B-tree 索引。

### 运维说明

该模块是链接 OpenSSL 的 C 扩展，服务端软件包必须提供匹配的共享库。它可迁移，不要求 `shared_preload_libraries`，也不需要重启。

上游实现没有提供与 `numeric`/`decimal` 之间的转换、密码学函数、哈希索引支持，也没有为 `bignum` 实现 `min`、`max`、`sum`、`avg` 聚合。除法是整数除法。由于只定义了文档所述的 `int4`/`int8` 转换路径，应显式测试混合类型表达式。
