## 用法

来源：

- [官方 README](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/README.md)
- [SQL 1.2](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban--1.2.sql)
- [1.1 升级脚本](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban--1.1--1.2.sql)
- [类型与校验器实现](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban.cpp)

`iban` 的 SQL 版本为 1.2，上游分发版本为 1.2.0。它增加经过校验的国际银行账号类型及 `iban_validate(text)`，通过各国 BBAN 结构与 mod-97 校验和拒绝格式错误的值。校验成功不代表账号存在、属于指定主体或能够通过支付网络访问。上游支持 PostgreSQL 14–18。

### 存储与比较

```sql
CREATE EXTENSION iban;
CREATE TABLE beneficiaries (
    beneficiary_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_iban iban NOT NULL
);
INSERT INTO beneficiaries (account_iban)
VALUES ('gb82 west 1234 5698 7654 32');
SELECT account_iban, iban_validate(account_iban::text)
FROM beneficiaries ORDER BY account_iban;
CREATE INDEX beneficiaries_iban_idx ON beneficiaries (account_iban);
```

文本输入转换为大写并移除普通空格，因此示例保存为 GB82WEST12345698765432。制表符、换行和其他空白不会移除。超过 64 字节的输入被拒绝，错误消息最多回显输入的 40 字节。二进制输入现在也按相同规则校验和规范化。

### 主要对象

- `iban_validate(text)` 返回校验结果而不保存值。
- `iban` 支持到文本类型的转换与比较运算符，`iban_cmp(iban, iban)` 按字节顺序比较。
- 默认 btree 与 hash 运算符类支持排序、相等比较、连接、分组和键；`iban_hash(iban)` 与 `iban_hash_extended(iban, bigint)` 提供哈希支持。

### 升级与校验边界

安装匹配文件，替换仍加载旧库的会话，再升级 SQL，并规范化旧版本写入的行：

```sql
ALTER EXTENSION iban UPDATE TO '1.2';
UPDATE beneficiaries SET account_iban = account_iban::text::iban;
```

在重写生产数据或创建唯一索引前，应检查已有值及规范化后的唯一性冲突：旧行可能仍保存小写或非规范表示。国家规则编译在扩展中，可能落后于注册表变化，应按本次确切版本检查应用需要的国家覆盖范围。
