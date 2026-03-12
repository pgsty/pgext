


## Usage

> [pg_ivm: Incremental View Maintenance for PostgreSQL](https://github.com/sraoss/pg_ivm)

The `pg_ivm` extension provides Incremental View Maintenance (IVM), updating materialized views by applying only incremental changes rather than recomputing from scratch. Views are updated immediately in AFTER triggers when base tables are modified.

```sql
CREATE EXTENSION pg_ivm;
```

### Configuration

Add `pg_ivm` to preload libraries for correct maintenance:

```ini
shared_preload_libraries = 'pg_ivm'
```

### Functions

#### create_immv

```sql
pgivm.create_immv(immv_name text, view_definition text) RETURNS bigint
```

Creates an Incrementally Maintainable Materialized View (IMMV). Triggers are automatically created to keep the view updated. A unique index is created automatically if possible.

```sql
SELECT pgivm.create_immv('myview', 'SELECT * FROM mytab');
```

#### refresh_immv

```sql
pgivm.refresh_immv(immv_name text, with_data bool) RETURNS bigint
```

Completely replaces IMMV contents. With `with_data = false`, the IMMV becomes unpopulated and triggers are dropped. With `with_data = true`, triggers and indexes are recreated.

```sql
SELECT pgivm.refresh_immv('myview', true);
```

#### get_immv_def

```sql
pgivm.get_immv_def(immv regclass) RETURNS text
```

Returns the reconstructed SELECT command for an IMMV.

### IMMV Metadata Catalog

The `pgivm.pg_ivm_immv` catalog stores IMMV information:

| Column | Type | Description |
|--------|------|-------------|
| `immvrelid` | regclass | OID of the IMMV |
| `viewdef` | text | Query tree for the view definition |
| `ispopulated` | bool | Whether IMMV is currently populated |

### Examples

Create an IMMV with aggregates:

```sql
SELECT pgivm.create_immv('immv_agg',
    'SELECT bid, count(*), sum(abalance), avg(abalance)
     FROM pgbench_accounts JOIN pgbench_branches USING(bid) GROUP BY bid');
```

Updates to base tables are reflected automatically:

```sql
UPDATE pgbench_accounts SET abalance = abalance + 1000 WHERE aid = 4112345;
SELECT * FROM immv_agg WHERE bid = 42;  -- aggregates updated automatically
```

List all IMMVs:

```sql
SELECT immvrelid AS immv, pgivm.get_immv_def(immvrelid) AS def
FROM pgivm.pg_ivm_immv;
```

Drop an IMMV with `DROP TABLE`:

```sql
DROP TABLE myview;
```

### Disable/Enable Maintenance

Disable immediate maintenance before bulk modifications, then refresh:

```sql
SELECT pgivm.refresh_immv('myview', false);   -- disable
-- ... bulk modifications ...
SELECT pgivm.refresh_immv('myview', true);    -- refresh and re-enable
```

### Supported Query Features

- Inner and outer joins (including self-join)
- `DISTINCT` clause
- Aggregate functions: `count`, `sum`, `avg`, `min`, `max`
- Simple subqueries in `FROM` clause
- `EXISTS` subqueries in `WHERE` clause
- Simple CTEs (`WITH` queries)
- `GROUP BY` clause
