## Usage

Sources:

- [Official README and operating guide](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/README.md)
- [Version 1.1 control file](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration.control)
- [Version 1.1 SQL implementation](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration--1.1.sql)

`cigration` 1.1 orchestrates online movement of colocated Citus shards for expansion, contraction, rebalancing, and worker replacement. It copies data and uses logical replication to catch up writes, briefly blocks writes while switching metadata, and preserves old shards in recycle-bin schemas until an administrator removes them. It is a privileged cluster-maintenance tool, not an application-level migration framework.

### Prerequisites and Installation

Upstream specifies PostgreSQL 10 or later, Citus 9.2 or later, and `dblink` on the coordinator. All participating nodes must use `wal_level = logical`; the executing database role must be able to connect without an interactive password among the coordinator, workers, standbys, and local server. Review this strong trust boundary before use.

Disable Citus DDL propagation while installing on the coordinator:

```sql
SET citus.enable_ddl_propagation TO off;
CREATE EXTENSION IF NOT EXISTS dblink;
CREATE EXTENSION cigration;
RESET citus.enable_ddl_propagation;
```

After installation, use `cigration.cigration_create_distributed_table` instead of the ordinary two-argument `create_distributed_table`, which the extension replaces with a guard.

### Rebalance Workflow

Create a job, inspect its colocated-shard tasks, and then run it:

```sql
SELECT * FROM cigration.cigration_create_rebalance_job();

SELECT jobid, taskid, status, source_nodename, target_nodename,
       all_colocated_logicalrels, total_shard_size
FROM cigration.pg_citus_shard_migration
ORDER BY jobid, taskid;

SELECT cigration.cigration_run_shard_migration_job();
```

Old source shards are renamed into per-job recycle-bin schemas. Inspect them before destructive cleanup:

```sql
SELECT * FROM cigration.cigration_get_recyclebin_metadata();
SELECT cigration.cigration_cleanup_recyclebin();
```

Other entry points include `cigration_create_drain_node_job`, `cigration_create_move_node_job`, `cigration_create_general_shard_migration_job`, `cigration_cancel_shard_migration_job`, and `cigration_cleanup_error_env`.

### Supported Data and Safety Boundaries

Upstream limits migration to single-replica hash-distributed tables. Reference tables, multi-replica shards, append-distributed tables, and unlogged tables are unsupported. Every table in a colocation group moves as one atomic task, so poor colocation design can create very large moves.

The extension installs event triggers that block dropping task tables and block `ALTER TABLE`, `CREATE INDEX`, `ALTER INDEX`, and `DROP INDEX` while any migration task is running. It also creates and drops publications, subscriptions, slots, shard tables, and Citus metadata through `dblink`; use a dedicated maintenance window and test recovery procedures.

Tables without a primary key or replica identity cannot accept concurrent UPDATE or DELETE during logical transfer. PostgreSQL 10 logical replication does not propagate TRUNCATE, so concurrent TRUNCATE can cause divergence. Long transactions and HA failover can also interrupt jobs. After fixing an error, use `cigration_cleanup_error_env` only when every migration job is stopped; remove recycle-bin data only after validating the new placements and backups.
