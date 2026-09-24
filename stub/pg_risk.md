## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/README.md)
- [Control file](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/pg_risk.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-risk/blob/858651c9c28f5c212ae4a8d5cb5e9138a106f8a0/docs/SECURITY.md)

`pg_risk` evaluates deterministic policies and records auditable decisions on PostgreSQL 14–18. It evaluates permissions or limits around an operation; the application still owns execution of that operation.

### Core Workflow

The following psql example creates a subject, policy, and rolling-volume rule, then evaluates a request. Amount limits use the asset’s smallest units.

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

### Decision Semantics

`risk_check` records a decision and applies idempotency semantics, while `risk_evaluate` provides evaluation without the same persisted-decision workflow. Policy configuration and event history determine the result; success does not execute a payment or reserve external funds. For strict concurrent limits, use the documented subject locking and keep the check and application operation in one transaction.

Core use has no mandatory companion extension. Optional adapters connect to the matching RustedBytes financial extensions; enabling them requires those exact APIs and appropriate grants. The ledger adapter must not be pointed at an unrelated extension with the same name.

### Operation

The extension needs no preload and is non-relocatable. Treat policy administration, event ingestion, and decision execution as separate privileges. Its immutable decision records do not make incorrect input data or an incorrectly configured policy trustworthy.
