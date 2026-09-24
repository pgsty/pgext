## 用法

来源：

- [README](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/README.md)
- [Control file / 控制文件](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/pg_money.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/docs/SECURITY.md)

`pg_money` 在 PostgreSQL 14–18 中提供精确的带币种金额、算术、舍入与分摊功能，其类型独立于受区域设置影响的内置金额类型。

### 核心工作流

```sql
CREATE EXTENSION pg_money;
CREATE TABLE invoices (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       total money_with_currency NOT NULL);
INSERT INTO invoices(total) VALUES ('USD 19.95'), (money_make(5.25, 'USD'));
SELECT sum(total), avg(total), money_format(sum(total)) FROM invoices;
SELECT money_split('USD 10.00', 3);
SELECT money_exchange('USD 100', 'EUR', 0.85);
```

### 对象与语义

`money_with_currency` 将币种与精确十进制金额一起存储，`money_minor` 处理最小货币单位金额。`money_make`、`money_from_minor` 和 `money_minor_make` 用于构造值。`money_round` 接受明确的舍入模式，`money_split` 在保持总额不变的前提下分配余数。

算术和聚合会拒绝不兼容币种，不会隐式兑换。汇率须显式提供，或由 `money_exchange_at` 从应用管理的表中读取；扩展不获取市场行情。应在应用边界检查舍入和金额范围。

### 运行

0.3.0 控制文件要求超级用户安装，并允许重定位，无需预加载。应将汇率表和修改权限与只读算术区分管理；SQL 输入本身不能证明汇率正确或及时。
