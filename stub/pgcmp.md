## Usage

Sources:

- [Upstream pg_comparator manual entry](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/README.pg_comparator)
- [Extension SQL definitions](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp.sql)
- [C support implementation](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp.c)

`pgcmp` is the PostgreSQL-side support extension for the `pg_comparator` command-line data comparison and synchronization tool. It supplies server-side checksums and casts used by the external program.

```sql
CREATE EXTENSION pgcmp;
SELECT extversion FROM pg_extension WHERE extname = 'pgcmp';
```

The source version is `3.1`, but the project is abandoned and its newest Git tag is older. Comparison and synchronization tools can place heavy scans and writes on both databases; use read-only comparison first, protect credentials, verify collation and type semantics, and rehearse recovery before enabling any repair action. Prefer maintained logical replication or purpose-built reconciliation for current systems.
