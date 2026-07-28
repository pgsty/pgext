## Usage

Sources:

- [Official upstream README](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/README)
- [Official extension control file (shark.control)](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/shark/shark.control)
- [Official extension SQL (shark--1.0.sql)](https://github.com/shicheng0104/opengauss/blob/0495a2328db5f409b6e29eda1d671964f35168d4/contrib/shark/shark--1.0.sql)

`shark` — extension for D compatibility. Use it when porting or emulating the corresponding database API. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION shark;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dbcc_check_ident_no_reseed(varchar, boolean, boolean)` is an extension function and returns `varchar`.
- `dbcc_check_ident_reseed(varchar, int16, boolean)` is an extension function and returns `varchar`.
- `fetch_status()` is an extension function and returns `int`.
- `objectproperty(id INT, property VARCHAR)` is an extension function and returns `INT`.
- `rowcount()` is an extension function and returns `int`.
- `rowcount_big()` is an extension function and returns `bigint`.
- `spid()` is an extension function and returns `bigint`.
- `sys.day(abstime)` is an extension function and returns `float8`.
- `sys.day(date)` is an extension function and returns `float8`.
- `sys.day(timestamp(0) with time zone)` is an extension function and returns `float8`.
- `sys.day(timestamptz)` is an extension function and returns `float8`.
- `sys.object_id(IN object_name VARCHAR, IN object_type VARCHAR DEFAULT '')` is an extension function and returns `integer`.
- `sys.pltsql_call_handler()` is an extension function and returns `language_handler`.
- `sys.pltsql_inline_handler(internal)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `2.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
