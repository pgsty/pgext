## Usage

Sources:

- [Official documentation](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/README.md)
- [Extension control file](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/pgck.control)
- [Build manifest](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/Cargo.toml)

`pgck` Concept Kernel runtime with NATS transport, SHACL validation, and RDF materialization.

### Enablement

Merge `pgck` into the existing preload list, restart PostgreSQL, and then create `pgck` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pgck'
```

```sql
CREATE EXTENSION pgck CASCADE;
```

The reviewed control or official workflow requires `pgrdf`, `pgcrypto`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT ckp.dispatch(
  'instance.create',
  '{"type":"urn:ckp:demo/type/Ship","name":"Aurora"}'::jsonb
);
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `ckp.dispatch` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 18; do not infer unlisted majors.
- Preloading `pgck` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
