## Usage

Sources:

- [Official source](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/README.md)
- [Official source](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/postgres/albanian_stemmer.control)
- [Official source](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/LICENSE)
- [Official source](https://github.com/edisnord/albanian-stemmer/blob/1.2.1/postgres/albanian_stemmer--1.0.sql)

albanian_stemmer supplies an Albanian Snowball dictionary and text-search configuration. Distribution 1.2.1 retains extension SQL version 1.0. Upstream provides PostgreSQL 14–18 support; install the matching native library before database creation. The control is trusted and relocatable and requires no preload.

### Text search

```sql
CREATE EXTENSION albanian_stemmer;
SELECT to_tsvector('albanian', 'Sëmundjet e mollës dhe të pemëve');
SELECT plainto_tsquery('albanian', 'mollët');
SELECT ts_lexize('albanian_stem', 'mollëve');
```

The albanian configuration copies simple parsing and maps word token classes to albanian_stem; numbers, URLs and other classes retain simple behavior. The dictionary folds ë/ç, removes stopwords and strips inflection. For a custom dictionary, albanian_stem_template accepts StopWords = false.

### Search-data maintenance

Use the same configuration when building document vectors and query terms. Stemming is a lossy normalization, not a full lemmatizer or a semantic-equivalence guarantee. Rule changes can change stored lexemes: recompute materialized vectors and rebuild affected indexes when adopting changed rules. SQL version 1.0 alone does not identify the distribution's stemmer revision.
