## Usage

Sources:

- [Official documentation](https://github.com/pgmoneta/pgmoneta_ext/blob/4b642dd88f3324c80aa9b30759a18df4773afe13/README.md)
- [Extension control file](https://github.com/pgmoneta/pgmoneta_ext/blob/4b642dd88f3324c80aa9b30759a18df4773afe13/sql/pgmoneta_ext.control)
- [Official repository](https://github.com/pgmoneta/pgmoneta_ext)

`pgmoneta_ext` Server-side block and delta-backup helpers for pgmoneta.

### Enablement

Install the files for the intended server, then create `pgmoneta_ext` in the target database:

```sql
CREATE EXTENSION pgmoneta_ext;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT pgmoneta_ext_version();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pgmoneta_ext_checkpoint` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_fips` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_get_file` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_get_files` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_get_oid` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_get_oids` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_promote` | FUNCTION | Callable function from the reviewed install surface. |
| `pgmoneta_ext_receive_file_chunk` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 14; do not infer unlisted majors.
