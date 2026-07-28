## Usage

Sources:

- [Official upstream README](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/README.md)
- [Official extension control file (workalendar.control)](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/workalendar.control)
- [Official extension SQL (workalendar--1.0.0.sql)](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/workalendar--1.0.0.sql)

`workalendar` — pg_workalendar is a plpythonu-based PostgreSQL extension for workalendar. Use it for the corresponding scheduling, temporal, or time-series workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION workalendar;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `holidays(year int, continent text, country text)` is an extension function and returns `SETOF`.
- `workon(venv text)` is an extension function and returns `void`.
- `holiday` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `plpython3u`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
