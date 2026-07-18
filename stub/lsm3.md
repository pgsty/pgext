## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/README.md)
- [Version 1.0 SQL objects](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3--1.0.sql)
- [C implementation](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3.c)
- [Extension control file](https://github.com/postgrespro/lsm3/blob/3bac10420e7738ddbd78e0eea7e3967afdd35aac/lsm3.control)

`lsm3` implements an LSM-like index access method from three ordinary B-tree structures: an active top index, a merge top index, and a base index. It requires `shared_preload_libraries` and a restart before creating the SQL extension.

```sql
CREATE EXTENSION lsm3;
CREATE TABLE event_log (id bigint, payload text);
CREATE INDEX event_log_id_lsm ON event_log USING lsm3 (id);
```

Each LSM3 index starts its own merge background worker. Size `max_worker_processes`, `lsm3.max_indexes`, and `lsm3.top_index_size` accordingly, and benchmark read amplification because scans merge three indexes.

Parallel scans and array keys are unsupported. Most importantly, LSM3 cannot enforce uniqueness: the index option `unique=true` is only a lookup optimization and does not reject duplicates. Never use it as a unique constraint. Version 1.0 has no current compatibility matrix; validate crash recovery, WAL replay, vacuum, concurrent merge, reindex, and upgrade behavior with disposable data.
