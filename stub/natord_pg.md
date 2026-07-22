## Usage

Sources:

- [Official README](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/README.md)
- [Official Rust implementation](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/src/lib.rs)
- [Official operator SQL](https://github.com/mpajkowski/natord_pg/blob/a1f2923934a1ae4179e559ea651641f1b975bef0/sql/operators.sql)

`natord_pg` 0.0.0 adds natural-order comparisons for `text`, so digit runs are compared numerically instead of purely byte by byte. It is useful for identifiers such as `file2` and `file10`, and includes a non-default B-tree operator class for matching predicates and indexes.

### Core Workflow

```sql
CREATE EXTENSION natord_pg;

CREATE TABLE artifacts (name text PRIMARY KEY);
INSERT INTO artifacts VALUES ('file10'), ('file2'), ('file1');

SELECT name
FROM artifacts
ORDER BY name USING <~~;

CREATE INDEX artifacts_name_natord_idx
  ON artifacts USING btree (name natord_ops);

SELECT name
FROM artifacts
WHERE name ~>= 'file2' AND name <~~ 'file20'
ORDER BY name USING <~~;
```

The natural ordering above places `file2` before `file10`. Because `natord_ops` is not declared as the default operator class for `text`, specify it explicitly when creating a natural-order index.

### Important Objects

- `natord_gt(text, text)`, `natord_lt(text, text)`, `natord_ge(text, text)`, and `natord_le(text, text)` implement natural comparisons.
- `natord_cmp(text, text)` returns -1, 0, or 1 and is the B-tree comparison support function.
- `<~~`, `~~>`, `<=~`, and `~>=` are natural less-than, greater-than, less-than-or-equal, and greater-than-or-equal operators.
- `natordfam` and `natord_ops` define the B-tree operator family and class.

### Operational Notes

The pinned manifest declares version `0.0.0` and build features for PostgreSQL 13 through 17, with PostgreSQL 17 as the default feature. The extension does not replace the database collation or ordinary `text` operators; queries must use the natural operators to match `natord_ops`. Before relying on uniqueness or range semantics, test punctuation, signs, leading zeros, non-ASCII text, and collation-sensitive data against the upstream `natord` crate behavior.
