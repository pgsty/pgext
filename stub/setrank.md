## Usage

Sources:

- [Official README](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/README)
- [Extension control file](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank.control)
- [Version 1.0 extension SQL](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank--1.0.sql)
- [IDF table loader and cache](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/idf.c)

`setrank` 1.0 adds TF-IDF, cover-density TF-IDF, and score functions for PostgreSQL full-text search. Unlike built-in ranking, it reads document-frequency statistics from a caller-maintained table so common lexemes can receive less weight.

### Build the Statistics Table

The table must have exactly two columns: lexeme `text` followed by document count `int4` or `int8`. A row whose first column is NULL stores the total document count.

```sql
CREATE EXTENSION setrank;

CREATE TABLE docs (
  id bigint GENERATED ALWAYS AS IDENTITY,
  body text,
  fts tsvector
);

INSERT INTO docs(body, fts)
VALUES ('x ray star', to_tsvector('english', 'x ray star'));

CREATE TABLE docs_stat AS
SELECT word AS value, ndoc::bigint AS ndoc
FROM ts_stat('SELECT fts FROM docs');

INSERT INTO docs_stat(value, ndoc)
VALUES (NULL, (SELECT count(*) FROM docs));
```

The second column must be positive for every row, lexemes must be unique, and the total must be at least every individual document frequency.

### Rank Documents

```sql
SELECT id,
       ts_rank_tfidf('docs_stat', fts, plainto_tsquery('english', 'x ray star')) AS rank
FROM docs
WHERE fts @@ plainto_tsquery('english', 'x ray star')
ORDER BY rank DESC;
```

`get_idf(table_name, lexeme)` returns the cached inverse document frequency. The three ranking families are `ts_rank_tfidf`, `ts_rank_cd_tfidf`, and `ts_score_tfidf`; each has overloads with optional `float4[]` weights and normalization flags.

### Cache and Maintenance Caveats

On first use in a backend, the C code reads the entire statistics table into `TopMemoryContext`. It never invalidates or refreshes that per-backend cache, so later table changes are invisible until the session reconnects. Long-lived connection pools can therefore serve stale rankings.

The functions are declared `IMMUTABLE` even though their results depend on table contents. Do not use them in expression indexes, generated columns, or other places that require true immutability. Rebuild the statistics table after corpus changes, then recycle all sessions that called `setrank` before trusting new scores.

The loader quotes the whole `table_name` argument as one identifier, so the reviewed implementation expects a simple unqualified table name rather than `schema.table`. Restrict who can call the functions with arbitrary table names, and keep the statistics table in the active search path. The project is abandoned; validate behavior against the target PostgreSQL version before retaining it.
