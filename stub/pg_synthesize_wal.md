## Usage

Sources:

- [Official README](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/README.md)
- [Official extension SQL](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/pg_synthesize_wal--1.0.sql)
- [Official extension control file](https://github.com/BRupireddy2/pg_synthesize_wal/blob/5eb6e95f7f3e615626cfc5fede3c789f8b2de852/pg_synthesize_wal.control)

`pg_synthesize_wal` deliberately writes custom WAL records with random or caller-supplied payloads. It is a developer, test, and sandbox tool for exercising WAL behavior—including records large enough to span WAL files—and is explicitly not recommended for production servers.

### Enable and Generate WAL

The library must register its custom resource manager during postmaster startup. The control file also declares `pg_walinspect` as an extension dependency, although the implementation uses it only in tests.

```sql
ALTER SYSTEM SET shared_preload_libraries = 'pg_synthesize_wal';
-- Restart PostgreSQL before continuing.

CREATE EXTENSION pg_walinspect;
CREATE EXTENSION pg_synthesize_wal;

SELECT pg_synthesize_wal_record(1024::bigint);
SELECT pg_synthesize_wal_record('test payload'::text);
SELECT pg_synthesize_wal_record(decode('deadbeef', 'hex'));
```

Each overload returns the `pg_lsn` of the inserted record. The size overload generates random bytes; the `text` and `bytea` overloads preserve caller-supplied payload bytes in the custom record.

### SQL Surface

- `pg_synthesize_wal_record(bigint)` writes a random payload of the requested size.
- `pg_synthesize_wal_record(text)` writes a text payload.
- `pg_synthesize_wal_record(bytea)` writes a binary payload.

All overloads are strict and parallel-unsafe. The upstream supports PostgreSQL `15+`, where custom WAL resource managers are available.

### Safety Boundary

Calling these functions consumes WAL bandwidth, archive space, replication bandwidth, and recovery time; very large inputs may also consume substantial backend memory. Redo intentionally performs no data change, but the records still flow through WAL, archiving, replication, and recovery. Run only on disposable or capacity-controlled systems, monitor WAL growth and replica lag, and remove the preload setting before returning a cluster to normal service.
