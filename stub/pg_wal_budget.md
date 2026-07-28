## Usage

Sources:

- [Official upstream README](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/README.md)
- [Official extension control file (pg_wal_budget.control)](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/pg_wal_budget.control)
- [Official extension SQL (pg_wal_budget--0.2.1--0.3.0.sql)](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/sql/pg_wal_budget--0.2.1--0.3.0.sql)

`pg_wal_budget` — A Rust/pgrx PostgreSQL 17 extension that observes, predicts, and optionally enforces WAL-generation budgets by policy scope. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_wal_budget;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pwb.clear_tenant()` is an extension function and returns `void`.
- `pwb.counters()` is an extension function and returns `table`.
- `pwb.create_policy(scope_kind text, scope_value text, wal_rate_bytes_per_sec bigint, wal_burst_bytes bigint, mode text default 'observe', priority integer default 100)` is an extension function and returns `integer`.
- `pwb.disable_policy(policy_id integer)` is an extension function and returns `void`.
- `pwb.flush_profiles()` is an extension function and returns `void`.
- `pwb.policies()` is an extension function and returns `setof`.
- `pwb.preload_status()` is an extension function and returns `text`.
- `pwb.query_profiles()` is an extension function and returns `table`.
- `pwb.recent_decisions(decision_limit integer default 100)` is an extension function and returns `table`.
- `pwb.reset_profiles()` is an extension function and returns `void`.
- `pwb.reset_stats()` is an extension function and returns `void`.
- `pwb.scope_names()` is an extension function and returns `table`.
- `pwb.scope_stats()` is an extension function and returns `table`.
- `pwb.set_policy_mode(policy_id integer, mode text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
