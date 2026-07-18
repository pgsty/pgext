## Usage

Sources:

- [undam README at the reviewed commit](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/README.md)
- [undam.control at the reviewed commit](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.control)
- [Version 0.1 install SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam--0.1.sql)
- [Table-access-method implementation](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.c)
- [Upstream preload configuration](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.conf)
- [Upstream regression SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/sql/undam.sql)

`undam` is an experimental table access method for in-place updates backed by undo chains and generic WAL. The extension creates the `undam` access method and the diagnostic function `undam_rel_info(regclass)`. Its own README still calls the design a preview and notes that vacuum remains necessary.

### Preload and Trial

```conf
shared_preload_libraries = 'undam'
```

Restart PostgreSQL before creating the extension:

```sql
CREATE EXTENSION undam;
CREATE TABLE event_state (
  id bigint PRIMARY KEY,
  value bigint NOT NULL
) USING undam;
INSERT INTO event_state VALUES (1, 0);
UPDATE event_state SET value = value + 1;
VACUUM event_state;
SELECT * FROM undam_rel_info('event_state'::regclass);
```

### Caveats

- `_PG_init` refuses to load unless `undam` is in `shared_preload_libraries`; changing this setting requires a server restart.
- This is preview storage-engine code built against old PostgreSQL internal table-access-method APIs. The source contains explicit `NOT_IMPLEMENTED` paths, and upstream provides no current compatibility matrix.
- Version `0.1` should be evaluated only on disposable data with the exact target server build. Exercise crash recovery, WAL replay, indexes, long transactions, vacuum, backup, and upgrade paths before drawing conclusions; do not treat the example as production endorsement.
