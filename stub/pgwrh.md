## Usage

Sources:

- [Release README](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/README.md)
- [Cluster and rollout guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/overview.md)
- [Core control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh/pgwrh.control)
- [Core SQL API](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh/src/master/api-management.sql)
- [Alpha release boundaries](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/releases/1.0.0-alpha1.md)

`pgwrh` 1.0.0-alpha1 scales PostgreSQL 18 reads by assigning partition leaves to logical replicas. The controller retains the source data and accepts writes; replicas query their local shards and reach other shards through foreign tables. This testing alpha supports fresh installations and does not promise an in-place upgrade to a later release.

### Enable and Configure

Install as an administrator on the controller and each replica. The core requires `pg_background` 1.6 or newer, the bundled `pgwrh_fdw`, and PL/pgSQL. Stock `postgres_fdw` is not a dependency. The SQL core has no preload requirement; the optional replication-wait component has its own startup requirement.

```sql
CREATE EXTENSION pgwrh CASCADE;
```

On the controller, create a replication group and register an already provisioned replica. The example assumes an existing non-superuser login with replication privileges, source-table grants, and working network authentication. Configure the source partition hierarchy and its policy in `pgwrh.sharded_table` before rolling out placement.

```sql
SELECT pgwrh.create_replica_cluster('readers');
SELECT pgwrh.add_replica('readers', 'replica_a', 'replica-a', 5432,
                        _dbname := 'read_a');
```

On each replica, `pgwrh.configure_controller` supplies the controller endpoint, replication login, authentication secret, and controller database. The controller and replica database names may differ. Managed replica-to-replica connections require SCRAM authentication; credential rotation is separate from topology changes.

### Review and Apply Placement

A replication factor is a percentage of eligible replicas, with an optional minimum copy count. Placement can also account for availability zones. The following call creates or reuses an editable pending configuration and previews its shard assignments:

```sql
SELECT * FROM pgwrh.preview_shard_placement(
    'readers', pgwrh.next_pending_version('readers')
);
SELECT pgwrh.start_rollout('readers');
```

Inspect `pgwrh.missing_subscribed_shard`, `pgwrh.missing_connected_local_shard`, and `pgwrh.missing_ready_remote_shard` for readiness blockers. Call `pgwrh.commit_rollout` only after preparation succeeds; it checks readiness again. Use `pgwrh.rollback_rollout` to abandon a rollout and wait for replica acknowledgements before another change. Existing copies are retained while replacement routes are prepared.

### Consistency and Operations

Logical replication is asynchronous. Use `pgwrh_wait` when reads must observe a known committed write, and `pgwrh_ui` for the optional controller console. Placement readiness does not certify current reachability or zero replication lag. Queries spanning replicas have no single cluster-wide snapshot.

Coordinate schema changes manually. This release does not elect a replacement controller or make schema changes atomic across nodes. Maintain controller backups and a PostgreSQL availability plan; shard redundancy does not replace either. Marking a host offline changes routing, not its stored data or assigned replication.
