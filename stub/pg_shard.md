## Usage

Sources:

- [Official README and deprecation notice](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/README.md)
- [Extension SQL for 1.2](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/sql/pg_shard.sql)
- [Official changelog](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/CHANGELOG.md)

`pg_shard` 1.2 is the retired Citus Data sharding extension for distributing hash-partitioned tables across PostgreSQL workers. Upstream declares it end-of-life and identifies the open-source Citus extension as its functional superset; use this document for maintaining or migrating an existing deployment, not for starting a new one.

### Cluster Setup

On the coordinator, preload the module and list one worker host and port per line in the data-directory file shown below. The coordinator must connect to workers over TCP without interactive authentication, and the same database must already exist on every worker.

```conf
shared_preload_libraries = 'pg_shard'
```

```text
# pg_worker_list.conf
worker-101  5432
worker-102  5432
```

Restart the coordinator, create the extension, define a prototype table, and create its shards.

```sql
CREATE EXTENSION pg_shard;

CREATE TABLE customer_reviews (
  customer_id text NOT NULL,
  review_rating integer
);

SELECT master_create_distributed_table('customer_reviews', 'customer_id');
SELECT master_create_worker_shards('customer_reviews', 16, 2);
```

### Main Operations and Limits

- `master_create_distributed_table(...)` chooses the hash partition column.
- `master_create_worker_shards(...)` creates the requested shard count and replication factor.
- `master_copy_shard_placement(...)` repairs an inactive placement; repair requires `pg_shard` on all workers.
- `pgs_distribution_metadata.partition` and `pgs_distribution_metadata.shard` expose distribution metadata.

The prototype table on the coordinator stores metadata and schema rather than user rows. Version 1.2 requires the partition column in the `WHERE` clause for `UPDATE` and `DELETE`. Worker authentication, failure handling, shard repair, and migration need explicit operational procedures. Keep the preload/restart requirement in mind, and plan migration to supported Citus before changing an inherited cluster.
