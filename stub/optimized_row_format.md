## Usage

Sources:

- [Official documentation](https://github.com/davinderatgithub/optimized-row-format/blob/d3d28fbb2b6ae5e826da92bd8bd0010275daba12/README.md)
- [Extension control file](https://github.com/davinderatgithub/optimized-row-format/blob/d3d28fbb2b6ae5e826da92bd8bd0010275daba12/optimized_row_format.control)
- [Official repository](https://github.com/davinderatgithub/optimized-row-format)

`optimized_row_format` Experimental table access method for compact, cache-friendly row storage.

### Enablement

Merge `optimized_row_format` into the existing preload list, restart PostgreSQL, and then create `optimized_row_format` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'optimized_row_format'
```

```sql
CREATE EXTENSION optimized_row_format;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Load the extension
CREATE EXTENSION optimized_row_format;

-- Create a test table
CREATE TABLE test_table (
    id integer,
    name text,
    value bigint,
    flag boolean
) USING optimized_row_format;

-- Insert test data
INSERT INTO test_table VALUES (1, 'test', 100, true);

-- Query the data
SELECT * FROM test_table;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `optimized_row_format_tableam_handler` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 17; do not infer unlisted majors.
- Preloading `optimized_row_format` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
