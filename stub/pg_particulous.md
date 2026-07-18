## Usage

Sources:

- [pg_particulous README at the reviewed commit](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/README.md)
- [pg_particulous.control at the reviewed commit](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous.control)
- [Version 1.0 install SQL](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/pg_particulous--1.0.sql)
- [C helper implementation](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/src/pg_particulous.c)
- [Upstream migration test](https://github.com/funbringer/pg_particulous/blob/25918b502ffecd0425ea7eefafc8770b2391fbf6/sql/vanilla_to_pathman.sql)

`pg_particulous` is a work-in-progress migration helper for moving PostgreSQL 10 native partition metadata into `pg_pathman`. Its current public operation is `migrate_to_pathman`; helper functions such as `build_vanilla_part_key` and `desugar_vanilla` derive partition expressions and rewrite native partition metadata before registering the relation with `pg_pathman`.

### Migration Entry Point

```sql
CREATE EXTENSION pg_pathman;
CREATE EXTENSION pg_particulous;

SELECT migrate_to_pathman(
  'public.measurements'::regclass,
  'recorded_at'
);
```

Install `pg_particulous` in the same schema as `pg_pathman`. For a PostgreSQL 10 native-partitioned relation, the function derives the partition key; for another relation, supply the expression explicitly.

### Caveats

- Despite the README's “vice versa” wording, the reviewed version `1.0` install SQL implements only migration toward `pg_pathman`.
- `desugar_vanilla` takes `ACCESS EXCLUSIVE` locks and directly updates `pg_class` and `pg_partitioned_table`. A failure or server-version mismatch can corrupt partition metadata.
- The code targets PostgreSQL 10-era internal APIs and has not been updated since 2017. Treat it as migration research code, not a routine production tool; rehearse on a disposable copy with verified backups.
