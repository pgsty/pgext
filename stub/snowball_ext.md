## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/README.md)
- [Version 1.0 SQL objects](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/snowball_ext--1.0.sql)
- [Nepali stemmer source](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/libstemmer/stem_UTF_8_nepali.c)
- [Extension control file](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/snowball_ext.control)

`snowball_ext` copies PostgreSQL's non-extensible Snowball dictionary template so an additional stemming algorithm can be registered. Version 1.0 supplies a Nepali stemmer and a `nepali` text-search configuration.

```sql
CREATE EXTENSION snowball_ext;
SELECT to_tsvector('nepali', 'अँगअँकाउछन्');
```

Use the configuration in the same way as built-in text-search configurations, and test stemming against a reviewed Nepali corpus before building production indexes. Tokenization and stemming quality directly affect recall, precision, and index contents.

The catalog marks this extension abandoned. It duplicates PostgreSQL internal Snowball template code and has no current compatibility matrix, so newer backend changes may break its native build. Because text-search index output depends on the algorithm, do not silently change versions on replicas or during upgrades; rebuild affected indexes only after comparing results and approving the linguistic change.
