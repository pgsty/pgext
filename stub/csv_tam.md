## Usage

Sources:

- [Official README](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/README.md)
- [Table access-method implementation](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam.c)
- [Extension SQL](https://github.com/InnerLife0/csv_tam/blob/5fe49b6d58c2688ae15a824069a719a95e5dd2c8/csv_tam--0.0.1.sql)

`csv_tam` 0.0.1 is a teaching prototype that stores a relation's main fork as comma-delimited rows through a custom table access method. It demonstrates PostgreSQL's TableAM callbacks; it is not a transactional CSV storage engine and must not be used for production data.

### Core Workflow

Install the handler and explicitly select `csv_tam` when creating a disposable table:

```sql
CREATE EXTENSION csv_tam;

CREATE TABLE tam_demo (
  id integer,
  measured_at timestamptz,
  value numeric
) USING csv_tam;

INSERT INTO tam_demo VALUES (1, now(), 42.5);
SELECT * FROM tam_demo;
```

The implementation only serializes `int2`, `int4`, `int8`, `float4`, `float8`, `date`, `timestamp`, `timestamptz`, and `numeric`. It has no CSV quoting or escaping; NULL becomes an empty field and is therefore ambiguous.

### Supported Surface

Sequential insert and scan are the intended demonstration path. Update, delete, truncate, indexes, vacuum, analyze, parallel scans, rescan, sampling, speculative insertion, and relation-copy operations are unimplemented and raise errors. Planner row estimates are fixed rather than derived from real tuple statistics.

### Operational Notes

Writes append directly to the relation file without MVCC, WAL, crash recovery, or transaction rollback. A failed or rolled-back transaction can therefore leave bytes behind, and concurrent writes are unsafe. Several callbacks emit NOTICE messages that can reveal row content or file paths. Use an isolated test cluster, disposable data, and no sensitive values; never rely on this extension for durability or visibility semantics.
