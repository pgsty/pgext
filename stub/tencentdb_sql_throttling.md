## Usage

Sources:

- [TencentDB PostgreSQL SQL throttling documentation](https://intl.cloud.tencent.com/zh/document/product/409/64750)
- [TencentDB for PostgreSQL product page](https://cloud.tencent.com/product/postgres)

`tencentdb_sql_throttling` is a TencentDB-managed extension that limits concurrent executions by normalized query ID. It is not a portable community package: Tencent Cloud controls availability, supported kernel builds, preload changes, privileges, HA behavior, and upgrades.

```sql
CREATE EXTENSION tencentdb_sql_throttling;
SELECT tencentdb_sql_throttling.add_rule_with_queryid(
  497939862935121343, 0, 10, true
);
SELECT * FROM tencentdb_sql_throttling.rules;
```

Tencent Cloud support must add `tencentdb_sql_throttling` to `shared_preload_libraries`, which restarts the instance. Official documentation requires at least the listed TencentDB kernel revisions and says Extended Query Protocol statements cannot be configured directly. Obtain query IDs from `pg_stat_activity` or `pg_stat_statements`, and use `work_node` 0 for primary plus read-only nodes, 1 for primary only, or 2 for read-only only.

Rules affect traffic only while loaded in memory. A restart, HA event, or major-version upgrade clears memory rules; persist them in `tencentdb_sql_throttling.persistent_rules_table` and explicitly reload them. `drop_rule`, `drop_all_rules`, `change_rule_status`, and the load/dump helpers distinguish memory from persistent state; verify both tables after every control operation.

Start with a conservative `max_concurrency`, monitor `current_concurrency`, `total_hit_count`, `reject_count`, latency, and application errors, and keep an emergency rollback path. Query-ID changes after SQL, schema, planner, or version changes can make a rule miss or affect an unintended normalized query, so rehearse failover and upgrade behavior with the provider.
