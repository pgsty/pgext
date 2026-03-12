

## Usage

> [pg_textsearch: Modern ranked text search for PostgreSQL with BM25](https://github.com/timescale/pg_textsearch)

Modern ranked text search using BM25 scoring with Block-Max WAND optimization. Simple syntax, fast top-k queries, parallel index builds, and partitioned table support.

Add to `shared_preload_libraries`:

```
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
```

### Quick Start

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
    ('PostgreSQL is a powerful database system'),
    ('BM25 is an effective ranking function'),
    ('Full text search with custom scoring');

-- Create a BM25 index
CREATE INDEX docs_idx ON documents USING bm25(content) WITH (text_config='english');

-- Query using the <@> operator (returns negative BM25 score, lower = better match)
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

### Querying

```sql
-- Auto-detect index from column
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;

-- Explicit index specification
SELECT * FROM documents
WHERE content <@> to_bm25query('database system', 'docs_idx') < -1.0;
```

### Filtering

**Pre-filtering** reduces rows before scoring (best with selective filters):

```sql
CREATE INDEX ON documents (category_id);
SELECT * FROM documents
WHERE category_id = 123
ORDER BY content <@> 'search terms'
LIMIT 10;
```

**Post-filtering** applies BM25 scan first, then filters:

```sql
SELECT * FROM documents
WHERE content <@> to_bm25query('search terms', 'docs_idx') < -5.0
ORDER BY content <@> 'search terms'
LIMIT 10;
```

### Index Options

| Option | Default | Description |
|--------|---------|-------------|
| `text_config` | (required) | PostgreSQL text search configuration |
| `k1` | 1.2 | Term frequency saturation parameter |
| `b` | 0.75 | Length normalization parameter |

```sql
CREATE INDEX ON documents USING bm25(content)
  WITH (text_config='english', k1=1.5, b=0.8);

-- Language-specific configurations
CREATE INDEX ON french_docs USING bm25(content) WITH (text_config='french');
CREATE INDEX ON german_docs USING bm25(content) WITH (text_config='german');
```

### Data Types

**bm25query** — represents queries for BM25 scoring:

```sql
SELECT to_bm25query('search query text', 'docs_idx');
-- docs_idx:search query text
```
