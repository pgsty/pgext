## Usage

Sources:

- [Official README at the catalog revision](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/README.md)
- [Extension SQL at the catalog revision](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws--1.0.sql)
- [Parser implementation at the catalog revision](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.c)

`pg_scws` 1.0 integrates the SCWS Chinese word segmenter as a PostgreSQL full-text-search parser. It installs the scws parser and the scwscfg text-search configuration, which maps several lexical classes through the simple dictionary.

### Core Workflow

```sql
CREATE EXTENSION pg_scws;

SELECT to_tsvector(
    'scwscfg',
    '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);

CREATE INDEX article_search_idx
ON article
USING gin (to_tsvector('scwscfg', body));

SELECT id
FROM article
WHERE to_tsvector('scwscfg', body)
      @@ to_tsquery('scwscfg', '计算 & 专家');
```

Use the same text-search configuration in the indexed expression and every matching query.

### Objects and Settings

- The extension creates the scws text-search parser and scwscfg configuration.
- Settings control dictionary-in-memory mode, charset selection, rules file, extra dictionaries, punctuation handling, duality segmentation, and multi-mode segmentation.
- Dictionary and rule files are resolved from PostgreSQL's text-search data directory. Every server that can execute queries needs identical files and configuration.

```sql
SHOW scws.dict_in_memory;
SHOW scws.charset;
SHOW scws.rules;
SHOW scws.extra_dicts;
SHOW scws.punctuation_ignore;
SHOW scws.seg_with_duality;
SHOW scws.multi_mode;
```

### Compatibility and Operations

- Upstream reports testing only on PostgreSQL 9.4 and merely expects later 9.x versions to work. The parser uses server APIs that can change between major versions; build and test it on each exact target major.
- The source bundles SCWS, a default UTF-8 dictionary, and rules. Custom dictionaries should be deployed atomically across primary, replicas, and failover candidates before rebuilding affected indexes.
- Segmentation or dictionary changes alter lexemes. REINDEX expression indexes after intentional rule changes and validate search recall/precision with representative Chinese text.
