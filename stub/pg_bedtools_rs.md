## Usage

Sources:

- [Official README at the catalog revision](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/README.md)
- [Rust implementation at the catalog revision](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/src/lib.rs)
- [Committed extension SQL at the catalog revision](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/sql/pg_bedtools_rs--0.1.0.sql)

`pg_bedtools_rs` 0.1.0 is a pgrx experiment that merges genomic intervals with the Rust bedrs library. It expects a relation containing chromosome, start, and end columns. Upstream targets PostgreSQL 13, and the pinned repository does not contain a directly usable generated extension SQL artifact.

### Core Workflow

First build a complete pgrx package and verify that its install script actually declares the function. Then use a trusted, fixed relation and predicate.

```sql
CREATE TABLE intervals (
    chromosome text NOT NULL,
    pos_start integer NOT NULL,
    pos_end integer NOT NULL
);

INSERT INTO intervals VALUES
    ('chr1', 10, 20),
    ('chr1', 18, 30),
    ('chr2', 5, 8);

SELECT *
FROM bed_merge('intervals', 0, '1 = 1');
```

### Function Index

- `bed_merge(text, integer, text)` returns chromosome, start, and end columns after sorting and merging intervals.
- The second argument is padding, defaulting to zero. It expands interval starts and ends before merging.
- The third argument is a SQL condition string, defaulting to a true predicate.

### Installation and Safety Boundaries

- The source expects columns named chromosome, pos_start, and pos_end, and upstream reports PostgreSQL 13 support only.
- The pinned committed SQL file recursively issues CREATE EXTENSION instead of declaring the Rust function. A source checkout is therefore not an installable artifact by itself; generate/package with the matching pgrx toolchain and inspect the resulting SQL.
- Table and condition strings are interpolated into SQL rather than identifier-quoted or parameter-bound. Never pass user-controlled values; expose only reviewed wrappers over fixed relations and predicates.
- The control file marks the extension superuser-only and non-relocatable. Validate interval ranges and NULL handling before production use.
