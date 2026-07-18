## Usage

Sources:

- [Upstream README](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/README.md)
- [Version 1.0 install SQL](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/jieba_parser--1.0.sql)
- [C++ and embedded Python implementation](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/jieba_parser.cc)
- [Python 2.7 build configuration](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/Makefile)

jieba_parser 1.0 is a 2014 full-text-search parser that embeds Python and calls the Python Jieba package for Chinese tokenization and keyword extraction.

### Parser configuration

Only evaluate a repaired, dependency-pinned build in a disposable backend:

```sql
CREATE EXTENSION jieba_parser;
CREATE TEXT SEARCH CONFIGURATION jiebacfg (PARSER = jiebaparser);
ALTER TEXT SEARCH CONFIGURATION jiebacfg
  ADD MAPPING FOR word WITH simple;
SELECT to_tsvector('jiebacfg', '我来到北京清华大学');
SELECT jieba_extract_tags('清华大学位于北京', 3);
```

The text search configuration is a separate database object; account for it when dropping or upgrading the extension.

### Caveats

- The Makefile is hard-coded to Python 2.7 and uses Python 2-only APIs. Python 2 is end-of-life, and this source does not build unchanged with supported Python 3 releases.
- The shared library initializes Python and imports Jieba during global object construction. Import and call failures are guarded by C assertions rather than PostgreSQL errors; release builds can continue with null pointers and crash a backend.
- Two global objects both own interpreter lifecycle calls. Embedding or finalizing the same interpreter alongside other Python-using extensions is not safely coordinated.
- jieba_rank accepts a vector and query but always returns zero. It is not a usable ranking function.
- Tokenization and extracted tags depend on the separately installed Jieba package, dictionary, and model version. Pin those assets for reproducible indexes and rebuild affected text-search indexes after intentional changes.
- Upstream has no license file or current PostgreSQL compatibility evidence. Prefer a maintained parser rather than deploying this implementation as published.
