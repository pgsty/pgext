## Usage

Sources:

- [isam.control](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/isam.control)
- [README.md](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/README.md)
- [sql/isam--1.0.sql](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/sql/isam--1.0.sql)
- [src/isam_am.c](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/src/isam_am.c)

`isam` is an academic PostgreSQL 18 index access method for one INTEGER key. It supports equality/range lookup and inserts with overflow pages; NULL keys are omitted and duplicate keys are allowed.

### Core Workflow

```sql
CREATE EXTENSION isam;
CREATE TABLE isam_sample (id integer);
INSERT INTO isam_sample VALUES (1), (2), (3);
CREATE INDEX isam_sample_idx ON isam_sample USING isam (id);
SELECT * FROM isam_sample WHERE id BETWEEN 1 AND 2;
SELECT isam_info('isam_sample_idx');
LOAD 'isam';
SHOW isam.fill_factor;
```

### Operational Boundaries

Superuser installation is required. The fill factor defaults to 70 and accepts 10–100. Index data lives in separate files below the PostgreSQL data directory, outside shared buffers and WAL: crash recovery is not provided. Readers and writers use a per-index lock, results are not ordered, and UNLOGGED tables are unsupported. Rebuilding limits overflow growth, but dropped/rebuilt indexes leave old files for manual cleanup. Use only for experiments with recoverable data; no explicit license was found.
