## Usage

Sources:

- [Extension control file](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.control)
- [Version 1.0 SQL object](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks--1.0.sql)
- [Hook implementation](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/pglogical_output_nzhooks.c)
- [Build integration](https://github.com/troels/pghooks/blob/f82688238e25bb220478404eeeb4d6a825f014a1/Makefile)

`pglogical_output_nzhooks` is a minimal example of a `pglogical_output` row-filter hook. It is not a general-purpose filtering extension: the reviewed version emits only rows whose relation name is exactly `nzmagic` and logs a warning for every relation it examines.

### Integration Boundary

```sql
CREATE EXTENSION pglogical_output_nzhooks;
```

The extension installs `pglo_nzhooks_setup_fn(internal)`. The `internal` argument is a pointer supplied by the matching `pglogical_output` hook ABI; users cannot safely call this function from ordinary SQL. Configure a compatible `pglogical_output` build to invoke this setup symbol through its hook mechanism. The repository does not document a stable end-user configuration syntax, so confirm it from the exact pglogical version being built rather than guessing an option.

At build time the Makefile expects pglogical output-plugin headers in the sibling path `../pglogical/pglogical_output`. There is no declared SQL extension dependency or version check. Compile both components from compatible source and test that their private `PGLogicalHooks` and `PGLogicalRowFilterArgs` layouts match.

### Example-Only Behavior

The setup function writes `HELLO THERE!` at warning level. For every changed relation, the row filter logs its name at warning level, compares only the unqualified relation name with `nzmagic`, and returns false for every other table. Schemas are ignored, so same-named tables in different schemas all pass.

Do not deploy this example unchanged: it can silently suppress almost all logical changes and flood server logs. Fork it to implement explicit schema-qualified policy, remove per-row warning logging, add regression tests for inserts/updates/deletes and partitions, and validate replication-slot restart, failover, binary compatibility, and upgrade behavior.
