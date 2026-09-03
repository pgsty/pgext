## Usage

Sources:

- [Official documentation](https://github.com/waaeer/pg_daterangex/blob/397429615dd2d89fe96f99c47ff5092a0f35af86/README.md)
- [Extension control file](https://github.com/waaeer/pg_daterangex/blob/397429615dd2d89fe96f99c47ff5092a0f35af86/daterangex.control)
- [Official repository](https://github.com/waaeer/pg_daterangex)

`daterangex` Date range type whose displayed upper bound is inclusive, with casts and operators.

### Enablement

Install the files for the intended server, then create `daterangex` in the target database:

```sql
CREATE EXTENSION daterangex;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION daterangex;
SELECT daterange('[2024-01-01,2024-06-01]');
        daterange
-------------------------
 [2024-01-01,2024-06-02)

SELECT daterangex('[2024-01-01,2024-06-01]');
       daterangex
-------------------------
 [2024-01-01,2024-06-01]
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `date_minus` | FUNCTION | Callable function from the reviewed install surface. |
| `daterangex_canonical` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
