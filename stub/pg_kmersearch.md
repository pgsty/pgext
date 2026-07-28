## Usage

Sources:

- [Official upstream README](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/README.md)
- [Official extension control file (pg_kmersearch.control)](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/pg_kmersearch.control)
- [Official extension SQL (pg_kmersearch--1.0.sql)](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/pg_kmersearch--1.0.sql)

`pg_kmersearch` — A PostgreSQL extension for DNA sequence similarity search using k-mer indexing. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_kmersearch;

-- Create a table with DNA sequences
CREATE TABLE sequences (id SERIAL, name TEXT, dna DNA4);

-- Create GIN index (default: int4 keys for k-mer size 16)
CREATE INDEX ON sequences USING gin(dna kmersearch_dna4_gin_ops_int4);

-- Insert sequences
INSERT INTO sequences (name, dna) VALUES ('seq1', 'ATCGATCG');

-- Search using k-mer similarity
SELECT * FROM sequences WHERE dna =% 'ATCGATCG';

-- Calculate similarity scores
SELECT name, kmersearch_matchscore(dna, 'ATCGATCG') AS score
FROM sequences ORDER BY score DESC;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bit_length(DNA2)` is an extension function and returns `integer`.
- `bit_length(DNA4)` is an extension function and returns `integer`.
- `char_length(DNA2)` is an extension function and returns `integer`.
- `char_length(DNA4)` is an extension function and returns `integer`.
- `kmersearch_actual_min_score_cache_free()` is an extension function and returns `integer`.
- `kmersearch_actual_min_score_cache_stats()` is an extension function and returns `TABLE`.
- `kmersearch_consistent_int2(internal, int2, text, int4, internal, internal)` is an extension function and returns `boolean`.
- `kmersearch_consistent_int4(internal, int2, text, int4, internal, internal)` is an extension function and returns `boolean`.
- `kmersearch_consistent_int8(internal, int2, text, int4, internal, internal)` is an extension function and returns `boolean`.
- `kmersearch_delete_tempfiles()` is an extension function and returns `TABLE`.
- `kmersearch_dna2_cmp(DNA2, DNA2)` is an extension function and returns `integer`.
- `kmersearch_dna2_eq(DNA2, DNA2)` is an extension function and returns `boolean`.
- `kmersearch_dna2_ge(DNA2, DNA2)` is an extension function and returns `boolean`.
- `kmersearch_dna2_gt(DNA2, DNA2)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
