## 用法

来源：

- [目录版本对应的官方 README](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/README.md)
- [目录版本对应的扩展 SQL](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/sql/pg_normalize_email--1.0.9.sql)
- [PGXN 发行页面](https://pgxn.org/dist/pg_normalize_email/)

`pg_normalize_email` 1.0.9 为若干指定邮件服务商提供一个 immutable 规范化函数。它把地址转为小写、映射一个服务商域名，并按服务商规则移除点号或加号标签。只有当账户去重策略明确接受把不同输入视为同一身份的风险时才应使用它。

### 核心流程

```sql
CREATE EXTENSION pg_normalize_email;

SELECT normalize_email('First.Last+news@googlemail.com');
-- firstlast@gmail.com

SELECT normalize_email('User+tag@outlook.com');
-- user@outlook.com

CREATE INDEX account_email_normalized_idx
ON account (normalize_email(email));
```

### 函数行为

- `normalize_email(text)` 要求输入恰好包含一个 at 符号；NULL 或无法正确拆分的输入会抛出异常。
- 所有输入都会转为小写，googlemail.com 会改写成 gmail.com。
- 对 gmail.com 和 live.com，函数移除本地部分中的点号以及第一个加号开始的后缀。
- 对 ya.ru、hotmail.com 和 outlook.com，只移除第一个加号开始的后缀；其他域名仅转为小写。

### 注意事项

- 该函数不是完整的 RFC 邮箱校验器：不会清理空白、检查两部分是否为空、规范化国际化域名，也不处理带引号的本地部分。
- 服务商行为可能变化，服务商特定的规范化也可能合并应用认为不同的地址。应保留原始地址，并把规范化唯一性视为明确的产品决策。
- 函数声明为 immutable；以后修改规则时，需要重建依赖的表达式索引并检查已存储的派生值。
