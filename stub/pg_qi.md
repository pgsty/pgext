## Usage

Sources:

- [Official source](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/README.md)
- [Official source](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/pg_qi.control)
- [Official source](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/LICENSE)
- [Official source](https://github.com/baremonconsulting/pg_qi/blob/cf34297f79f4a9c0463153af77e685d0407df724/sql/pg_qi--0.0.1.sql)

pg_qi 0.0.1 is a pre-alpha shadow classifier for pg_duckdb integration. It always returns pass_through intent and preserves the downstream plan. Automatic routing, telemetry export and policy management are not implemented.

### Enable diagnostics

Preload pg_duckdb before pg_qi and restart. Install pg_qi as a superuser. PostgreSQL 16–18 are tested; PostgreSQL 19 testing is pinned to RC2 and does not establish final-release support. The integration requires a real resettable duckdb.force_execution setting and the documented dependency revision.

```ini
shared_preload_libraries = 'pg_duckdb,pg_qi'
```

```sql
CREATE EXTENSION pg_qi;
SELECT qi.version(), qi.is_preloaded();
SHOW pg_qi.enabled;
SHOW pg_qi.shadow_reason;
SHOW pg_qi.shadow_classifier_status;
SHOW pg_qi.shadow_observed_engine;
SHOW pg_qi.shadow_verification;
```

### Interpretation

Use SHOW to inspect the completed backend-local shadow record without generating another plan. unproved_outermost is not proof of a client-originated plan; a negative engine-tree match remains unknown/unverified. pg_qi_spike_test is an explicitly conditional testing companion and is not part of normal installation or a production router.
