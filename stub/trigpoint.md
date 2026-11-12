## Usage

Sources:

- [Release metadata](https://api.pgxn.org/dist/trigpoint/0.5.0/META.json)
- [Operator guide](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/README.md)
- [Control file](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/trigpoint.control)
- [Installation SQL](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/pgext/trigpoint--0.5.sql)
- [Changelog](https://api.pgxn.org/src/trigpoint/trigpoint-0.5.0/CHANGELOG.md)

`trigpoint` adds a vector index access method for pgvector on PostgreSQL 16–18. It combines an HNSW router over compact sketches with calibrated or exact pruning before fetching heap rows. Distribution 0.5.0 installs SQL version 0.5. The release requires pgvector 0.8.0 or later.

### Core Workflow

Install the native library as an administrator and create both extensions:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION trigpoint;
CREATE TABLE items (id bigint, embedding vector(3));
CREATE INDEX items_embedding_idx
    ON items USING trigpoint (embedding vector_l2_ops);

SELECT id, embedding <-> '[1,2,3]'::vector AS distance
FROM items
ORDER BY embedding <-> '[1,2,3]'::vector
LIMIT 10;
```

The `vector_l2_ops` class supports L2 distance; `vector_cosine_ops` supports cosine distance through `<=>`. Only one vector column is indexed. Inner product, backward scans, covering columns and parallel query scans are unsupported.

### Search Controls

The default `trigpoint.iterative_scan = relaxed_order` expands filtered scans in batches; later batches can contain nearer rows. `strict_order` keeps returned distances ordered by dropping out-of-order candidates. For exact prefixes, also enable `trigpoint.exact`:

```sql
BEGIN;
SET LOCAL trigpoint.iterative_scan = strict_order;
SET LOCAL trigpoint.exact = on;
SELECT id
FROM items
ORDER BY embedding <-> '[1,2,3]'::vector
LIMIT 10;
COMMIT;
```

`trigpoint.max_scan_tuples` defaults to 20000; `trigpoint.max_probes` defaults to 0, meaning all cells are permitted. Hitting a cap can return fewer rows without an error. Filters on other columns run above the index; an index qualification on the vector column itself is unsupported. Approximate recall depends on routing and calibration settings.

### Memory and Preload

An index build needs roughly six bytes per row per vector dimension in `maintenance_work_mem`, plus sufficient dynamic shared memory. On PostgreSQL 17–18 the shared image registry needs no preload. On PostgreSQL 16, sharing requires adding the library to the existing preload list and restarting:

```ini
shared_preload_libraries = 'trigpoint'
```

Without that preload on PostgreSQL 16, the extension still works but each backend keeps its own index image. `trigpoint_shared_available()` reports whether sharing is available. The superuser setting `trigpoint.max_shared_mb` caps shared images; `trigpoint.on_shared_cap = private` permits a private fallback instead of an error.

### Maintenance and Access

```sql
SELECT trigpoint_version();
SELECT trigpoint_info('items_embedding_idx'::regclass);
SELECT * FROM trigpoint_check('items_embedding_idx'::regclass, false);
SELECT * FROM trigpoint_shared_total();
VACUUM items;
```

`VACUUM` records deletions and compacts the image when thresholds are reached; `REINDEX` rebuilds it. Catalog upgrades use `ALTER EXTENSION trigpoint UPDATE`; dumping and restoring rebuilds indexes. The release documents page format 3 and tested PostgreSQL 17→18 upgrades, not an unrestricted guarantee for future formats.

Helpers that accept an index require permission to read its table; deep checking requires ownership. `trigpoint_image()` and `trigpoint_route()` expose vectors or neighbors and are revoked from `PUBLIC`. Release builds exclude page-modification test hooks. Use PostgreSQL data checksums as an additional corruption defense; the structural checker does not prove every model value is correct.
