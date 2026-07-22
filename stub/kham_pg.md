## Usage

Sources:

- [Official kham-pg README](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/README.md)
- [Version 0.8.2 extension SQL](https://github.com/preedep/kham/blob/ec2394cc70ab9687b6e0bcd7c02c2f75a4cf9fbd/kham-pg/sql/kham_pg--0.8.2.sql)
- [PGXN 0.8.2 documentation](https://pgxn.org/dist/kham_pg/0.8.2/README.html)

`kham_pg` provides Thai full-text search for PostgreSQL. Its `kham` parser segments text without relying on spaces, while its dictionaries add normalized words, phonetic codes, RTGS romanization, number normalization, part-of-speech lexemes, and named-entity handling to standard `tsvector` and `tsquery` workflows.

### Core Workflow

```sql
CREATE EXTENSION kham_pg;

SELECT *
FROM ts_parse('kham', 'ทักษิณเดินทางไปกรุงเทพ');

SELECT kham_tsvector('กินข้าวกับปลา');
SELECT kham_tsquery('ปลา');

CREATE TABLE articles (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text,
    body text NOT NULL
);

CREATE INDEX articles_kham_fts_idx
ON articles USING gin (kham_tsvector(body));

SELECT id, title
FROM articles
WHERE kham_tsvector(body) @@ kham_tsquery('ปลา');
```

`kham_tsvector(document)` is a convenience wrapper around `to_tsvector('kham', document)`, and `kham_tsquery(query)` wraps `plainto_tsquery('kham', query)`. Standard highlighting also works:

```sql
SELECT ts_headline('kham', body, kham_tsquery('ปลา'))
FROM articles;
```

### Text Search Objects

- Parser and ready-to-use configuration: `kham`.
- Default dictionary: `kham_fts_dict`, using `lk82` Thai Soundex.
- Alternative dictionaries: `kham_fts_dict_udom83` and `kham_fts_dict_metasound`.
- Simple lowercasing pass-through dictionary: `kham_dict`.
- Token types include `thai`, `latin`, `number`, `punct`, `emoji`, `unknown`, and `named`.

Thai and named-entity tokens can expand to several colocated lexemes, allowing phonetic or Latin-script queries to match Thai text. Thai stopwords are suppressed, Thai digits are normalized alongside ASCII forms, and punctuation and emoji are discarded by the default configuration.

### Caveats

Version `0.8.2` is non-relocatable and targets PostgreSQL `14` through `18` in the reviewed upstream release. Segmentation, phonetic expansion, stopword removal, and named-entity recognition affect recall and index size, so validate them with domain-specific Thai text. A GIN expression index matches queries only when they use the same expression and configuration. No preload or restart is required after the library is installed.
