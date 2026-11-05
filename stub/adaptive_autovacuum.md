## Usage

Sources:

- [Upstream source 1](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/6712f2bc62cd7c4c7cc2c264fc8977c7f23b111f/README.md)
- [Upstream source 2](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/6712f2bc62cd7c4c7cc2c264fc8977c7f23b111f/adaptive_autovacuum.control)

`adaptive_autovacuum` 1.3.1 is a beta cluster-wide controller for autovacuum settings, statistics repair and emergency maintenance.

### Core Workflow

First preload the library, restart and install once in `adaptive_autovacuum.control_database` (default postgres). Health functions are available to `pg_monitor`. Settings under `adaptive_autovacuum.policy` are table columns, not GUCs. Table recommendations include apply and revert SQL; this version does not execute those table-setting changes.

```sql
CREATE EXTENSION adaptive_autovacuum;
SELECT * FROM adaptive_autovacuum.doctor();
SELECT * FROM adaptive_autovacuum.status();
SELECT * FROM adaptive_autovacuum.table_recommendations;
```

### Operational Boundaries

Supports PostgreSQL 17/18; worker-count changes are advice-only on 17. Fresh policy enables active global tuning, so for observation disable the cluster switch before preload, set dry_run in the policy, then re-enable it. Reserve background-worker slots and coordinate with Patroni/operators. Pausing or dropping does not undo applied global changes or cancel all running work; restore intended settings before removal.
