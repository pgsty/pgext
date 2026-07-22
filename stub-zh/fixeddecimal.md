## 用法

来源：

- [官方 README](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/README.md)
- [扩展控制文件](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.control)
- [扩展 SQL](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal--1.1.0.sql)
- [类型实现](https://github.com/2ndQuadrant/fixeddecimal/blob/970f56b53a2eef74130527778190597dbd9a6a30/fixeddecimal.c)

`fixeddecimal` 提供一个 8 字节定标小数类型，适合金额等恰好需要两位小数并希望紧凑存储的值。它能在自身范围内精确存储并精确加减，但不提供任意 `numeric` 精度，也不能配置其他小数位数。

### 核心流程

```sql
CREATE EXTENSION fixeddecimal;

CREATE TABLE invoice (
  invoice_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  amount fixeddecimal(12,2) NOT NULL
);

INSERT INTO invoice(amount) VALUES ('19.95'), ('0.05');
SELECT sum(amount), avg(amount) FROM invoice;
CREATE INDEX invoice_amount_idx ON invoice (amount);
```

声明形式为 `fixeddecimal(precision,2)`，精度可从 3 到 17。数值以放大 100 倍的有符号 64 位整数存储，范围为 `-92233720368547758.08` 到 `92233720368547758.07`。扩展提供比较和算术操作符、与常见数值类型间的转换、`sum`、`avg`、B-tree 与哈希操作符类，以及 BRIN minmax 操作符类。

### 精度与兼容性

输入超过两位的小数会被截断而不是四舍五入，乘除法重新定标时也向零截断。业务规则需要显式舍入时，应使用 `round(value::numeric, 2)::fixeddecimal`。溢出会报错，并且该类型没有 `NaN` 表示。

不需要预加载。上游说明支持 PostgreSQL 9.5 及以上版本，但只声明最后测试到 PostgreSQL 12；固定的 1.1.0 源码来自 2019 年。在较新的 PostgreSQL 大版本上采用前，应分别验证编译、二进制升级、转储恢复、类型转换、聚合与索引。以后若需要改变列的小数位数，必须转换为其他类型，因为实现只支持两位小数。
