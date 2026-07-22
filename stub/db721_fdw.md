## Usage

Sources:

- [Upstream README](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/README.md)
- [Cargo package metadata](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/Cargo.toml)
- [Rust FDW implementation](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/src/lib.rs)
- [DB721 parser](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/src/parser/db721.rs)

`db721_fdw` version `0.0.0` is an educational Rust/pgrx foreign data wrapper for the columnar file format used by CMU 15-721 Project 1. A foreign table supplies a server-local `filename`; the implementation exposes read-only scan callbacks and experiments with projection and simple comparison-filter pushdown.

### Example

```sql
CREATE EXTENSION db721_fdw;
CREATE SERVER db721_server FOREIGN DATA WRAPPER db721_fdw;
CREATE FOREIGN TABLE chickens (
  identifier integer,
  farm_name varchar,
  weight_g real
) SERVER db721_server
  OPTIONS (filename '/data/data-chickens.db721', table 'Farm');
SELECT * FROM chickens WHERE identifier = 1;
```

The README example includes `table 'Farm'`, but the reviewed FDW collects that option and never consumes it; it always reads metadata from `filename` and does not use the embedded metadata table name to select among logical tables. Do not rely on `table` for routing or isolation.

This is course/project code, not a general storage FDW. The validator accepts options without checking them, several types and expressions are unsupported, planner estimates are fixed provisional values, and source paths contain unfinished panic branches. Installation is marked superuser-only and file reads occur as the PostgreSQL OS account. Use only with trusted files in an isolated test instance.
