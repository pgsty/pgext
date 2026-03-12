

## 用法

> [currency: ISO 4217 货币代码类型](https://github.com/adjust/pg-currency)

`currency` 扩展提供了 ISO 4217 货币代码数据类型，每个值仅占用一个字节的存储空间。

```sql
CREATE EXTENSION currency;

CREATE TABLE transactions (
    id                serial,
    payment_currency  currency
);

INSERT INTO transactions VALUES (1, 'USD'), (2, 'EUR'), (3, 'USD');

SELECT * FROM transactions ORDER BY payment_currency;
 id | payment_currency
----+------------------
  2 | EUR
  1 | USD
  3 | USD
```

### 运算符

支持标准比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

内置 B-tree 索引支持，可进行高效排序和查找。

### 函数

```sql
-- 列出所有支持的货币代码
SELECT * FROM supported_currencies() currency ORDER BY currency;
```
