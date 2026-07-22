## Usage

Sources:

- [Official pg_comparator manual entry](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/README.pg_comparator)
- [Command-line manual source](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pg_comparator.pl)
- [Version 3.1 extension SQL](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp--3.1.sql)

`pgcmp` is the PostgreSQL-side support extension for the external `pg_comparator` command-line program. Together they compare keyed tables by hierarchical checksums and can optionally synchronize the second table to the first.

### Core Workflow

Install `pgcmp` in each PostgreSQL database involved, then run a read-only comparison from the operating system. The URL suffix identifies the table, key columns, and compared columns.

```sql
CREATE EXTENSION pgcmp;
```

```sh
pg_comparator --ask-pass \
  /family/calvin?id:c1,c2 \
  /family/hobbes
```

The command reports `UPDATE`, `INSERT`, and `DELETE` keys needed to make the second table match the first. `--synchronize` alone is a dry run; adding `--do-it` performs writes. Begin with comparison only, inspect every difference, back up the target, and rehearse recovery before enabling both options.

### Extension Surface

- `cksum2(text)`, `cksum4(text)`, and `cksum8(text)` provide the default checksum family.
- `fnv2(text)`, `fnv4(text)`, and `fnv8(text)` provide the FNV checksum family.
- The extension also installs XOR aggregates and casts used by the command's PostgreSQL comparison paths.

### Operational Boundaries

The command requires a non-empty key and may scan, lock, create temporary summary tables, and, in synchronization mode, write the target. Keep passwords out of process arguments by using `--ask-pass` or protected environment handling. Heterogeneous comparisons can disagree because of types, encodings, collations, NULL handling, or casts, and checksum collisions can produce false negatives. Version `3.1` is abandoned; prefer maintained replication or reconciliation tooling for current systems and validate load, transaction, timeout, and rollback behavior on copies before touching production data.
