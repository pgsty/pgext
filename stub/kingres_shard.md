## Usage

Sources:

- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/README.md)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/kingres_shard.control)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/kingres_shard--0.5.sql)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_shard` 0.5 adds hash-bucket distribution, reference tables, statement routing and two-phase commit to KingresDB. It requires the KingresDB kernel marker and does not build against stock PostgreSQL. The source supports the PostgreSQL 18 and 19 branches; the latter is based on a release candidate.

### Configure a cluster

Preload on every node and restart. All nodes must use one PostgreSQL major and compatible extension files, with libpq connectivity and matching roles/databases. The prepared-transaction limit must cover concurrent distributed work. Administrative APIs require superuser access.

```ini
shared_preload_libraries = 'kingres_shard'
max_prepared_transactions = 100
```

```sql
CREATE EXTENSION kingres_shard;
SELECT kingres.init_cluster('host1', 5432);
SELECT kingres.add_node('host2', 5432);
CREATE TABLE public.users (id int PRIMARY KEY, name text);
SELECT kingres.create_distributed_table('public.users', 'id');
SELECT * FROM kingres.check_cluster();
```

### Routing and maintenance

`kingres.create_reference_table` replicates a small table across nodes; `kingres.move_bucket` and `kingres.rebalance_plan` support placement changes. `kingres.recover_prepared_transactions` resolves recorded decisions. DDL is coordinated through the map-group primary. The shared `kingres` schema can coexist with the other family extensions. Upgrade libraries on all nodes, restart, and apply the supplied SQL chain before relying on new helpers.

### Correctness boundary

Two-phase commit is not a global snapshot: distributed reads can observe different snapshots, including between scans of one statement. Unsupported multi-row writes are rejected. Rolling upgrades between PostgreSQL majors are unsupported. The reviewed revision reports unresolved high-severity failover defects, and automatic control-plane failover is unfinished. Do not rely on failover durability or treat this preview as a production-ready distributed database.
