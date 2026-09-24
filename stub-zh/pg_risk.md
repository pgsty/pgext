## 用法

来源：

- [README](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/README.md)
- [Control file / 控制文件](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/pg_risk.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/docs/SECURITY.md)

`pg_risk` 在 PostgreSQL 14–18 中评估确定性规则并记录可审计的决策。它检查操作的许可或限额，实际操作仍由应用执行。

### 核心工作流

以下 psql 示例创建主体、策略和滚动金额限额规则，再评估请求。金额限额使用资产的最小单位。

```sql
CREATE EXTENSION pg_risk;
SELECT risk_subject_create('CUSTOMER', 'customer-123', 'retail') AS customer_id \gset
SELECT risk_policy_create(name => 'retail-withdrawal', operation => 'WITHDRAWAL',
  segment => 'retail') AS policy_id \gset
SELECT risk_rule_create(policy_id => :'policy_id', name => 'daily-usd-50k',
  rule_kind => 'ROLLING_VOLUME_LIMIT',
  config => '{"asset":"USD","window":"24 hours","max_units":"5000000"}');
SELECT * FROM risk_check(subject_id => :'customer_id', operation => 'WITHDRAWAL',
  input_amount => 'USD 1000', idempotency_key => 'withdrawal:request-456');
```

### 决策语义

`risk_check` 记录决策并应用幂等语义，`risk_evaluate` 提供评估而不走相同的持久决策流程。结果由策略配置和事件历史决定；评估通过不代表已经付款或预留外部资金。严格并发限额需要使用上游的主体锁定流程，并将检查与应用操作置于同一事务。

核心用法不强制依赖配套扩展。可选适配器连接匹配的 RustedBytes 金融扩展，需要对应的准确 API 和授权。账本适配器不能指向名称相同但不相关的扩展。

### 运行

扩展不需要预加载，且不可重定位。应分别管理规则管理、事件输入和决策执行权限。决策记录不可变，并不能让错误输入或错误规则变得可信。
