## Usage

Sources:

- [README](https://github.com/plv8/plv8/blob/v3.2.5/README.md)
- [Built-ins](https://github.com/plv8/plv8/blob/v3.2.5/docs/BUILTINS.md)
- [Configuration](https://github.com/plv8/plv8/blob/v3.2.5/docs/CONFIGURATION.md)
- [Changes](https://github.com/plv8/plv8/blob/v3.2.5/Changes)
- [Control](https://github.com/plv8/plv8/blob/v3.2.5/plv8.control.common)
- [SQL template](https://github.com/plv8/plv8/blob/v3.2.5/plv8.sql.common)

`plv8` provides a trusted JavaScript procedural language powered by V8. This page follows upstream 3.2.5, including its effective-user crash fixes and PostgreSQL 19 support.

### Basic use

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;

CREATE FUNCTION plv8_test(keys text[], vals text[]) RETURNS json AS $$
  let out = {};
  for (let i = 0; i < keys.length; i++) out[keys[i]] = vals[i];
  return out;
$$ LANGUAGE plv8 IMMUTABLE STRICT;
```

### Common built-ins

- `plv8.elog(level, ...)`: emit PostgreSQL log or client messages.
- `plv8.execute(sql [, args])`: run SQL and return rows or affected-row count.
- `plv8.prepare(...)`, `PreparedPlan.execute()`, `PreparedPlan.cursor()`: prepared SPI access.
- `plv8.subtransaction(fn)`: run a group of SPI operations atomically.
- `plv8.find_function(...)`: call another PLV8 function by name.
- `plv8.memory_usage()`: inspect V8 heap usage for the current session.
- `plv8.run_script(source, name)`: evaluate named script text.

### Runtime settings

```sql
SET plv8.start_proc = 'plv8_init';
SET plv8.execution_timeout = 60;
SET plv8.memory_limit = 512;
```

- `plv8.start_proc`
- `plv8.v8_flags`
- `plv8.execution_timeout`
- `plv8.memory_limit`

### Caveats

- The 3.2.5 CI matrix covers PostgreSQL 14–19. PostgreSQL-major support is separate from the versions available in downstream packages.
- Creating the extension requires a superuser; the installed JavaScript language is trusted, which does not make the extension itself installable by an unprivileged user. `plv8_info()` has its public execution privilege revoked by the installation SQL.
- Version 3.2.5 fixes crashes when a cached function runs under a different effective user via `SECURITY DEFINER` or `SET ROLE`, or after `plv8_reset()`, and when an exception message cannot be converted to a string.
- Each session has its own global JavaScript runtime; switching roles initializes a separate runtime context.
- `plv8.execution_timeout` only applies when the extension is compiled with execution-timeout support.
