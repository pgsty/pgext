## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/README.md)
- [Version 1.0 SQL objects](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/pgstrtranslate--1.0.sql)
- [C implementation](https://github.com/AbdulYadi/pgstrtranslate/blob/b6f77ac939f09a47f5ed797b3763437a06c9825c/pgstrtranslate.c)

`pgstrtranslate` extends PostgreSQL's character-wise `translate` operation with ordered multi-character replacements. Its function `pgstrtranslate(boolean, text, text[], text[])` returns text.

```sql
CREATE EXTENSION pgstrtranslate;
SELECT pgstrtranslate(
  false,
  'abcdefghijkl',
  ARRAY['ab', 'efg', '2cd'],
  ARRAY['012', '3', '78']
);
```

With the first argument false, all search terms are matched against the original input, so replacements do not create new matches. With it true, replacements run sequentially and later search terms can match text produced by earlier replacements. Order therefore changes results.

Require search and replacement arrays to have the intended pairing and test empty elements, duplicates, overlaps, nulls, multibyte text, and expansion. The function is marked immutable, so any persisted index or generated value assumes stable byte-for-byte semantics. Version 1.0 has no current compatibility matrix; benchmark worst-case overlapping patterns and validate output before using it for canonicalization, security filtering, or identifiers.
