## Usage

Sources:

- [pg_dropbuffers README at the reviewed commit](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/README.md)
- [pg_dropbuffers 0.0.1 install SQL at the reviewed commit](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers--0.0.1.sql)
- [pg_dropbuffers C implementation at the reviewed commit](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.c)

`pg_dropbuffers` 0.0.1 is a testing-only cache eviction tool. `pg_drop_current_db_buffers` removes buffers for the current database without restarting the cluster. Use it only on an isolated disposable cluster with concurrent writers stopped.

```sql
CREATE EXTENSION pg_dropbuffers;

CHECKPOINT;
SELECT pg_drop_current_db_buffers();
```

`pg_drop_system_cache` invokes the following command through sudo, while `pg_drop_caches` calls both cache-dropping functions:

```text
/usr/bin/sudo /sbin/sysctl vm.drop_caches=3
```

### Caveats

- Never use this extension on production. The implementation warns that buffers can still be dirty when removed, so data loss is possible even if a checkpoint was requested first.
- `pg_drop_system_cache` is Linux-specific and requires the PostgreSQL operating-system user to run the exact sysctl command through passwordless sudo. That is a significant host-level privilege.
- Evicting shared and operating-system caches can severely distort unrelated workloads. Run benchmarks on a dedicated host and record the cache-reset method.
- The extension exposes cluster-impacting functions without a built-in environment guard; restrict EXECUTE privileges explicitly.
