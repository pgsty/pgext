

## 用法

> [country: ISO 3166-1 alpha-2 国家代码类型](https://github.com/adjust/pg-country)

`country` 扩展提供了一个枚举数据类型，用于表示 ISO 3166-1 alpha-2 国家代码，存储开销极低。

```sql
CREATE EXTENSION country;

CREATE TABLE locations (
    id    serial PRIMARY KEY,
    code  country
);

INSERT INTO locations VALUES (1, 'US'), (2, 'GB'), (3, 'FR');

SELECT * FROM locations ORDER BY code;
```

### 数据类型

`country` 类型将值限制为有效的 ISO 3166-1 alpha-2 代码（两字母国家代码，如 `US`、`GB`、`FR`、`DE` 等）。无效代码在输入时会被拒绝。

### 运算符

支持标准比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

### 索引支持

提供 B-tree 和 Hash 索引运算符类，可在 `country` 列上进行高效查询和排序。
