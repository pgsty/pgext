## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/README.md)
- [Version 1.0 SQL declaration](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch--1.0.sql)
- [Extension control file](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.control)

`norasearch` provides one C function that compares two strings at each possible offset and returns the offsets whose aligned byte-match count exceeds a threshold. Results are a two-dimensional integer array of `{offset,count}` pairs.

```sql
CREATE EXTENSION norasearch;
SELECT norasearch('abracadabra', 'abra', 2);
-- {{0,4},{7,4}}
```

The signature is `norasearch(hay text, needle text, minmatch int)`. Offsets are zero-based, and only rows with `count > minmatch` are emitted. The implementation compares bytes rather than locale-aware characters, so verify behavior with multibyte text before using it for user-facing search.

The current README gives Windows/MSYS2 and PostgreSQL 18 build instructions. There is no index operator class or planner integration: every call performs function-level matching, so benchmark it on representative string sizes and row counts.
