## Usage

Sources:

- [Create a ScaNN index](https://cloud.google.com/alloydb/docs/ai/create-scann-index)
- [AlloyDB ScaNN index reference](https://cloud.google.com/alloydb/docs/reference/ai/scann-index-reference)
- [Supported AlloyDB extensions](https://cloud.google.com/alloydb/docs/reference/extensions)

`alloydb_scann` is Google AlloyDB's managed ScaNN access method for approximate nearest-neighbor searches over `vector` columns. Use it when vector search latency, index build time, or memory footprint matters more than exact-neighbor guarantees. It is available in AlloyDB and AlloyDB Omni, not as a portable community PostgreSQL extension.

### Core Workflow

Install the managed extension; `CASCADE` installs the required `vector` extension when necessary:

```sql
CREATE EXTENSION IF NOT EXISTS alloydb_scann CASCADE;
```

Create an automatically tuned cosine-distance index and run a matching nearest-neighbor query:

```sql
CREATE INDEX items_embedding_scann
ON items USING scann (embedding cosine);

SELECT id, embedding <=> $1::vector AS distance
FROM items
ORDER BY embedding <=> $1::vector
LIMIT 20;
```

The indexed distance function and query operator must represent the same metric. ScaNN supports `l2`, `dot_product`, and `cosine`. A simple `USING scann` definition uses automatic tuning and maintenance by default.

### Index Modes and Options

- `mode='AUTO'` lets AlloyDB choose and maintain the tree structure. `optimization='SEARCH_OPTIMIZED'` favors recall and latency; `BALANCED` reduces build cost.
- `auto_maintenance='ON'` is the automatic-mode default and keeps the index adapted as data changes.
- `mode='MANUAL'` exposes settings such as `num_leaves`, `quantizer`, and `max_num_levels`; use it only after measuring recall and latency on representative queries.
- Quantizers include `SQ8` and `FLAT`; additional choices and deeper trees can be preview features with separate instance flags.

### Requirements and Caveats

The source table needs stored `vector` embeddings. Empty, nearly empty, and partitioned tables have special creation constraints; automatically tuned indexes on tables with fewer than 10,000 rows require deferred index creation. Approximate results require application-level recall testing, and index build or maintenance consumes instance resources.

Version availability, supported PostgreSQL majors, preview flags, privileges, upgrades, and maintenance behavior are controlled by the AlloyDB service. Check the provider documentation for the target instance rather than assigning a community extension version.
