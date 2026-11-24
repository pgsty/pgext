## Usage

Sources:

- [Upstream source 1](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/v1.4.0/README.md)
- [Upstream source 2](https://github.com/secp256k1-sha256/adaptive_autovacuum/blob/v1.4.0/adaptive_autovacuum.control)

`adaptive_autovacuum` 1.4.0 is a cluster-wide controller for autovacuum settings, statistics repair and emergency maintenance.

### Core Workflow

First preload the library, restart and install once in `adaptive_autovacuum.control_database` (default postgres). Health functions are available to `pg_monitor`. Settings under `adaptive_autovacuum.policy` are table columns, not GUCs. Table recommendations include apply and revert SQL; this version does not execute those table-setting changes.

```sql
CREATE EXTENSION adaptive_autovacuum;
SELECT * FROM adaptive_autovacuum.doctor();
SELECT * FROM adaptive_autovacuum.status();
SELECT * FROM adaptive_autovacuum.table_recommendations;
```

### Operational Boundaries

Supports PostgreSQL 17/18/19; worker-count changes are advice-only on 17. Fresh policy enables active global tuning, so for observation disable the cluster switch before preload, set dry_run in the policy, then re-enable it. Reserve background-worker slots and coordinate with Patroni/operators. Pausing or dropping does not undo applied global changes or cancel all running work; restore intended settings before removal.

### Version 1.4.0 Maintenance Changes

Emergency takeover now terminates a provably stalled anti-wraparound autovacuum worker and queues replacement maintenance without waiting. It does not target human-started vacuums or vacuums making progress; PostgreSQL 19 failsafe-mode vacuums are excluded. Review this intervention policy before enabling it.

Workers stay idle during `pg_upgrade` binary-upgrade mode. Install the matching new-major library before upgrade, then run `ALTER EXTENSION adaptive_autovacuum UPDATE` in the control database. An upgrade from 1.1.0 requires recreation and loses the old policy/history; follow the version-specific procedure.
