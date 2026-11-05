## Usage

Sources:

- [Installation and upgrades](https://vchord.pgsty.com/docs/vchord/getting-started/installation/)
- [Partitioned vector indexes](https://vchord.pgsty.com/docs/vchord/usage/indexing/)
- [Graph indexes and result ordering](https://vchord.pgsty.com/docs/vchord/usage/graph-index/)
- [Per-index search defaults](https://vchord.pgsty.com/docs/vchord/usage/fallback-parameters/)
- [Quantized vector types](https://vchord.pgsty.com/docs/vchord/usage/quantization-types/)
- [Recall evaluation](https://vchord.pgsty.com/docs/vchord/usage/measure-recall/)
- [1.2.0 release notes](https://vchord.pgsty.com/blog/release/vchord-1.2.0/)
- [1.2.0 control file](https://github.com/pgsty/vchord/blob/1.2.0/vchord.control)

`vchord` 1.2.0 provides approximate vector search on PostgreSQL 14–18. It uses pgvector's types and adds the partitioned `vchordrq` index and the disk graph `vchordg` index. It requires the `vector` extension, shared preloading, and superuser privileges to install. Graph indexing is documented as a preview feature.

### Enable and Search

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'vchord'
```

```sql
CREATE EXTENSION vchord CASCADE;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]'), ('[3,1,2]');
CREATE INDEX items_embedding_idx ON items
USING vchordrq (embedding vector_l2_ops);

SELECT id, embedding <-> '[3,1,2]' AS distance
FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

Small datasets can start with an unpartitioned index. For larger datasets, choose the partition hierarchy through TOML `build.internal.lists`, then set `vchordrq.probes` to the number of partitions to search at each level. More probes trade query work for recall; the probe list must match the hierarchy. The graph alternative uses the same column and distance operator classes:

```sql
CREATE INDEX items_embedding_graph_idx ON items
USING vchordg (embedding vector_l2_ops)
WITH (options = $$
m = 32
ef_construction = 64
$$, ef_search = 128);
```

Use the graph definition as an alternative to the partitioned index above. The recall example below assumes only the partitioned index exists; inspect the query plan when comparing index methods.

### Types, Operators, and Tuning

| Object | Purpose |
| --- | --- |
| `vector_l2_ops`, `<->` | Euclidean distance, ordered ascending. |
| `vector_ip_ops`, `<#>` | Negative inner product, ordered ascending. |
| `vector_cosine_ops`, `<=>` | Cosine distance, ordered ascending. |
| `halfvec` | Half-precision vectors, with corresponding `halfvec_*_ops` classes. |
| `rabitq4`, `rabitq8` | Quantized vector storage; conversion loses precision. |
| `sphere`, `<<->>`, `<<#>>`, `<<=>>` | Distance-radius predicates for L2, inner product, and cosine search. |
| `vchordrq.epsilon` | Reranking tradeoff; default 2.5 in 1.2.0, previously 1.9. |
| `vchordg.ef_search` | Graph search breadth; default 64. |
| `vchordrq_prewarm`, `vchordg_prewarm` | Load hot index pages; require table ownership or table-level `SELECT`. |

```sql
SELECT id FROM items
WHERE embedding <<->> sphere('[1,2,3]'::vector, 0.5);

SELECT dequantize_to_vector(quantize_to_rabitq8('[1,2,3]'::vector));
SELECT dequantize_to_halfvec(quantize_to_rabitq4('[1,2,3]'::halfvec));
```

Sphere radii use the distance operator's units. PostgreSQL rechecks graph sphere predicates, but approximate search can still miss matching rows. Dequantization cannot recover the original vector exactly.

Per-index `probes`, `epsilon`, and `ef_search` options supply search defaults. Explicit GUC settings, including database or server settings, override them. Resetting a session setting restores its configured default, which can still override the index.

### Ordering and Recall

Graph scans can discover closer rows after returning earlier candidates; nearly equal half-precision distances can also appear out of order. Increasing the search breadth reduces some ordering errors but does not guarantee strict ordering. Materialize a larger candidate set and sort it explicitly when ordering matters:

```sql
WITH candidates AS MATERIALIZED (
    SELECT id, embedding <-> '[3,1,2]' AS distance
    FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 100
)
SELECT * FROM candidates ORDER BY distance LIMIT 10;
```

This sorts the retrieved candidates; recall still depends on the candidate count and search parameters. For partitioned indexes, compare approximate results with an exact baseline:

```sql
SELECT vchordrq_evaluate_query_recall(
    query => $$SELECT ctid FROM items
                ORDER BY embedding <-> '[3,1,2]' LIMIT 5$$,
    exact_search => true
);
```

The query must return `ctid`; keep sequential scans enabled for the exact pass. Multi-vector MaxSim queries using `@#` require `exact_search => true` for recall evaluation.

### Maintenance and Upgrade

Keep normal vacuuming enabled: graph searches exclude deleted vertices after `VACUUM`, and index maintenance updates tuple/page statistics. `ANALYZE` alone does not rebuild or compact the index.

Install matching 1.2.0 library and SQL files, restart PostgreSQL to load the new library, then update each database as the extension owner or an administrator:

```sql
ALTER EXTENSION vchord UPDATE TO '1.2.0';
```

The release provides upgrade paths from 0.2.0 through 1.1.1 without reindexing. Recheck representative rankings after upgrading: 1.2.0 fixes MaxSim candidate estimates, graph sphere filters, deleted-vertex visibility, and quantization edge cases, and raises the default reranking budget. The extension is relocatable; include its chosen schema in the search path or qualify its objects.
