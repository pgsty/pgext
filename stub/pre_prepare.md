


## Usage

> [pre_prepare: Pre Prepare your Statement server side](https://github.com/dimitri/preprepare)

pre_prepare automatically prepares SQL statements at connection time so clients can directly use `EXECUTE` without calling `PREPARE` first.

### Setup

Configure in `postgresql.conf`:

```
preprepare.relation = 'preprepare.statements'
preprepare.at_init = on    -- auto-prepare on connection (requires local_preload_libraries)
```

Create a table to store statements:

```sql
CREATE TABLE preprepare.statements (name text PRIMARY KEY, statement text);
```

Insert statements (include the full `PREPARE` syntax):

```sql
INSERT INTO preprepare.statements VALUES ('test', 'prepare test as select 1');
```

### Functions

**`prepare_all()`** -- Prepare all statements from the configured relation:

```sql
SELECT prepare_all();
```

**`prepare_all('schema.table')`** -- Prepare statements from a specific table:

```sql
SELECT prepare_all('public.expensive_planning');
```

**`discard()`** -- Like `DISCARD ALL` but without `DEALLOCATE ALL` (preserves prepared statements):

```sql
SELECT discard();
```

### With PgBouncer

Set `connect_query` to auto-prepare on each server connection:

```ini
[databases]
foo = port=5432 connect_query='SELECT prepare_all();'
```

Avoid using `DISCARD ALL` as your `reset_query` (it would deallocate the prepared statements).
