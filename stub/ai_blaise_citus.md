## Usage

Sources:

- [Official upstream README](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/README.md)
- [Official extension control file (ai_blaise_citus.control)](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/ai_blaise_citus.control)
- [Official implementation source](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/src/lib.rs)

`ai_blaise_citus` — Rust pgrx companion extension for SQL surfaces that coordinate Citus, TimescaleDB, bundled extensions, and sidecars. Use it when an application needs this specific database capability. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION ai_blaise_citus;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add_compression_policy_distributed` is an extension function.
- `add_continuous_aggregate_distributed` is an extension function.
- `add_reorder_policy_distributed` is an extension function.
- `add_retention_policy_distributed` is an extension function.
- `companion_feature_status()` is an extension function.
- `distribute_hypertable` is an extension function.
- `time_range_shard_pruner` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- Upstream explicitly says the project is not production-ready.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
