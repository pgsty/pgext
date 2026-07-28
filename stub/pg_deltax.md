## Usage

Sources:

- [Official upstream README](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/README.md)
- [Official extension control file (pg_deltax.control)](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/pg_deltax.control)
- [Official implementation source](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/src/lib.rs)

`pg_deltax` — DeltaX (δx) is a PostgreSQL extension offering compression and columnar storage for time-series data. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_deltax;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.2.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
