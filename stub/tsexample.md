## Usage

Sources:

- [Extension control file](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample.control)
- [Version 1.0 installation SQL](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample--1.0.sql)
- [C parser and dictionary implementation](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample.c)
- [Regression example SQL](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/sql/tsexample.sql)

`tsexample` 1.0 is a compact C example of PostgreSQL's extensible full-text-search interfaces. It installs the `sample_parser` parser, the `cutdict` dictionary template, a `cut3` dictionary, and the `sample` text-search configuration.

### Inspect parsing and lexemes

```sql
CREATE EXTENSION tsexample;

SELECT * FROM ts_parse('sample_parser', 'abc def 123 1xx yy3');
SELECT to_tsvector('sample', 'abcdef 12345678 xyz');
SELECT plainto_tsquery('sample', 'abcdef 12345678 xyz');
```

`sample_parser` groups letters, digits, and underscore characters into tokens, classifying an all-digit token as `number` and any token containing a non-digit as `word`. The `sample` configuration maps numbers through PostgreSQL's `simple` dictionary and words through `cut3`. `cut3` lowercases a word and, for longer inputs, emits lexemes made from its first three and last three characters.

### Example status

This project is instructional code, not a general linguistic analyzer. Its fixed object names and three-character dictionary settings are examples to adapt rather than a complete search design. The repository supplies no README, explicit license file, release history, or compatibility matrix. Build and test the C module against the exact PostgreSQL major version, and use distinct object names if incorporating the pattern into another extension.
