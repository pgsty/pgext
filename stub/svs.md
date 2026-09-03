## Usage

Sources:

- [Official documentation](https://github.com/intel/SVS-Extension-for-PostgreSQL/blob/d504d1c49536fd8f9ff08cff0c0432d84978c8b2/README.md)
- [Extension control file](https://github.com/intel/SVS-Extension-for-PostgreSQL/blob/d504d1c49536fd8f9ff08cff0c0432d84978c8b2/svs.control)
- [Official repository](https://github.com/intel/SVS-Extension-for-PostgreSQL)

`svs` SVS Vamana approximate-nearest-neighbor index with LeanVec and LVQ compression for pgvector.

### Enablement

Install the files for the intended server, then create `svs` in the target database:

```sql
CREATE EXTENSION svs CASCADE;
```

The reviewed control or official workflow requires `vector`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION svs;

CREATE TABLE items (id serial PRIMARY KEY, embedding vector(768));

-- Create a Vamana index
CREATE INDEX ON items USING vamana (embedding vector_l2_ops);

-- Nearest-neighbor search
SELECT id FROM items ORDER BY embedding <-> '[0.1, 0.2, ...]' LIMIT 10;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `vamana` | ACCESS METHOD | Access method used in table or index definitions. |
| `vector_l2_ops` | OPERATOR CLASS | Operator class used by indexes. |
| `halfvec_cosine_ops` | OPERATOR CLASS | Operator class used by indexes. |
| `halfvec_ip_ops` | OPERATOR CLASS | Operator class used by indexes. |
| `halfvec_l2_ops` | OPERATOR CLASS | Operator class used by indexes. |
| `pg_stat_vamana_worker` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stat_vamana_worker_slot` | FUNCTION | Callable function from the reviewed install surface. |
| `svs_restart_worker` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
