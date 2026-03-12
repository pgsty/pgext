

## 用法

> [numeral: 文本数词数据类型（英语、德语、罗马数字）](https://github.com/df7cb/postgresql-numeral)

`numeral` 扩展提供了三种自定义数值数据类型，使用文本数词而非数字。

```sql
CREATE EXTENSION numeral;
```

### 数据类型

- **`numeral`**：英语数词，使用短级差制（10^9 = billion）
- **`zahl`**：德语数词，使用长级差制（10^9 = Milliarde）
- **`roman`**：罗马数字

三种类型在内部与 `bigint` 二进制兼容，并可隐式转换为 `bigint`。

### 示例

```sql
-- 英语数词
SELECT 'thirty'::numeral + 'twelve'::numeral;
-- forty-two

-- 德语数词
SELECT 'siebzehn'::zahl * 'dreiundzwanzig'::zahl;
-- dreihunderteinundneunzig

-- 罗马数字
SELECT 'MCMLV'::roman + 'II'::roman * 'XXX'::roman;
-- MMXV
```

### 运算符

标准算术运算符（`+`、`-`、`*`、`/`）通过隐式 `bigint` 转换实现。所有现有的 bigint 运算符和函数均可使用。
