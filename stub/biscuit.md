
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION biscuit;
> CREATE INDEX idx_users_name ON users USING biscuit(name);
> SELECT * FROM users WHERE name LIKE '%john%';
> ```
>
> Sources: [README](https://github.com/CrystallineCore/Biscuit), [Docs](https://biscuit.readthedocs.io/)

`biscuit` is a PostgreSQL index access method for fast `LIKE` and `ILIKE` pattern matching, including multi-column searches. The upstream project positions it as a deterministic bitmap index that avoids the false-positive recheck overhead common in trigram-based searches.

### Quick Start

Create the extension and build a Biscuit index on one or more text columns:

```sql
CREATE EXTENSION biscuit;

CREATE INDEX idx_users_name ON users USING biscuit(name);

CREATE INDEX idx_products_search
ON products USING biscuit(name, description, category);
```

Basic wildcard queries work with the index:

```sql
SELECT * FROM users WHERE name LIKE '%john%';
SELECT * FROM users WHERE name NOT LIKE 'a%b%c';
SELECT COUNT(*) FROM users WHERE name LIKE '%test%';

SELECT *
FROM products
WHERE name LIKE '%widget%'
  AND description LIKE '%blue%'
  AND category LIKE 'electronics%'
LIMIT 10;
```

## Index Behavior

Biscuit stores bitmap position indexes for each string and can match both forward and backward character positions. The upstream design highlights:

- positive indexes for characters at exact positions
- negative indexes for characters counted from the string end
- case-insensitive variants for `ILIKE`
- exact-length and minimum-length bitmaps for fast length filtering

For a pattern such as `LIKE 'abc%def'`, Biscuit can intersect prefix and suffix bitmaps plus a minimum-length filter, producing exact matches without a heap recheck phase.

### Pattern Cases

The implementation documents optimized paths for common pattern types:

- exact matches such as `'abc'`
- prefix patterns such as `'abc%'`
- suffix patterns such as `'%xyz'`
- substring patterns such as `'%abc%'`
- multi-column predicates, where Biscuit reorders predicates by estimated selectivity

## Performance Notes

The upstream README emphasizes bitmap-only evaluation and several execution optimizations, including:

- early termination when an intermediate bitmap becomes empty
- direct use of roaring bitmaps for sparse and dense cases
- negative-position lookups for suffix predicates
- sorted TID output to improve heap access locality
- special handling for aggregate queries and `LIMIT`

The project README also includes a benchmark setup comparing Biscuit indexes with trigram-based approaches on a 1M-row table.

## Requirements

The current upstream README lists these requirements for source builds:

- PostgreSQL 16 or newer
- standard build tools such as `gcc`, `make`, and `pg_config`
- optional CRoaring for improved performance

The project publishes packages on [PGXN](https://pgxn.org/dist/biscuit/) and maintains a dedicated documentation site on Read the Docs.
