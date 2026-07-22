## Usage

Sources:

- [Official README](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/README.md)
- [Extension SQL for version 1](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number--1.sql)
- [Type implementation](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number.c)

`number` is a variable-width signed 64-bit integer type. It stores zero in one byte and other values in two to nine bytes depending on magnitude, while preserving the same numeric range as `bigint`. It is intended for integer columns dominated by small values where reduced row and index size may outweigh extension and conversion overhead.

### Core Workflow

Install the extension, use `number` as a column type, and compare its storage with `bigint` on representative data:

```sql
CREATE EXTENSION number;

CREATE TABLE compact_counters (
  counter_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value number NOT NULL
);

INSERT INTO compact_counters (value)
VALUES (0), (127), (128), (2147483648);

SELECT value, pg_column_size(value)
FROM compact_counters
ORDER BY value;
```

Text input and output use ordinary signed integer notation. Values outside the signed 64-bit range are rejected; this is not an arbitrary-precision decimal type and has no scale.

### Casts, Comparison, and Indexing

- Assignment casts convert `integer` and `bigint` to `number`.
- An implicit cast converts `number` to `bigint`, allowing PostgreSQL's integer arithmetic to handle expressions.
- Version 1 defines comparison operators between two `number` values and cross-type comparisons with `bigint`.
- The default `number_ops` B-tree operator class supports ordered indexes.

```sql
CREATE INDEX compact_counters_value_idx
ON compact_counters (value);

SELECT *
FROM compact_counters
WHERE value >= 100::bigint
ORDER BY value;
```

The README's statement that the type has no operators predates comparison support present in the version 1 SQL; the SQL file is the authoritative object surface. There are no native arithmetic operators returning `number`, so inspect implicit casts and result types in expressions where that distinction matters.

### Operational Notes

Smaller values save space, but large values use up to nine bytes before normal tuple overhead and can be larger than fixed-width `bigint`. Benchmark the actual distribution, table compression, indexes, comparisons, and write workload rather than assuming a net saving.

Version 1 is from 2017, is relocatable, and declares no preload or extension dependency. The custom on-disk representation makes the extension binary part of backup and upgrade compatibility: install and test it on the destination before restoring dependent columns. Validate current PostgreSQL header compatibility, logical replication behavior, dump/restore, binary upgrades, casts used by ORMs, and query plans before production use.
