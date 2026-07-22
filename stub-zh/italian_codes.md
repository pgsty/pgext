## 用法

来源：

- [上游英文文档](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/doc/italian_codes.rst)
- [1.0 版本 SQL 实现](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/sql/italian_codes.sql)
- [扩展控制文件](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/italian_codes.control)

`italian_codes` 为意大利税务代码提供文本域和校验函数：`codice_fiscale` 接受自然人的 16 字符代码或法人的 11 位数字代码，`partita_iva` 接受意大利增值税号。它只校验给定标识符，不会根据个人资料推导税号。

```sql
CREATE EXTENSION italian_codes;

SELECT codice_fiscale('mss trs 53b19 h892p'::text);
SELECT partita_iva('0575346 048 3'::text);
```

构造函数会移除空白，`codice_fiscale` 还会把输入转成大写。应显式传入 `text` 表达式：PostgreSQL 可能把 `codice_fiscale('X')` 这样的裸字面量解释为域字面量，从而绕过规范化函数。

希望把校验错误作为文本返回而不是抛出 `check_violation` 时，可使用 `codice_fiscale_error(text)` 或 `partita_iva_error(text)`；这些辅助函数要求输入已经规范化。底层规范化函数本身不校验内容，校验错误消息为意大利语。

扩展有意不提供税号生成，因为同码问题可能需要税务机关签发替代码。应保留权威原始标识符，区分格式校验与身份核验，也不要把这些域当作个人或企业真实存在的证明。
