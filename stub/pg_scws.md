## Usage

Sources:

- [pg_scws README at the reviewed commit](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/README.md)
- [pg_scws.control at the reviewed commit](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.control)
- [Version 1.0 install SQL](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws--1.0.sql)
- [Parser and configuration implementation](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.c)
- [Bundled SCWS sources and dictionaries](https://github.com/jaiminpan/pg_scws/tree/338d1ebe911372165acad331264bbf48afa56a9e/libscws)

`pg_scws` embeds the SCWS Chinese word segmenter as PostgreSQL's `scws` text-search parser. Installation creates the `scwscfg` text-search configuration and maps noun, verb, adjective, idiom, exclamation, and temporary-word token classes through the `simple` dictionary.

### Chinese Full-Text Parsing

```sql
CREATE EXTENSION pg_scws;

SELECT to_tsvector(
  'scwscfg',
  '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);
```

Session settings include `scws.charset`, `scws.rules`, `scws.extra_dicts`, `scws.punctuation_ignore`, `scws.seg_with_duality`, and `scws.multi_mode`. Rule and dictionary names resolve under PostgreSQL's shared `tsearch_data` directory.

### Caveats

- Version `1.0` was tested upstream only with PostgreSQL 9.4, and the reviewed source dates from 2016. Port and regression-test the parser against the exact target release.
- The bundled default dictionary and rules determine segmentation quality and vocabulary. Changes to those files are server-wide deployment changes and must be consistent across replicas.
- `scws.extra_dicts` accepts server-side dictionary filenames, not client files. Upstream's user-defined-dictionary documentation is empty; establish your own controlled build, distribution, and rollback procedure.
