## Usage

Sources:

- [Official upstream README](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/README)
- [Official extension control file (sqlbench.control)](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/src/storeproc/pgsql/c/sqlbench.control)
- [Official extension SQL (sqlbench--1.0.0.sql)](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/src/storeproc/pgsql/c/sqlbench--1.0.0.sql)

`sqlbench` — a forked project from dbt2, can do a standard TPC-C test. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION sqlbench;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `delivery(INTEGER, INTEGER)` is an extension function and returns `INTEGER`.
- `new_order(INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER,…)` is an extension function and returns `INTEGER`.
- `order_status(INTEGER, INTEGER, INTEGER, TEXT)` is an extension function and returns `SETOF`.
- `payment(INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, TEXT, REAL)` is an extension function and returns `INTEGER`.
- `stock_level(INTEGER, INTEGER, INTEGER)` is an extension function and returns `INTEGER`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
