## 用法

来源：

- [官方 PostgreSQL-IBAN README](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/README.md)
- [1.1 版扩展 SQL](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban--1.1.sql)
- [IBAN 类型与校验器实现](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban.cpp)

`iban` 增加 `iban` 数据类型和 `iban_validate(text)`，用于国际银行账号的语法校验。输入函数会检查各国 BBAN 结构和 ISO 7064 mod-97 校验和，因此适合在数据库边界拒绝格式错误的值。

### 存储已校验的值

```sql
CREATE EXTENSION iban;

CREATE TABLE beneficiaries (
    beneficiary_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_iban   iban NOT NULL
);

INSERT INTO beneficiaries (account_iban)
VALUES ('GB82WEST12345698765432');

SELECT account_iban, iban_validate(account_iban::text)
FROM beneficiaries;
```

无效文本在转换或赋值给 `iban` 时会被拒绝。若希望校验失败时只返回布尔值而不触发输入错误，可先调用辅助函数：

```sql
SELECT iban_validate('GB82WEST12345698765432'); -- true
SELECT iban_validate('GB00INVALID');             -- false
```

扩展提供从 `iban` 到 `text` 和 `varchar` 的隐式转换，以及到 `bpchar` 的赋值转换。

### 校验边界

校验过程会把临时副本转换为大写，但输入函数保存原始表示。因此小写字母可能通过校验，输出时仍保持小写。空格和视觉分隔符不会被规范化；如果要求单一显示格式，应由应用在写入前完成规范化。

国家与 BBAN 规则被编译进此版本，可能落后于当前 IBAN 注册表。校验成功只表示编码结构和校验和有效，并不证明账号存在、仍然开放、属于指定主体或能被支付网络访问。二进制接收函数不会重复文本输入校验，因此只有可信的 PostgreSQL 二进制协议编码器才应发送此自定义类型。把它作为长期域约束前，应测试实际构建与应用需要的注册表覆盖范围。
