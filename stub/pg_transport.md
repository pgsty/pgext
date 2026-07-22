## Usage

Sources:

- [Amazon RDS transport overview and limitations](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.html)
- [Amazon RDS setup procedure](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Setup.html)
- [Transport procedure](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Transporting.html)
- [Function reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.transport.import_from_server.html)
- [Parameter reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Parameters.html)

`pg_transport` is an Amazon RDS for PostgreSQL extension that physically pulls one database from a source RDS instance into a destination RDS instance. It is designed for large database moves with less processing and downtime than dump/restore, and is not available for self-managed PostgreSQL, Amazon EC2, or non-RDS targets.

### Core Workflow

Source and destination must run the same PostgreSQL version. Use custom parameter groups on both instances, preload the extension, size worker capacity on the destination, reboot, and install `pg_transport` on both sides.

```conf
shared_preload_libraries = 'pg_transport'
pg_transport.num_workers = 3
max_worker_processes = 18
```

```sql
CREATE EXTENSION pg_transport;

SELECT transport.import_from_server(
  'source.example.region.rds.amazonaws.com',
  5432,
  'transport_source',
  'source-temporary-password',
  'appdb',
  'destination-temporary-password',
  true
);
```

Run `transport.import_from_server` on the destination with `dry_run` set to `true` first. After compatibility, space, security-group, credential, and worker checks pass, repeat with `dry_run` set to `false`. Both the source and destination caller must have `rds_superuser`; use temporary roles or rotate both supplied passwords immediately afterward.

`pg_transport.num_workers` accepts 1–32 and defaults to 3. AWS defines the destination worker requirement as `(3 * pg_transport.num_workers) + 9`. `pg_transport.work_mem` limits memory per worker, while `pg_transport.timing` controls progress timing messages.

### Migration Boundaries

Starting transport terminates sessions on the source database and places that database in a special read-only state. The destination database is inaccessible while in transit. Transport is not transactional and is not recorded through normal WAL for point-in-time recovery; destination PITR is unavailable during the operation until the post-transport backup completes. A failed cleanup can leave the source read-only and require explicit recovery.

Only `pg_transport` may be installed in the source database during transport. Source objects must use `pg_default`, `reg` types are unsupported, the destination must not already contain the database name, and read replicas or parents of read replicas cannot participate. Roles, ownership, and ACLs are not transported; destination objects become owned by the local destination user. Inventory and recreate security separately, validate backups and application cutover, and follow current AWS documentation because availability depends on the managed RDS engine release.
