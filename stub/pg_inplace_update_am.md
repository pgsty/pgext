## Usage

Sources:

- [Official usage](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/.github/README.md)
- [Extension control](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/pg_inplace_update_am.control)
- [Versioned SQL](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/sql/pg_inplace_update_am--0.1.sql)
- [Initialization and settings](https://github.com/igor-suhorukov/pg_inplace_update_am/blob/7b868ca7aee12af6faf2c1c3804964d7aff81a10/src/iheap_handler.c)

`pg_inplace_update_am` 0.1 provides the `iheap` table access method and `dmbtree` index access method on unmodified PostgreSQL 19. Rows can be updated in place while undo records retain older versions for snapshots and rollback. This is early-development storage code; qualify recovery, replication and upgrades before using it with important data.

### Enable and Use

A superuser must preload the module and restart PostgreSQL to register its WAL resource managers:

```ini
shared_preload_libraries = 'pg_inplace_update_am'
```

Then install the SQL objects in the target database:

```sql
CREATE EXTENSION pg_inplace_update_am;
CREATE TABLE accounts (id int PRIMARY KEY, branch int, balance numeric)
    USING iheap;
CREATE INDEX ON accounts (branch);
INSERT INTO accounts VALUES (1, 1, 100);
UPDATE accounts SET balance = balance + 1 WHERE id = 1;
UPDATE accounts SET branch = 2 WHERE id = 1;
```

The ordinary index above defaults to `dmbtree`; primary keys and unique indexes default to `btree`. Changing an indexed column can stay in place with `dmbtree`; ordinary `btree` indexes require moving that row.

### Settings and Inspection

- `iheap.enable_inplace` enables in-place updates; disabling it moves each updated row.
- `iheap.default_index_am` defaults to `dmbtree`; `iheap.unique_index_am` defaults to `btree`.
- `iheap.foreground_undo_limit` defaults to 1024 kB. Larger aborts leave work for later accesses and VACUUM.
- `iheap.logical_decoding` defaults to `on`; `ignore` skips these tables' changes during decoding.
- `iheap_relation_stats`, `iheap_metapage_info` and `dmbtree_index_stats` inspect storage; `iheap_check`, `dmbtree_check` and `dmbtree_parent_check` verify it.
- `iheap_stat_tables` exposes statistics. `iheap_discard` removes undo no snapshot still needs; `dmbtree_copy_btree_opclasses` copies B-tree operator classes added after installation.

Most raw inspection, verification and maintenance functions have public execution revoked by the installation SQL. Keep undo cleanup and VACUUM in the operational plan; the existence of upstream test suites is not a production-readiness guarantee.
