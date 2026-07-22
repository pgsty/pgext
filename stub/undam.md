## Usage

Sources:

- [undam README at the reviewed commit](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/README.md)
- [undam.control at the reviewed commit](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.control)
- [Version 0.1 install SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam--0.1.sql)
- [Table-access-method implementation](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.c)
- [Upstream preload configuration](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.conf)
- [Upstream regression SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/sql/undam.sql)

`undam` is an experimental table access method that performs in-place tuple updates while retaining older versions in undo chains. It stores fixed-size chunks across the main and extension forks and uses PostgreSQL generic WAL. Upstream calls the design a preview, and vacuum is still required.

### Core Workflow

The module must allocate shared state during server startup:

```conf
shared_preload_libraries = 'undam'
```

Restart PostgreSQL, then create the extension and opt individual tables into the access method:

```sql
CREATE EXTENSION undam;

CREATE TABLE event_state (
  id bigint PRIMARY KEY,
  value bigint NOT NULL
) USING undam;

INSERT INTO event_state VALUES (1, 0);
UPDATE event_state SET value = value + 1 WHERE id = 1;
VACUUM event_state;

SELECT * FROM undam_rel_info('event_state'::regclass);
```

Existing heap tables are not converted by installing the extension. Use the `USING undam` clause only for tables whose data can be recreated during evaluation.

### Object Index

- `undam`: the table access method selected by `CREATE TABLE ... USING undam`.
- `undam_rel_info(regclass)`: reports implementation diagnostics for an undam relation, including scan and chunk counters.
- `undam.auto_chunk_size`, `undam.chunk_size`, `undam.alloc_chains`, and `undam.max_relations`: startup/configuration settings that control chunk selection and shared relation state. Keep the upstream defaults unless the exact source and workload have been tested.

### Operational Boundaries

- `_PG_init()` rejects normal session-time loading; `undam` must be present in `shared_preload_libraries`, and changing that setting requires a restart.
- Version `0.1` is preview storage-engine code tied to PostgreSQL internal table-access APIs. The reviewed source includes explicit `NOT_IMPLEMENTED` error paths and no current supported-version matrix.
- In-place updates do not eliminate MVCC history or maintenance. Undo-chain versions and reusable chunks still depend on vacuum and global visibility decisions.
- Because this replaces heap storage behavior, ordinary SQL success is insufficient validation. Test indexes, unique constraints, row locks, concurrent update/delete, long-running snapshots, vacuum, truncation, bulk load, TOAST-sized values, crash recovery, WAL replay, checksums, physical backup, replication, upgrade, and extension removal.
- Use only disposable data and the exact PostgreSQL build being evaluated. Do not move an undam relation between server builds or treat the sample as production endorsement without upstream compatibility and recovery evidence.
