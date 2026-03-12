


## Usage

> [hll: type for storing hyperloglog data](https://github.com/citusdata/postgresql-hll)

The `hll` extension provides a HyperLogLog data type for probabilistic distinct-value counting. It enables efficient approximate `COUNT(DISTINCT)` operations with configurable accuracy, and supports set union operations that allow pre-aggregated data to be combined without loss of precision.

### Data Types

- **`hll`** -- HyperLogLog accumulator with parameters: `hll(log2m, regwidth, expthresh, sparseon)`
- **`hll_hashval`** -- Hashed value type for insertion into HLL structures

### Core Functions

| Function | Description |
|----------|-------------|
| `hll_empty()` | Create an empty HLL |
| `hll_add(hll, hll_hashval)` | Add a hashed value to an HLL |
| `hll_cardinality(hll)` | Estimate distinct count |
| `hll_union(hll, hll)` | Combine two HLLs |
| `hll_add_agg(hll_hashval)` | Aggregate hashed values into a single HLL |
| `hll_union_agg(hll)` | Merge multiple HLLs into one |
| `hll_print(hll)` | Display HLL parameters and contents |

### Hash Functions

| Function | Input Type |
|----------|-----------|
| `hll_hash_boolean(boolean [, seed])` | boolean |
| `hll_hash_smallint(smallint [, seed])` | smallint |
| `hll_hash_integer(integer [, seed])` | integer |
| `hll_hash_bigint(bigint [, seed])` | bigint |
| `hll_hash_bytea(bytea [, seed])` | bytea |
| `hll_hash_text(text [, seed])` | text |
| `hll_hash_any(any [, seed])` | any (dynamic dispatch, slower) |

### Operators

| Operator | Function | Example |
|----------|----------|---------|
| `\|\|` | `hll_add` / `hll_union` | `users \|\| hll_hash_integer(123)` |
| `#` | `hll_cardinality` | `#users` |

### Example: Daily Unique User Tracking

```sql
-- Store daily unique user counts
CREATE TABLE daily_uniques (
    date  date UNIQUE,
    users hll
);

-- Aggregate daily data
INSERT INTO daily_uniques(date, users)
    SELECT date, hll_add_agg(hll_hash_integer(user_id))
    FROM facts GROUP BY 1;

-- Weekly uniques (unions are lossless)
SELECT hll_cardinality(hll_union_agg(users))
FROM daily_uniques
WHERE date >= '2012-01-02' AND date <= '2012-01-08';

-- Monthly breakdown
SELECT EXTRACT(MONTH FROM date) AS month,
       #hll_union_agg(users) AS approx_uniques
FROM daily_uniques
WHERE date >= '2012-01-01' AND date < '2013-01-01'
GROUP BY 1;

-- 7-day sliding window
SELECT date, #hll_union_agg(users) OVER seven_days
FROM daily_uniques
WINDOW seven_days AS (ORDER BY date ASC ROWS 6 PRECEDING);
```

### Configuration Parameters

- **`log2m`** (4--31): Number of registers as log-base-2. Controls accuracy with relative error of +/-1.04/sqrt(2^log2m). Default: 11.
- **`regwidth`** (1--8): Bits per register. Tuned alongside log2m for maximum cardinality estimation. Default: 5.
- **`expthresh`** (-1 to 18): Controls EXPLICIT-to-SPARSE promotion. `-1` for auto mode, `0` to skip EXPLICIT. Default: -1.
- **`sparseon`** (0 or 1): Enables/disables SPARSE representation. Default: 1.

All inputs to a given HLL must use the same hash seed. HLLs intended for union operations must have been populated with identically-seeded hash values.
