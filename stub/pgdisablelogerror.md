## Usage

Sources: [README](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/README.md), [v1.0 release](https://github.com/fmbiete/pgdisablelogerror/releases/tag/v1.0), [control file](https://github.com/fmbiete/pgdisablelogerror/blob/v1.0/pgdisablelogerror.control)

`pgdisablelogerror` suppresses PostgreSQL server log entries for configured SQLSTATE error codes. It is useful when expected application errors, such as duplicate-key violations, are too noisy in the server log.

### Enable The Hook

Load the module and restart PostgreSQL:

```conf
shared_preload_libraries = 'pgdisablelogerror'
```

Create the extension once in the `postgres` database:

```sql
CREATE EXTENSION pgdisablelogerror;
```

### Configure SQLSTATE Codes

Set `pgdisablelogerror.sqlerrcode` to a comma-separated list of SQLSTATE codes:

```conf
pgdisablelogerror.sqlerrcode = '23505,23503'
```

An empty or NULL value disables suppression:

```conf
pgdisablelogerror.sqlerrcode = ''
```

To identify SQLSTATE values in normal PostgreSQL logs, add `%e` to `log_line_prefix`.

### Caveats

- Version 1.0 supports PostgreSQL 14-18.
- This extension affects logging, not error behavior. Clients still receive the original error.
- Use narrow SQLSTATE lists. Suppressing broad error classes can hide real operational problems.
