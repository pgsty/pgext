## Usage

Sources:

- [AlloyDB columnar engine overview](https://docs.cloud.google.com/alloydb/docs/columnar-engine/about?hl=en)
- [Configure the columnar engine](https://docs.cloud.google.com/alloydb/docs/columnar-engine/configure)
- [Manage columnar content manually](https://docs.cloud.google.com/alloydb/docs/columnar-engine/manage-content-manually)
- [AlloyDB supported extensions](https://docs.cloud.google.com/alloydb/docs/reference/extensions)

`google_columnar_engine` is the AlloyDB extension that exposes Google Cloud's in-memory column store and vectorized planner/executor for analytical scans, joins, and aggregates. It runs only inside AlloyDB; it has no standalone community package or independent upstream version.

### Enablement

Set the instance database flag `google_columnar_engine.enabled=on`. This restarts the instance. After it returns to service, a member of `alloydbsuperuser` can enable the SQL extension:

```sql
CREATE EXTENSION IF NOT EXISTS google_columnar_engine;

SELECT google_columnar_engine_add(
    relation => 'sales.orders',
    columns  => 'customer_id,created_at,total'
);

SELECT * FROM g_columnar_relations;

EXPLAIN (ANALYZE, COLUMNAR_ENGINE)
SELECT customer_id, sum(total)
FROM sales.orders
GROUP BY customer_id;
```

Automatic columnarization is enabled by default with the engine and can select useful columns from observed workload. `google_columnar_engine_add` manages content explicitly, `google_columnar_engine_drop` removes it from the store, and `google_columnar_engine_refresh('sales.orders')` rebuilds stale columnar blocks.

### Persistence and Memory

Manually added relations are not automatically restored after an instance restart. Use the `google_columnar_engine.relations` database flag for selections that must persist; conversely, dropping flag-managed content only lasts until the next restart unless the flag is updated. The engine reserves 30% of instance memory by default. Changing its memory-size flag also restarts the instance, so size the store and schedule flag changes operationally.

### Query and Table Boundaries

A scan can use the columnar engine only when all columns needed by that scan are loaded and use supported built-in types. Leaf partitions may be loaded individually, but non-leaf partitioned tables and foreign tables are unsupported. Very small tables (roughly fewer than 5,000 rows in the provider guidance) and indexed lookups may remain in the row store even when columns are loaded.

Updates invalidate columnar blocks. Monitor `invalid_block_count` versus `total_block_count` in `g_columnar_relations` and refresh heavily changed tables; frequent-update OLTP tables may be poor candidates. Use `EXPLAIN (COLUMNAR_ENGINE)` to confirm actual execution rather than inferring it from extension installation. Provider flags, supported types, and eligibility rules can change with the AlloyDB engine version, so validate them against the current instance documentation.
