## Usage

Sources:

- [Pinned current upstream README](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/README.md)
- [Version 0.1.0 installation SQL](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/pg_recyclebin--0.1.0.sql)
- [Pinned Cargo metadata](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/Cargo.toml)
- [Pinned backup and restore test matrix](https://github.com/CaghanTU/pg_recyclebin/blob/0b78576e548863c598acd0f46545e146592f2132/docs/BACKUP_RESTORE_TEST_MATRIX.md)

pg_recyclebin version 0.1.0 intercepts selected DROP and TRUNCATE operations. Dropped tables are moved and renamed under flashback_recycle; truncated data is copied there. Metadata under flashback records the original owner, schema, operation, dependencies, and retention deadline so an entry can be listed, restored, or permanently purged.

### Preload and recovery example

```conf
shared_preload_libraries = 'pg_recyclebin'
```

```sql
CREATE EXTENSION pg_recyclebin;

CREATE TABLE recycle_demo (id integer);
INSERT INTO recycle_demo VALUES (1), (2);

DROP TABLE recycle_demo;

SELECT * FROM flashback_list_recycled_tables();
SELECT flashback_restore('recycle_demo');

SELECT * FROM recycle_demo ORDER BY id;
```

The current project documents PostgreSQL 14 through 18 and pgrx 0.16.1. A background worker purges expired entries. Its default connection database is postgres; installations in another database should set flashback.database_name. Runtime limits control retention days, table count, total size, worker interval, and excluded schemas.

### Recovery is not backup

The hook deliberately changes normal DROP and TRUNCATE semantics, including multi-table commands, CASCADE, and DROP SCHEMA CASCADE. It may move partitions, recreate dependent views and incoming foreign keys, preserve policies and triggers, or copy a complete truncated table. These operations add DDL latency, storage and WAL pressure, lock interactions, and new failure modes. Test transactional rollback, dependencies, partitions, identity sequences, ownership, RLS, and quoted names against the exact release.

The recycle bin cannot protect against DROP DATABASE, cluster or disk loss, temporary or excluded tables, privileged purge, retention eviction, extension failure, or corruption. Keep independent tested backups and point-in-time recovery. Restrict restore and purge functions, monitor recycle storage, and rehearse both ordinary restore and mass recovery before relying on it for operator mistakes.
