## Usage

Sources:

- [pg_bedtools_rs README at the reviewed commit](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/README.md)
- [pg_bedtools_rs.control at the reviewed commit](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/pg_bedtools_rs.control)
- [Cargo package metadata](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/Cargo.toml)
- [bed_merge implementation](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/src/bed/merge.rs)
- [Generated extension SQL](https://github.com/shencangsheng/pg_bedtools_rs/blob/a04ef932a131323de5f9d7f85e78d0e5affd101c/sql/pg_bedtools_rs--0.1.0.sql)

`pg_bedtools_rs` exposes `bed_merge(table_name, padding, conditions)` for merging overlapping genomic intervals with the Rust bedrs library. The input relation must provide `chromosome`, `pos_start`, and `pos_end` columns; optional padding expands both ends before merging.

### Merge Intervals

```sql
CREATE EXTENSION pg_bedtools_rs;

CREATE TABLE genomic_interval (
  chromosome varchar NOT NULL,
  pos_start integer NOT NULL,
  pos_end integer NOT NULL
);
INSERT INTO genomic_interval VALUES
  ('chr1', 5, 10), ('chr1', 7, 15), ('chr1', 22, 30);

SELECT * FROM bed_merge('genomic_interval', 0, '1 = 1');
```

### Caveats

- Both `table_name` and `conditions` are interpolated directly into SQL with Rust `format!`. They are not identifier-quoted or parameterized; accept them only from trusted code and do not expose the function to untrusted callers.
- The implementation loads all selected intervals into backend memory before returning merged rows. Validate interval bounds and size large jobs conservatively.
- Version `0.1.0` supports PostgreSQL 13 only. One secondary upstream document calls the function `bed_overlap`, but the reviewed Rust export and README use `bed_merge`.
- The control file requires `superuser` installation and marks the extension untrusted. Installation privilege does not make dynamically supplied filter text safe.
