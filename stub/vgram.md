## Usage

Sources:

- [Official README](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/README.md)
- [Official extension SQL](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram--1.0.sql)
- [Official GIN implementation](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram_gin.c)

`vgram` accelerates `LIKE` and `ILIKE` substring searches with GIN indexes built from variable-length q-grams. Unlike fixed trigrams, it learns frequent short grams from the target corpus and indexes longer or rarer grams, which can improve selectivity for a particular data distribution.

### Build a Corpus-Specific Index

Version `1.0` requires PostgreSQL 13 or later. In psql, collect frequent grams, save the result, and pass it to the `vgram_gin_ops` options:

```sql
CREATE EXTENSION vgram;

SELECT qgram_stat(title, 2, 4, 0.05) AS vgrams
FROM documents
\gset

CREATE INDEX documents_title_vgram_idx
ON documents USING gin (
  title vgram_gin_ops (minq=2, maxq=4, vgrams=:'vgrams')
);

SELECT *
FROM documents
WHERE title ILIKE '%indexing%';
```

`minq` and `maxq` bound extracted gram lengths. The threshold passed to `qgram_stat()` marks grams at or above that corpus frequency as too common to index. Use `get_vgrams()` to inspect extraction for a candidate value and learned gram array.

### Objects and Advanced Statistics

- `qgram_stat(text, int, int, float8)` aggregates frequent q-grams.
- `get_vgrams(text, int, int, text[])` shows the grams selected for text.
- `vgram_gin_ops` indexes ordinary `text` for `LIKE` and `ILIKE`.
- `vgram_text` is text-compatible but installs custom frequency statistics and selectivity estimation.
- `vgram_gin_ops2` is the GIN operator class for `vgram_text`.

### Maintenance Boundaries

The learned option array is embedded in the index definition and can be very large. It reflects the sample at build time; major language or distribution changes require new statistics, comparison of recall/plans, and usually `REINDEX` or a replacement index. GIN candidates are rechecked, but overly common or too-short patterns may still fall back to broad scans. Benchmark build time, index size, write amplification, pending-list behavior, VACUUM, collation, multibyte text, case folding, and planner estimates. Keep the training query and parameters reproducible so replicas and rebuilt environments use the same operator-class options.
