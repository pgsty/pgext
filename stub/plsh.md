


## Usage

> [plsh: PL/sh procedural language](https://github.com/petere/plsh)

`plsh` allows writing PostgreSQL functions as shell scripts. The function body must start with a shebang line specifying the interpreter.

```sql
CREATE EXTENSION plsh;
```

### Create Shell Functions

```sql
CREATE FUNCTION concat(text, text) RETURNS text AS '
#!/bin/sh
echo "$1$2"
' LANGUAGE plsh;

SELECT concat('hello ', 'world');  -- 'hello world'
```

### Argument Passing

Function arguments are passed as positional shell variables (`$1`, `$2`, etc.):

```sql
CREATE FUNCTION file_line_count(filename text) RETURNS int AS '
#!/bin/sh
wc -l < "$1"
' LANGUAGE plsh;
```

### Return Values

- Standard output becomes the return value (trailing newlines stripped)
- Empty output returns NULL
- Standard error output causes the function to abort with that error message
- Non-zero exit status triggers an error

```sql
CREATE FUNCTION system_uptime() RETURNS text AS '
#!/bin/sh
uptime
' LANGUAGE plsh;
```

### Database Access

Direct SPI access is not available, but `psql` can be used since libpq environment variables are preconfigured:

```sql
CREATE FUNCTION query_db(x int) RETURNS text AS $$
#!/bin/sh
psql -At -c "SELECT name FROM users WHERE id = $1"
$$  LANGUAGE plsh;
```

### Trigger Functions

Trigger context is available via environment variables:

| Variable | Description |
|----------|-------------|
| `PLSH_TG_NAME` | Trigger name |
| `PLSH_TG_WHEN` | `BEFORE`, `INSTEAD OF`, or `AFTER` |
| `PLSH_TG_LEVEL` | `ROW` or `STATEMENT` |
| `PLSH_TG_OP` | `DELETE`, `INSERT`, `UPDATE`, or `TRUNCATE` |
| `PLSH_TG_TABLE_NAME` | Target table name |
| `PLSH_TG_TABLE_SCHEMA` | Target table schema |

Event trigger variables: `PLSH_TG_EVENT`, `PLSH_TG_TAG`.

### Inline Execution

```sql
DO E'#!/bin/sh\necho "running shell command"' LANGUAGE plsh;
```

### Security

`plsh` should not be declared as `TRUSTED` since shell scripts have full OS-level access under the PostgreSQL user. Only superusers should create `plsh` functions.
