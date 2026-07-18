## Usage

Sources:

- [Upstream README](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/README.md)
- [Extension SQL](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/pg_hspell--1.0.sql)
- [Extension control file](https://github.com/IgKh/pg_hspell/blob/89de5bffea99a918bc285096e70fcc4c5f102b34/pg_hspell.control)
- [PostgreSQL text-search configuration documentation](https://www.postgresql.org/docs/current/textsearch-configuration.html)

`pg_hspell` provides an `hspell` text-search dictionary template that lemmatizes Hebrew words. It requires GNU/Linux and an external `libhspell` built with `--enable-linginfo`.

The extension creates a dictionary but no complete text-search configuration, so map it into a configuration explicitly:

```sql
CREATE EXTENSION pg_hspell;

CREATE TEXT SEARCH CONFIGURATION hebrew (COPY = simple);
ALTER TEXT SEARCH CONFIGURATION hebrew
  ALTER MAPPING FOR word, hword, hword_part
  WITH hspell, simple;

SELECT to_tsvector('hebrew', 'הרכבת');
```

### Linguistic limits

Version `1.0` documents PostgreSQL 9.6 and later. Recognized words emit every possible lemma without contextual disambiguation, improving recall at the cost of precision. Tokens containing Niqqud are not recognized by `libhspell`, and PostgreSQL's default parser can split Hebrew acronyms or abbreviations around quotes. Strip Niqqud before this dictionary and test parser mappings with representative text.
