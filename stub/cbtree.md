## Usage

Sources:

- [Official README](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/README.md)
- [Extension SQL](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree--1.0.sql)
- [Access-method implementation](https://github.com/yanliu8/cbtree/blob/871f57358cb297f22be6f419a8e52dc86567ab7c/cbtree.c)

`cbtree` is an experimental counted B-tree access method that treats an indexed `integer` value as a one-based insertion or lookup position. It was written against PostgreSQL 10.1 and is useful for studying positional indexes, not as a general-purpose or currently supported B-tree replacement.

### Core Workflow

Create an integer dummy column and a single-column `cbtree` index. The inserted value selects the desired position; a value larger than the current row count appends the row. Equality is the only supported search strategy.

```sql
CREATE EXTENSION cbtree;

CREATE TABLE demo (
    payload text,
    sequence_position integer NOT NULL
);
CREATE INDEX demo_position_idx
    ON demo USING cbtree (sequence_position);

INSERT INTO demo VALUES ('first', 1);
INSERT INTO demo VALUES ('new first', 1);
INSERT INTO demo VALUES ('append', 999);

SELECT payload
FROM demo
WHERE sequence_position = 1;
```

An index built over existing rows derives its initial sequence from heap scan order. The stored dummy value is an insertion instruction, not a durable rank that PostgreSQL automatically renumbers.

### Objects and Limits

- `cbtree` is the index access method.
- `int4_ops` is its only operator class, for `integer`; only `=` is registered.
- The implementation consumes only the first index value. Use one non-null `integer` column even though the handler advertises multicolumn support.
- The extension also installs `delta`, `delta_actual_pos`, `delta_sel`, `delta_del`, `delta_ins`, and `auto_vacuum`. These prototype helpers contain unsafe dynamic SQL and a broken trigger-time `VACUUM` path; do not expose or rely on them.

### Operational Notes

Use only positive positions. Zero, negative, or null inputs are not safely rejected in production builds. The access method does not support uniqueness, ordering, backward scans, null searches, array searches, or parallel index scans.

The code uses PostgreSQL 10-era index APIs and has no maintained compatibility matrix. Build and destructive-test it against the exact PostgreSQL source tree before any lab use; do not use it to hold authoritative ordering or production data.
