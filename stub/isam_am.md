## Usage

Sources:

- [README.md](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/README.md)
- [isam_am/isam_am--1.0.sql](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am--1.0.sql)
- [isam_am/isam_am--1.0--1.1.sql](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am--1.0--1.1.sql)
- [isam_am/Makefile](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/Makefile)
- [isam_am/isam_am.control](https://github.com/lMafa1813l/BD2-Gestor-isam-postgres/blob/b6b2d9fec5abadbafeaa01a6d17bca4b1597d1b0/isam_am/isam_am.control)

`isam_am` 1.1 implements an ISAM index for a PostgreSQL 18.6 course project. The installation path combines the 1.0 SQL definition with the 1.0-to-1.1 upgrade script.

### Core Workflow

```sql
CREATE EXTENSION isam_am;
CREATE TABLE isam_example (k integer);
SET isam_am.fill_pct = 80;
CREATE INDEX isam_example_idx ON isam_example USING isam_am (k);
SELECT * FROM isam_index_stats('isam_example_idx');
SELECT isam_reorg_needed('isam_example_idx', 25);
```

### Operational Boundaries

The default operator class covers `int4` comparisons. `isam_am.fill_pct` controls the fill setting used by the example. `isam_index_stats(regclass)` returns key, data-block, overflow and chain statistics; `isam_reorg_needed(regclass, float8)` compares overflow percentage with a threshold, defaulting to 25. Reorganization uses REINDEX and has its normal locking and storage cost. Installation defaults to superuser-only. Treat the supplied tests as project evidence, not a production support promise; no repository license was found.
