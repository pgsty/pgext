## Usage

Source: [README](https://github.com/plv8/plv8/blob/r3.2/README.md), [Docs site](https://plv8.github.io/), [Built-ins](https://github.com/plv8/plv8/blob/r3.2/docs/BUILTINS.md), [Runtime configuration](https://github.com/plv8/plv8/blob/r3.2/docs/CONFIGURATION.md), [Tag v3.2.4](https://github.com/plv8/plv8/tree/v3.2.4)

`plv8` is a trusted JavaScript procedural language for PostgreSQL, powered by the V8 engine. Upstream currently tags the extension as `v3.2.4`; Pigsty's `3.2.4-2` package version is a packaging revision rather than a new upstream extension release.

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

- Current docs state support for PostgreSQL 13 and above.
- Each session has its own global JavaScript runtime; switching roles initializes a separate runtime context.
- `plv8.execution_timeout` only applies when the extension is compiled with execution-timeout support.
