## 用法

来源：

- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/pg_ledger.control)
- [账本 SQL API 与事务检查](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/src/lib.rs)
- [账户规则实现](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ledger/src/rules.rs)

`pg_ledger` 版本 `0.2.0` 是固定安装在 `pgledger` 模式中的复式记账引擎。它记录不可变的日记账头与明细行，检查事务级借贷平衡，管理币种与会计期间，并可通过附着在应用表上的规则生成分录。

### 核心流程

在一个事务中提交平衡分录，然后查询账户汇总：

```sql
CREATE EXTENSION pg_ledger;

BEGIN;
SELECT pgledger.debit('cash', 100.00, 'sale received');
SELECT pgledger.credit('revenue', 100.00, 'sale received');
SELECT pgledger.check_balance();
COMMIT;

SELECT pgledger.account_balance('cash');
SELECT * FROM pgledger.trial_balance();
SELECT * FROM pgledger.journal_entries();
```

核心类型是 `pgledger.ledgeramount`，支持算术、比较、B-tree 索引以及 `sum`、`min` 和 `max` 聚合。主要分录函数是 `debit`、`credit`、`entry`、`reverse` 和 `check_balance`。管理函数包括 `create_rule`、`enable_rule`、`disable_rule`、`drop_rule`、`set_exchange_rate`、`convert`、`open_period`、`close_period` 和 `period_status`。

### 运维说明

启用 `pg_ledger.enabled` 时，扩展会在提交前拒绝不平衡的账本活动。`pg_ledger.strict_mode` 控制修改没有规则的账本金额列时是否报错；两者都是超级用户可设置的会话 GUC。日记账表由不可变触发器保护，因此更正应使用 `reverse`，而不是 UPDATE 或 DELETE。control 文件不可重定位且仅限超级用户安装。在把日记账当作会计控制前，应验证规则生成的账户、舍入、保存点、期间关闭、备份恢复及应用授权。
