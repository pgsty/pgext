## 用法

来源：

- [Version 2.0.0 README](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/README.md)
- [Version 2.0.0 extension SQL](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/rule_engine_postgre_extensions--2.0.0.sql)
- [Engine selection guide](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/docs/ENGINE_SELECTION.md)
- [External data-source guide](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/docs/EXTERNAL_DATASOURCES.md)

rule_engine_postgre_extensions 在 PostgreSQL 内运行以 Grule Rule Language 表达的业务规则，并处理 JSON facts。2.0.0 版提供 RETE 与前向链执行、规则仓库与规则集、调试事件持久化、内置转换函数，以及可选出站集成。规则应被视为可执行应用逻辑，API 只能授予可信角色。

### 核心流程

先安装 `pgcrypto`，因为较新模式版本引入的仓库凭据依赖它；然后对复制的 JSON facts 执行一条小规则：

```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION rule_engine_postgre_extensions;

SELECT run_rule_engine(
  '{"Order":{"total":150,"discount":0}}',
  'rule "Discount" {
     when Order.total > 100
     then Order.discount = Order.total * 0.10;
   }'
)::jsonb;
```

如需明确使用 RETE 行为，可调用 `run_rule_engine_rete`；更简单的前向链路径使用 `run_rule_engine_fc`。把返回 JSON 应用于权威表之前必须先校验。

### 重要对象

- `run_rule_engine`、`run_rule_engine_rete` 和 `run_rule_engine_fc` 使用不同引擎执行 facts 与规则文本。
- `run_rule_engine_debug`、`debug_get_events` 和 `debug_list_sessions` 暴露执行轨迹。
- 调试控制函数可启用、禁用、持久化、删除或清空轨迹会话。
- 仓库函数存储、版本化、激活并执行命名规则。
- 规则集和统计函数对规则分组并暴露执行指标。
- Webhook、外部数据源和 NATS 函数把规则处理连接到远程系统。
- 内置 GRL 函数覆盖日期、字符串、数学和 JSON 转换。

### 安全与运维说明

根据文档错误契约，规则和 fact 输入限制为 1 MB，但复杂规则仍可能大量消耗后端 CPU 与内存。调试持久化会保存业务事实和决策，必须有意识地设置保留及保护。仓库凭据依赖 `pgcrypto`；应限制底层表访问，并在保存密钥前核实加密密钥流程。Webhook、数据源和 NATS 会产生外部且非事务性的副作用，并可能泄露 facts。性能声明必须在真实负载上复现；在网络、认证、超时、重试和审计行为得到验证前，应关闭出站集成。
