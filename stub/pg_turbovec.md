## Usage

Sources:

- [v2.10.3 README](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/README.md)
- [v2.10.3 changelog](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/CHANGELOG.md)
- [Upgrade matrix](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/docs/UPGRADING.md)
- [Control file](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/pg_turbovec.control)
- [Filtering guide](https://codeberg.org/gregburd/pg_turbovec/src/tag/v2.10.3/docs/FILTERING.md)

`pg_turbovec` 2.10.3 provides the `turbovec.vector` type and compact vector indexes with candidate reranking against the original heap vectors. Flat search scans quantized codes; IVF additionally searches selected cells. Both are approximate unless the candidate set contains every true neighbour. **Upgrading any 1.x index to 2.x requires rebuilding it.**

### Create and Query Vectors

```sql
CREATE EXTENSION pg_turbovec;
SET search_path = public, turbovec;

CREATE TABLE items (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding turbovec.vector CHECK (turbovec.vector_dims(embedding) = 8)
);
INSERT INTO items (embedding) VALUES
  ('[1,2,3,4,5,6,7,8]'),
  ('[2,1,3,5,4,7,6,8]'),
  ('[8,7,6,5,4,3,2,1]');
CREATE INDEX items_embedding_idx ON items
USING turbovec (embedding turbovec.vec_cosine_ops)
WITH (bit_width = 4);

SELECT id, embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector AS distance
FROM items
ORDER BY embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector
LIMIT 3;
```

Indexed vectors must have a consistent dimension that is a multiple of 8; the type accepts up to 16,000 coordinates. The example uses eight-dimensional vectors so it is valid for indexing. Use `turbovec.vec_cosine_ops` with `<=>` and `turbovec.vec_ip_ops` with `<#>`; `<->` and `<+>` also provide exact distance operators.

### Choose an Index and Tune Queries

- `lists = 0` is the default flat quantized scan. Candidate reranking does not guarantee exact recall for every dataset.
- `WITH (lists = N)` enables IVF. Train on enough representative rows, choose the cell count for the corpus, and measure recall against an exact baseline; increasing the cell count alone does not ensure faster queries.
- `bit_width = 4` is the default. Two- and three-bit quantization are also available. `bit_width = 1` uses centered sign binary quantization, supports IVF, and reranks from the heap; it is a different scheme from TurboQuant.
- `WITH (graph = true)` is deprecated. Prefer flat or IVF for new indexes and consult the upstream migration guide for an existing graph index.

Indexes and queries can run without preload, but upstream recommends adding the library to `shared_preload_libraries` for its tuning GUCs. Merge it with existing entries and restart PostgreSQL; do not replace other required libraries:

```conf
shared_preload_libraries = 'pg_turbovec'
```

```sql
SELECT count(*) FROM pg_settings WHERE name LIKE 'turbovec.%';
SET turbovec.probes = 16;
SET turbovec.search_k = 64;
```

The settings query must return a nonzero count. A custom dotted parameter accepted by SET is not proof that the extension's GUC registered. Important controls are `turbovec.probes`, `turbovec.search_k`, `turbovec.oversample`, `turbovec.iterative_scan` and `turbovec.cache_size_mb`. Use partial indexes for stable filters or the documented allowlist/iterative-scan paths; compare filtered results with an exact baseline. Native partitions have separate indexes that must each be maintained.

### Upgrade from 1.x to 2.10.3

Schedule a maintenance window and retain the heap vectors plus a verified backup. Install the matching new binaries, update the SQL extension, and restart PostgreSQL so every backend uses the new library:

```sql
ALTER EXTENSION pg_turbovec UPDATE TO '2.10.3';
```

The index format changed from 7 to 8 in 2.0.0. Old indexes cannot be read in place: **rebuild every TurboVec index from the heap before resuming queries**, including every partition's index. In psql, generate and run the rebuild statements:

```sql
SELECT format('REINDEX INDEX %I.%I;', n.nspname, c.relname)
FROM pg_class c
JOIN pg_am a ON a.oid = c.relam
JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE a.amname = 'turbovec';
\gexec
```

This is ordinary blocking REINDEX. If choosing concurrent rebuilding, issue each command as a separate top-level statement, outside a transaction block or DO function, and plan for old-format scans to fail until the rebuilt index is available. Do not treat the SQL extension update alone as a completed 1.x migration.

The 2.10.2-to-2.10.3 patch preserves format 8 and does not itself require reindexing. Earlier transitions still need the actions in the upstream upgrade matrix. An already corrupted index needs repair even when a patch preserves its format; `turbovec.turbovec_check(regclass)` reports detected corruption and a reason.

### Compatibility and Maintenance

The control file fixes objects in schema `turbovec`, sets `superuser = false`, and is not relocatable. Upstream covers PostgreSQL 13–18 and labels PostgreSQL 19 experimental in this release; current Pigsty packages cover 14–18. Binary replacement, tuning preload and major-format migration have separate restart/rebuild requirements. Version 2.10.3 parallelizes cold-backend code repacking; it does not change existing index bytes. Budget index-build memory, temporary space, WAL and vacuum work from the actual corpus, rather than treating upstream benchmark numbers as guarantees.
