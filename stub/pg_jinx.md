## Usage

Sources:

- [Official documentation](https://github.com/sourcewave/pg_jinx/blob/651430bb03e981ecef9d3e3e1c4441ccdbf522db/README.md)
- [Extension control file](https://github.com/sourcewave/pg_jinx/blob/651430bb03e981ecef9d3e3e1c4441ccdbf522db/pg_jinx.control)
- [Official repository](https://github.com/sourcewave/pg_jinx)

`pg_jinx` JNI bridge for PostgreSQL stored procedures, triggers, and foreign data wrappers written in Java.

### Enablement

Install the files for the intended server, then create `pg_jinx` in the target database:

```sql
CREATE EXTENSION pg_jinx;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM jproperties;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `jinx.TestTrigger` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `jinx.countRows` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.example_rs` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.fdw_handler` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.fdw_validator` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.getGravatar` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.getOptions` | FUNCTION | Callable function from the reviewed install surface. |
| `jinx.inline_handler` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
