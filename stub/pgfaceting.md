

## Usage

> [pgfaceting: fast faceted search using inverted indexes with roaring bitmaps](https://github.com/cybertec-postgresql/pgfaceting)

The `pgfaceting` extension enables rapid facet counting via inverted indexes built with roaring bitmaps. Requires the `pg_roaringbitmap` extension.

```sql
CREATE EXTENSION pgfaceting;
```

### Facet Types

- **`plain_facet(column)`**: Use column values directly as facets
- **`datetrunc_facet(column, precision)`**: Apply date truncation (e.g., monthly/yearly buckets)
- **`bucket_facet(column, buckets)`**: Assign continuous variables to predefined ranges

### Key Functions

```sql
-- Create facet infrastructure for a table
SELECT pgfaceting.add_faceting_to_table(
    'products',
    'id',
    ARRAY[
        plain_facet('color'),
        plain_facet('size'),
        bucket_facet('price', ARRAY[0, 10, 50, 100, 500])
    ]
);

-- Run maintenance to merge incremental changes
SELECT pgfaceting.run_maintenance();

-- Merge deltas for a specific table
SELECT pgfaceting.merge_deltas('products');

-- Get top N facet values
SELECT pgfaceting.top_values('products', 10);

-- Count results with facet filters
SELECT pgfaceting.count_results('products', filters);
```

### Architecture

The extension maintains two auxiliary tables per indexed table: a main facets table with roaring bitmaps mapping facet values to row IDs, and a delta table for incremental changes between maintenance runs.

Currently supports only 32-bit integer ID columns.
