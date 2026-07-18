## Usage

Sources:

- [Upstream README](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/README.md)
- [Extension SQL](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/sql/kor_search--1.0.0.sql)
- [C implementation](https://github.com/junhkang/kor_search/blob/f10c6ff5f474b6cb455de56f6fb5b5af0dbf6c96/src/kor_search.c)
- [PGXN documentation](https://pgxn.org/dist/kor_search/1.0.0/doc/kor_search.html)

`kor_search` performs Korean/English word matching with database-resident keyword and synonym tables; it does not call an external translator or morphological analyzer. Install `pg_trgm` first because the version `1.0.0` SQL creates a `gin_trgm_ops` index but does not declare that dependency in its control file:

```sql
CREATE EXTENSION pg_trgm;
CREATE EXTENSION kor_search;

SELECT kor_search_like('이것은 엘지 제품입니다', 'lg');
SELECT kor_search_similar('I eat rice', '밥 먹다');
SELECT kor_search_tsvector('데이터 과학은 중요합니다', 'data science');
SELECT kor_search_regex('자바와 파이썬', 'java|python');
```

### API and security notes

The README examples use stale names `kor_like` and `kor_regex_search`; the published install SQL actually defines `kor_search_like` and `kor_search_regex`, as used above. Search quality and cost depend on the contents and size of `kor_search_word_transform` and `kor_search_word_synonyms`. The C implementation interpolates caller-provided text into internal SPI statements without SQL escaping, so do not expose these functions to untrusted input; revoke `EXECUTE` or patch the implementation before such use.
