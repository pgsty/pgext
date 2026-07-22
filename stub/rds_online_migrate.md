## Usage

Sources:

- [Official Alibaba Cloud online-partitioning guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-online-migrate-extension-to-partition-tables-online)
- [Official ApsaraDB RDS for PostgreSQL product documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_online_migrate` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that converts a regular table to a partitioned table, or repartitions an existing table, while writes continue. It copies existing rows, streams changes through logical replication, and atomically renames the source and destination after synchronization. It is a provider-managed feature, not a portable community extension.

### Prerequisites and Core Workflow

Alibaba Cloud's current guide requires RDS PostgreSQL 13 or later with a qualifying minor engine version, `wal_level = logical`, row identity through a primary key or `REPLICA IDENTITY`, enough privileges for publications/subscriptions/slots and table rename, a verified backup, and free storage of at least twice the source table plus indexes.

Create the destination partitioned table in the same schema, then start the conversion:

```sql
CREATE EXTENSION rds_online_migrate;

CREATE TABLE public.events_new
    (LIKE public.events INCLUDING ALL)
    PARTITION BY HASH (id);
CREATE TABLE public.events_new_0 PARTITION OF public.events_new
    FOR VALUES WITH (MODULUS 2, REMAINDER 0);
CREATE TABLE public.events_new_1 PARTITION OF public.events_new
    FOR VALUES WITH (MODULUS 2, REMAINDER 1);

SELECT rds_online_migrate.rewrite_table(
    'public.events',
    'public.events_new'
);
```

A successful call returns `true`. The destination becomes the original table name, while the source is retained with an `_rds_bkp` suffix for verification.

### Post-migration and Failure Handling

Views, triggers, and foreign-key constraints associated with the original table are not automatically transferred to the replacement and must be rebuilt. If the source was in a logical-replication publication, add the renamed replacement back. Reapply and verify business-role privileges, compare row counts and partitions, and retain the backup until application validation completes.

The procedure is non-atomic. An interruption can leave metadata rows, a publication, subscription, or replication slot that requires manual cleanup following the official guide. Concurrent tasks may require more `max_worker_processes`. Test the complete conversion and recovery path in a non-production RDS instance and use the service's current eligibility and support process to enable the feature.
