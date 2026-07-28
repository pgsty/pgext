## Usage

Sources:

- [Official upstream README](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/README.md)
- [Official extension control file (pg_change_point_detection.control)](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/pg_change_point_detection.control)
- [Official extension SQL (pg_change_point_detection--1.0.sql)](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/pg_change_point_detection--1.0.sql)

`pg_change_point_detection` — pg_changepoint is a PostgreSQL extension for detecting change points in the table data. It is a port of Andrey Akinshin's implementation of the ED-PELT algorithm. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_change_point_detection;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_change_point_detection(data double precision[])` is an extension function and returns `integer[]`.
- `pg_change_point_detection_in_column(table_name text, column_name text, order_column text DEFAULT NULL)` is an extension function and returns `integer[]`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
