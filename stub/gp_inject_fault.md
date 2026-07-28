## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/README)
- [Official extension control file (gp_inject_fault.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/gp_inject_fault.control)
- [Official extension SQL (gp_inject_fault--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/gp_inject_fault--1.0.sql)

`gp_inject_fault` — infinite_loop loop until query cancel or terminate signal is received. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_inject_fault;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `force_mirrors_to_catch_up()` is an extension function and returns `VOID`.
- `gp_inject_fault(faultname text, type text, db_id int4)` is an extension function and returns `text`.
- `gp_inject_fault(faultname text, type text, db_id int4, gp_session_id int4)` is an extension function and returns `text`.
- `gp_inject_fault(faultname text, type text, ddl text, database text, tablename text, start_occurrence int4, end_occurrence int4, extra_arg int4, db_id int4)` is an extension function and returns `text`.
- `gp_inject_fault(faultname text, type text, ddl text, database text, tablename text, start_occurrence int4, end_occurrence int4, extra_arg int4, db_id int4, gp_session_id int4)` is an extension function and returns `text`.
- `gp_inject_fault_infinite(faultname text, type text, db_id int4)` is an extension function and returns `text`.
- `gp_wait_until_triggered_fault(faultname text, numtimestriggered int4, db_id int4)` is an extension function and returns `text`.
- `insert_noop_xlog_record()` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
