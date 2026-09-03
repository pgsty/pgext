## Usage

Sources:

- [Official documentation](https://github.com/fossas/pg_fossa/blob/1ea06cfa082d51ab2a5cf722f6eb57fd7688928f/README.md)
- [Extension control file](https://github.com/fossas/pg_fossa/blob/1ea06cfa082d51ab2a5cf722f6eb57fd7688928f/pg_fossa.control)
- [Official repository](https://github.com/fossas/pg_fossa)

`pg_fossa` Archived SQL extension supporting FOSSA hasGraph data structures and queries.

### Enablement

Install the files for the intended server, then create `pg_fossa` in the target database:

```sql
CREATE EXTENSION pg_fossa;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT fossa_version();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `fossa_dependencies` | FUNCTION | Callable function from the reviewed install surface. |
| `fossa_version` | FUNCTION | Callable function from the reviewed install surface. |
| `array_intersect` | FUNCTION | Callable function from the reviewed install surface. |
| `array_intersect_agg` | AGGREGATE | Aggregate exposed by the extension. |
| `array_sort_unique` | FUNCTION | Callable function from the reviewed install surface. |
| `array_symmetric_difference` | FUNCTION | Callable function from the reviewed install surface. |
| `array_union` | FUNCTION | Callable function from the reviewed install surface. |
| `array_union_agg` | AGGREGATE | Aggregate exposed by the extension. |

### Operations and Boundaries

- Catalog lifecycle is archived; test upgrades, dump/restore, and server compatibility before production use.
