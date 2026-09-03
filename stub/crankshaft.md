## Usage

Sources:

- [Official documentation](https://github.com/CartoDB/crankshaft/blob/4d7bc1acb5cd3167cb0c9e8341beb1f32964ae3e/README.md)
- [Extension control file](https://github.com/CartoDB/crankshaft/blob/4d7bc1acb5cd3167cb0c9e8341beb1f32964ae3e/release/crankshaft.control)
- [Official repository](https://github.com/CartoDB/crankshaft)

`crankshaft` Archived CARTO spatial-analysis functions for clustering, segmentation, and spatial statistics.

### Enablement

Install the files for the intended server, then create `crankshaft` in the target database:

```sql
CREATE EXTENSION crankshaft CASCADE;
```

The reviewed control or official workflow requires `plpython3u`, `postgis`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT cdb_crankshaft.cdb_crankshaft_version();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `_cdb_random_seeds` | FUNCTION | Callable function from the reviewed install surface. |
| `cdb_crankshaft.cdb_crankshaft_version` | FUNCTION | Callable function from the reviewed install surface. |
| `cdb_crankshaft.cdb_moran_local` | FUNCTION | Callable function from the reviewed install surface. |
| `cdb_crankshaft.cdb_moran_local_rate` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The extension fixes or creates schema objects under `cdb_crankshaft`; include them in privilege and backup review.
- Catalog lifecycle is archived; test upgrades, dump/restore, and server compatibility before production use.
