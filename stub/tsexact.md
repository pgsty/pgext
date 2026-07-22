## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/tsexact.control)
- [Official upstream documentation](https://github.com/postgrespro/tsexact/blob/b08a6ce7ed40b5b62cccdb12d7e138c653d2efd0/README.md)

`tsexact` 1.0 adds exact-fragment helpers for PostgreSQL full-text vectors. It was designed to emulate phrase search on PostgreSQL 9.5 and earlier; upstream recommends built-in phrase-search operators on PostgreSQL 9.6 and later.

### Core Workflow

```sql
CREATE EXTENSION tsexact;

SELECT ts_exact_match(
  'cat:3 fat:2 sad:5'::tsvector,
  'cat:2 fat:1 sad:4'::tsvector
);

SELECT ts_squeeze('cat:2,9 fat:1,6 sad:4'::tsvector);
```

`ts_exact_match(tsvector, tsvector)` checks whether the fragment's relative positions occur at some offset in the document and ignores weights. Its three-argument overload accepts a weight mask for the document. `ts_squeeze(tsvector)` removes gaps in position numbering, `setweight(tsquery, text)` applies weights to query lexemes, and `poslen(tsvector)` returns the total position span.

### Indexing Pattern and Caveats

`ts_exact_match` is not directly supported by full-text indexes. Use an indexed `@@` condition to obtain candidates, then apply `ts_exact_match` as an exact recheck.

```sql
SELECT *
FROM documents
WHERE search_vector @@ plainto_tsquery('fat rat')
  AND ts_exact_match(
        search_vector,
        ts_squeeze(to_tsvector('fat rat'))
      );
```

Tokenization, dictionaries, stop words, weights, and position gaps all affect results. Prefer native phrase queries on current PostgreSQL, and retain `tsexact` only when its legacy vector semantics are specifically required and covered by regression tests.
