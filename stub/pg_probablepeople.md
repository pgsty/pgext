## Usage

Sources:

- [Official documentation](https://github.com/513analytics/pg_probablepeople/blob/a0587406b3d843e39f0bde52cc8587a972f317aa/README.md)
- [Extension control file](https://github.com/513analytics/pg_probablepeople/blob/a0587406b3d843e39f0bde52cc8587a972f317aa/pg_probablepeople.control)
- [Official repository](https://github.com/513analytics/pg_probablepeople)

`pg_probablepeople` CRF-based parsing of personal and organization names into structured components.

### Enablement

Install the files for the intended server, then create `pg_probablepeople` in the target database:

```sql
CREATE EXTENSION pg_probablepeople;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM parse_name('Mr. John Doe');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `parse_name` | FUNCTION | Callable function from the reviewed install surface. |
| `tag_name` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is preview; test upgrades, dump/restore, and server compatibility before production use.
