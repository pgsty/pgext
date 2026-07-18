## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/README.md)
- [Extension control file](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/dict_id.control)
- [Dictionary implementation](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/dict_id.c)

`dict_id` is an alpha Indonesian stemming dictionary for PostgreSQL full-text search, backed by the cSastrawi stemmer. Upstream explicitly says it is intended only for people helping with development and documents PostgreSQL 9.4.4.

```sql
CREATE EXTENSION dict_id;
SELECT ts_lexize('dict_id', 'bertapa');
```

The example returns the stem `tapa`. To use the dictionary in search, create a dedicated text-search configuration and map selected token classes only after testing a representative Indonesian corpus; stemming choices directly affect indexed lexemes and query matches.

Treat this as historical experimental code. There is no current linguistic quality, thread-safety, memory-safety, or PostgreSQL compatibility statement in the reviewed material. Compare output with a maintained Indonesian search pipeline, measure false positives and false negatives, and rebuild every dependent text-search index if dictionary behavior changes. Do not silently replace a production dictionary because old and new index entries can become semantically inconsistent.
