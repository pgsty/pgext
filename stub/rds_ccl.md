## Usage

Sources:

- [Alibaba Cloud rds_ccl SQL-throttling documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-ccl-extension-to-perform-sql-throttling-on-an-apsaradb-rds-for-postgresql-instance)
- [Alibaba Cloud RDS product overview](https://www.alibabacloud.com/help/en/rds/product-overview/what-is-apsaradb-rds-what-is-apsaradb-rds)

`rds_ccl` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL service extension that limits concurrent executions of selected normalized statements and queues excess work. It is provider-only and is not a portable extension for self-managed PostgreSQL.

### Core Workflow

Enable the instance parameters `rds_enable_ccl=on` and `compute_query_id=auto` or `on`, install the extension through RDS extension management, then create rules with a privileged RDS account. This example limits one statement to three active executions and two waiters on the primary:

```sql
CREATE EXTENSION rds_ccl;

SELECT rds_add_ccl_rule(
    $$SELECT * FROM ccl_tbl;$$,
    1,
    3,
    2,
    true,
    'limit constant select',
    ''
);

SELECT * FROM rds_show_current_db_ccl_rule();
```

Rules match normalized `query_id` values, so statements that differ only in literals can share a rule. Use `rds_get_query_id(text)` to calculate an identifier, or `rds_add_ccl_rule_with_query_id(...)` to create a rule from a known identifier.

### Rule Operations

- `node_tag` values 1, 2, and 3 allow loading on the primary, read-only instances, or both; a read-only instance still requires an explicit `rds_load_ccl_rule(integer)` call.
- `rds_enabled_ccl_rule` shows enabled rules across databases, while `rds_show_current_db_ccl_rule()` shows full current-database details.
- `rds_enable_ccl_rule(integer)`, `rds_disable_ccl_rule(integer)`, and `rds_disable_all()` control persisted rule state.
- `rds_update_ccl_rule(integer, integer, integer)` changes only concurrency and queue limits. `rds_unload_ccl_rule(integer, varchar)` removes a rule from memory, while `rds_del_ccl_rule(integer)` deletes it.

### Service Boundaries

Rules are created on the primary. Enabled primary rules load automatically when created or enabled, but in-memory rules do not survive instance restart or primary/secondary failover; reload them afterward. A rule for a read-only instance must have a compatible `node_tag` and be loaded on that instance.

Low concurrency can stall legitimate traffic, and exceeding `max_waiting` aborts the new transaction. Derive identifiers in the same database and `search_path` context as production, stage limits gradually, monitor active and queued statements, and keep tested disable and unload procedures. Availability, version, privileges, and behavior remain governed by the current Alibaba Cloud RDS engine and service policy.
