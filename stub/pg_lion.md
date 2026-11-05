## Usage

Sources:

- [Upstream source 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [Upstream source 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion.control)

`pg_lion` 0.3.0 adds roaring-bitmap inverted indexes for filtering, counts and grouped aggregates on PostgreSQL 16 and later.

### Core Workflow

The `lion` access method supports scalar equality/ranges, arrays, JSONB and text-search vectors. Index options `summaries` and `slices` trade extra storage and write work for range or bit-sliced aggregation. `lion_bm25_score` ranks text-search matches; ordinary PostgreSQL operators remain the query interface.

```sql
CREATE EXTENSION pg_lion;
CREATE TABLE lion_events (country text, amount integer);
INSERT INTO lion_events VALUES ('JP', 10), ('JP', 20), ('NZ', 30);
CREATE INDEX ON lion_events USING lion (country);
VACUUM ANALYZE lion_events;
SELECT count(*) FROM lion_events WHERE country = 'JP';
```

### Operational Boundaries

Superuser installation is required. Preloading is optional: append `pg_lion` to `shared_preload_libraries` and restart to enable its custom WAL format. An index built in that mode needs the preload for writes and VACUUM until rebuilt. Heap visibility and vacuum frequency affect index-only counting; do not assume every query avoids heap access.
