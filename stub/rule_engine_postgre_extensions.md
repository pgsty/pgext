## Usage

Sources:

- [Version 2.0.0 README](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/README.md)
- [Version 2.0.0 extension SQL](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/rule_engine_postgre_extensions--2.0.0.sql)
- [Engine selection guide](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/docs/ENGINE_SELECTION.md)
- [External data-source guide](https://github.com/KSD-CO/rule-engine-postgres/blob/178bb1f2230c81b35bb03ebd3d9a931346b00204/docs/EXTERNAL_DATASOURCES.md)

rule_engine_postgre_extensions runs business rules expressed in Grule Rule Language against JSON facts inside PostgreSQL. Version 2.0.0 offers RETE and forward-chaining execution, rule repositories and sets, debug-event persistence, built-in transformation functions, and optional outbound integrations. Treat rules as executable application logic and grant the API only to trusted roles.

### Core Workflow

Install `pgcrypto` first because repository credentials introduced in later schema versions use it, then execute a small rule against copied JSON facts:

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

Use `run_rule_engine_rete` explicitly for RETE behavior or `run_rule_engine_fc` for the simpler forward-chaining path. Validate returned JSON before applying it to authoritative tables.

### Important Objects

- `run_rule_engine`, `run_rule_engine_rete`, and `run_rule_engine_fc` execute facts and rule text with different engines.
- `run_rule_engine_debug`, `debug_get_events`, and `debug_list_sessions` expose execution traces.
- Debug control functions enable, disable, persist, delete, or clear trace sessions.
- Repository functions store, version, activate, and execute named rules.
- Rule-set and statistics functions group rules and expose execution metrics.
- Webhook, external-data-source, and NATS functions connect rule processing to remote systems.
- Built-in GRL functions cover date, string, math, and JSON transformations.

### Security and Operational Notes

Rule and fact inputs are limited to 1 MB according to the documented error contract, but complex rules can still consume backend CPU and memory. Debug persistence stores business facts and decisions and must be retained and protected deliberately. Repository credentials depend on `pgcrypto`; restrict the underlying tables and verify the encryption-key procedure before storing secrets. Webhooks, data sources, and NATS create external, nontransactional effects and can leak facts. Reproduce performance claims on the real workload and keep outbound integrations disabled until network, authentication, timeout, retry, and audit behavior are verified.
