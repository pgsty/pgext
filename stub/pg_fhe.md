## Usage

Sources:

- [Official documentation](https://github.com/FHE-Postgres/pg_fhe/blob/e15047067c197f07dc7b85bd42e72a0cb9ec932e/pg_fhe/README.md)
- [Extension control file](https://github.com/FHE-Postgres/pg_fhe/blob/e15047067c197f07dc7b85bd42e72a0cb9ec932e/pg_fhe/sql/pg_fhe.control)
- [Official repository](https://github.com/FHE-Postgres/pg_fhe)

`pg_fhe` Microsoft SEAL-backed CKKS operations over encrypted values stored in PostgreSQL.

### Enablement

Install the files for the intended server, then create `pg_fhe` in the target database:

```sql
CREATE EXTENSION pg_fhe;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE TABLE test_ckks_mult as
    SELECT id, ckks_mult(data, 2.0) as data FROM test;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ckks_mult` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
