## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/README.md)
- [Compiled extension SQL](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/tsp_semantic_headline--1.0.sql)
- [Extension control file](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/tsp_semantic_headline.control)

`tsp_semantic_headline` is a pure-SQL experiment for phrase-aware full-text highlighting. It supplies `TS_SEMANTIC_HEADLINE`, preprocessed query/vector helpers, and `TS_FAST_HEADLINE`, which can use aligned precomputed `TSVECTOR` and `text[]` columns for faster recall.

```sql
CREATE EXTENSION unaccent;
CREATE EXTENSION tsp_semantic_headline;
SELECT TS_SEMANTIC_HEADLINE(
  'english',
  'phrase matches are highlighted as complete phrases',
  TO_TSPQUERY('english', 'phrase<->match')
);
```

Use the same preprocessing for indexed text, recall arrays, and queries; otherwise token positions diverge and highlights are wrong. The fast path trades extra stored columns and update work for lower runtime cost. Reproduce performance claims with your language, corpus, query shapes, and options.

Upstream calls the project in development and documents unimplemented options. Its special-character treatment inserts separators and explicitly warns that space-based assumptions can change or destroy meaning in many languages. Treat highlighted output as presentation, escape it for the destination UI, and never mark untrusted HTML as safe. Test stemming, negation, phrase distance, punctuation, Unicode, XSS, long documents, and divergence after content updates.
