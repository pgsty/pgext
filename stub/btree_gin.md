

## Usage

> [btree_gin: B-tree equivalent GIN operator classes](https://www.postgresql.org/docs/current/btree-gin.html)

Provides GIN index operator classes for data types that normally only support B-tree indexing. Useful for multicolumn GIN indexes that combine GIN-indexable and B-tree-indexable columns.

```sql
CREATE EXTENSION btree_gin;
```

### Supported Data Types

`int2`, `int4`, `int8`, `float4`, `float8`, `numeric`, `timestamp with time zone`, `timestamp without time zone`, `time with time zone`, `time without time zone`, `date`, `interval`, `oid`, `money`, `char`, `varchar`, `text`, `bytea`, `macaddr`, `macaddr8`, `inet`, `cidr`, `uuid`, `bit`, `varbit`, `bool`, `name`, `bpchar`, and all `enum` types.

### Examples

```sql
-- GIN index on an integer column
CREATE INDEX idx ON test USING GIN (a);
SELECT * FROM test WHERE a < 10;

-- Multicolumn GIN index combining full-text search with a scalar filter
CREATE INDEX idx ON articles USING GIN (body_tsvector, category);
SELECT * FROM articles
WHERE body_tsvector @@ to_tsquery('PostgreSQL')
  AND category = 'tech';
```

Note: btree_gin does not outperform standard B-tree indexes for single-column queries. Its main benefit is combining scalar columns with GIN-native columns (like tsvector or arrays) in a single multicolumn index.
