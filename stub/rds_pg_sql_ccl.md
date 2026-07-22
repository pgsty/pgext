## Usage

Sources:

- [Huawei Cloud RDS for PostgreSQL rds_pg_sql_ccl guide](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0073.html)

`rds_pg_sql_ccl` is a Huawei Cloud RDS for PostgreSQL kernel extension for SQL concurrency control. It can cap total concurrent SQL or apply per-query-ID rules with separate running and waiting limits. It is provider-managed and is not an installable community package.

### Availability and Enablement

The provider guide lists PostgreSQL 17 all minor releases, 16.2+, 15.4+, 14.8+, 13.11+, 12.15+, and 11.20+. Check the actual instance, then create the kernel extension through Huawei's control function and enable its kernel parameter:

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_pg_sql_ccl';

SELECT control_extension('create', 'rds_pg_sql_ccl');

SHOW rds_pg_sql_ccl.enable_ccl;
SHOW rds_pg_sql_ccl.max_concurrent_sql;
SHOW rds_pg_sql_ccl.max_enabled_rules;
```

Set `rds_pg_sql_ccl.enable_ccl` to on through the RDS parameter controls. `rds_pg_sql_ccl.max_concurrent_sql` applies the higher-priority instance-wide cap; values at or below zero mean unlimited. `rds_pg_sql_ccl.max_enabled_rules` limits how many rules can be active.

### Per-Query Rules

Create a rule from SQL text, explicitly enable it, and inspect the active set:

```sql
WITH new_rule AS (
  SELECT rds_pg_sql_ccl.add_ccl_rule_by_query(
    'SELECT pg_sleep(1)',
    2,
    10,
    'public'
  ) AS rule_id
)
SELECT rds_pg_sql_ccl.enable_ccl_rule(rule_id)
FROM new_rule;

SELECT * FROM rds_pg_sql_ccl.get_all_enabled_rule;
SELECT * FROM rds_pg_sql_ccl.get_activity_query_status;
SELECT * FROM rds_pg_sql_ccl.get_current_db_ccl_rule;
```

The running limit is `max_concurrency`; `max_waiting` controls how many additional executions may wait. A newly created rule does not take effect until `enable_ccl_rule` is called.

### Important Objects

- `get_query_id(query_string, search_path)` calculates a query ID for SQL without bind variables.
- `add_ccl_rule_by_query(...)` creates a rule from SQL text; `add_ccl_rule_by_queryid(...)` creates one from a known query ID.
- `enable_ccl_rule(bigint)`, `disable_ccl_rule(bigint)`, `disable_all_ccl_rule()`, `delete_ccl_rule(bigint)`, and `update_ccl_rule(...)` manage rule state.
- `get_all_enabled_rule`, `get_activity_query_status`, and `get_current_db_ccl_rule` expose enabled rules, waiting activity, and current-database rule definitions.

### Bind Variables and Lifecycle

`get_query_id` and text-based rule creation cannot identify prepared statements with bind variables. Execute the prepared statement first, obtain its `queryid` from `pg_stat_statements`, then call `add_ccl_rule_by_queryid`. Query IDs and normalized text must be verified against the target engine and search path.

Enabled state is not retained across an instance reboot; rules must be enabled again. Read replicas synchronize primary rules and enable them, but a rule applies only to statements that begin after enablement. Rules are unique by query ID within a database, while different databases can reuse the same query ID.

### Operational Notes

Concurrency limits can reject or queue production queries, so derive values from workload and resource measurements. The instance-wide cap takes precedence over per-query rules. Manage installation, parameters, engine upgrades, and availability through Huawei Cloud RDS, and test failover, reboot re-enablement, replica behavior, queue saturation, and application timeouts before relying on the feature.
