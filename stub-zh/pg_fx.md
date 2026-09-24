## 用法

来源：

- [README](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/README.md)
- [Control file / 控制文件](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/pg_fx.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/docs/SECURITY.md)

`pg_fx` 在 PostgreSQL 14–18 中存储外部提供的汇率、计算精确定价并创建有期限的报价。行情由应用采集，扩展不访问行情源，也不转移余额。

### 核心工作流

```sql
CREATE EXTENSION pg_fx;
SELECT fx_source_upsert('primary_bank', source_priority => 10,
                        source_max_age => interval '30 seconds');
SELECT fx_rate_insert(source => 'primary_bank', rate_pair => 'USD/EUR',
  rate_bid => 0.8510, rate_ask => 0.8520,
  rate_observed_at => clock_timestamp(), rate_volume => 100000);
SELECT fx_bid('USD/EUR'), fx_ask('USD/EUR'), fx_mid('USD/EUR');
SELECT fx_vwap('USD/EUR'), fx_weighted_median('USD/EUR');
```

### 定价与报价

`fx_rule_create` 按客户分组和金额范围配置加价与费用，`fx_create_quote` 记录有期限的报价。`fx_execute_quote` 只改变报价状态，不向账户记账。应用账本记账和报价状态迁移应放在同一个事务中。来源优先级、过期阈值、资产精度与舍入都会影响返回价格。

安装相关扩展后，`fx_enable_pg_money` 与 `fx_enable_pg_cryptocurrency` 可添加可选类型适配器。独立的账本集成针对 RustedBytes 项目，不应与另一个使用相同扩展名称的项目混淆。

### 权限与维护

无需预加载。创建时可选择安装模式，但之后不能重定位。核心功能不硬性依赖配套扩展。应按安全文档仅授予必需的来源、汇率、规则和报价操作权限，不应允许非受信调用者发布权威汇率。
