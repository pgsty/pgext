

## Usage

> [topn: top-N values approximation for PostgreSQL](https://github.com/citusdata/postgresql-topn)

Provides approximate top-N value tracking using an approximation algorithm that keeps a predefined number of frequent items and counters. Supports materialization, incremental updates, and merging across time intervals.

```sql
CREATE EXTENSION topn;
```

### Data Type

Uses `JSONB` to store the frequent items and their frequencies.

### Aggregates

| Function | Description |
|---|---|
| `topn_add_agg(text)` | Aggregate that creates a JSONB counter from a text column |
| `topn_union_agg(jsonb)` | Aggregate that merges multiple JSONB counter lists |

### Functions

| Function | Description |
|---|---|
| `topn(jsonb, n)` | Return the top-N elements and frequencies as rows |
| `topn_add(jsonb, text)` | Add a text value to a JSONB counter |
| `topn_union(jsonb, jsonb)` | Merge two JSONB counter lists |

### Configuration

- `topn.number_of_counters` -- number of counters to track (default: 1000)

### Examples

```sql
-- Materialize top products by date
CREATE TABLE popular_products (
  review_date date UNIQUE,
  agg_data jsonb
);

INSERT INTO popular_products
SELECT review_date, topn_add_agg(product_id)
FROM customer_reviews GROUP BY review_date;

-- Get the top-1 product for each day
SELECT review_date, (topn(agg_data, 1)).*
FROM popular_products ORDER BY review_date;

-- Top 10 across a time range (merging daily summaries)
SELECT (topn(topn_union_agg(agg_data), 10)).*
FROM popular_products
WHERE review_date >= '2000-01-01' AND review_date < '2000-02-01'
ORDER BY 2 DESC;
```
