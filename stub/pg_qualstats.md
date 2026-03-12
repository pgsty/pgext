

## Usage

> [pg_qualstats: predicate statistics collector for PostgreSQL](https://github.com/powa-team/pg_qualstats)

pg_qualstats keeps statistics on predicates found in `WHERE` clauses and `JOIN` conditions. It tracks which columns are most frequently queried and which are queried together, enabling index recommendations.

### Viewing Predicate Statistics

```sql
-- Raw predicate statistics for current database
SELECT * FROM pg_qualstats;

-- Human-readable aggregated form
SELECT * FROM pg_qualstats_pretty;

-- Aggregated per-attribute statistics
SELECT * FROM pg_qualstats_all;

-- Predicates aggregated by query
SELECT * FROM pg_qualstats_by_query;
```

### Index Advisor

Generate index suggestions based on collected predicate statistics:

```sql
-- Suggest indexes (filtering predicates with >1000 rows and >30% selectivity)
SELECT v FROM json_array_elements(
    pg_qualstats_index_advisor(min_filter => 50)->'indexes') v;

-- Show predicates that couldn't be optimized
SELECT v FROM json_array_elements(
    pg_qualstats_index_advisor(min_filter => 50)->'unoptimised') v;
```

### Utility Functions

```sql
-- Get stored query text for a queryid
SELECT pg_qualstats_example_query(queryid);

-- Get all stored query texts
SELECT * FROM pg_qualstats_example_queries();

-- Reset all statistics
SELECT pg_qualstats_reset();
```

### Configuration

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_qualstats.enabled` | `true` | Enable/disable collection |
| `pg_qualstats.track_constants` | `true` | Track individual constant values |
| `pg_qualstats.max` | 1000 | Maximum tracked predicates and query texts |
| `pg_qualstats.resolve_oids` | `false` | Resolve OIDs at query time (uses more space) |
| `pg_qualstats.track_pg_catalog` | `false` | Track predicates on pg_catalog objects |
| `pg_qualstats.sample_rate` | -1 | Fraction of queries to sample (-1 = auto: 1/max_connections) |
