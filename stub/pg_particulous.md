## Usage

Sources:

- [pg_particulous README at the reviewed commit](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/README.md)
- [Version 1.0 install SQL](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous--1.0.sql)
- [C helper implementation](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/src/pg_particulous.c)
- [Upstream migration test](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/sql/vanilla_to_pathman.sql)

`pg_particulous` is experimental PostgreSQL 10-era migration code for converting native declarative partition metadata into `pg_pathman` metadata. Version `1.0` implements only the native-to-`pg_pathman` direction despite the README's “vice versa” description. It directly rewrites system catalogs and storage metadata, so treat it as historical migration research rather than a routine extension.

### Migration Entry Point

The extension requires `plpgsql` and `pg_pathman`. A PostgreSQL 10 native-partitioned table can supply its partition expression implicitly; a non-native relation requires an explicit expression.

```sql
CREATE EXTENSION pg_pathman;
CREATE EXTENSION pg_particulous;

SELECT migrate_to_pathman(
  'public.measurements'::regclass,
  'recorded_at'
);
```

`migrate_to_pathman(regclass, text, boolean)` returns a boolean. For a native partitioned relation it calls `build_vanilla_part_key`, recursively runs `desugar_vanilla`, adds check constraints to former partitions, and finally calls `add_to_pathman_config`. In the reviewed SQL, the `run_tests` argument is accepted but not used.

Other helper functions inspect PostgreSQL 10 partition keys and bounds, build check conditions, and identify native partitions. They expose implementation machinery, not a supported general partition-management API.

### Destructive Boundary

`desugar_vanilla` locks `pg_class` and `pg_inherits` in exclusive mode and the target relation in `ACCESS EXCLUSIVE` mode. It swaps `relkind`, `relfilenode`, transaction-age fields, and `pg_partitioned_table` references with a temporary table, clears native partition flags, and drops that temporary table. These operations depend on exact PostgreSQL 10 internal catalog layouts.

### Operational Notes

- Do not run this on modern PostgreSQL or a production relation based only on extension-install success. The repository has not been updated since 2017 and targets unstable internal APIs.
- The migration has no implemented reverse operation. A successful call is not evidence that constraints, indexes, triggers, foreign keys, ownership, grants, sequences, replication, and query plans remain equivalent.
- Rehearse the exact schema on a disposable physical copy, take and verify a recoverable backup, stop concurrent writes, and compare catalog state and query results before and after migration.
- The function requires privileges to lock and alter the relation, modify system catalogs, and register `pg_pathman` metadata. Limit execution to a migration superuser and remove the extension after the experiment if it is no longer required.
