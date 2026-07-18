## Usage

Sources:

- [Upstream storage and configuration documentation](https://github.com/adjust/pg_cryogen/blob/4c8b4f382215861e613b1e9be66233a10e058af5/README.md)
- [Extension control file](https://github.com/adjust/pg_cryogen/blob/4c8b4f382215861e613b1e9be66233a10e058af5/pg_cryogen.control)

`pg_cryogen` is an experimental append-only table access method for PostgreSQL 12 or later. It compresses heap tuples with `lz4` or `zstd` and implements index and bitmap scans.

```sql
CREATE EXTENSION pg_cryogen;
SET pg_cryogen.compression_method = 'zstd';
CREATE TABLE events (id bigint, payload jsonb) USING pg_cryogen;
INSERT INTO events VALUES (1, '{"ok":true}');
```

Only `INSERT` and `COPY` modify data, and the reviewed implementation permits inserts into only one such table per transaction. Each statement creates at least one separate block, so bulk loads are preferable. The repository is archived; pin the PostgreSQL major version, verify backup and recovery, and do not assume support for update, delete, vacuum behavior, replication, or future access-method APIs.
