## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/README.md)
- [Version 1.0 SQL declaration](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch--1.0.sql)
- [Matching implementation](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.c)
- [Extension control file](https://github.com/lemmerelassal/NoraSearch/blob/0c11af31b5aea1c58ca401e57966870b13b19ddf/norasearch.control)

`norasearch` provides one byte-oriented similarity function. At every nonnegative alignment of a needle against a haystack, it counts equal byte positions and returns the alignments whose count is strictly greater than a threshold. This is cross-correlation-like scoring, not ordinary contiguous substring search.

### Core Workflow

```sql
CREATE EXTENSION norasearch;

SELECT norasearch('abracadabra', 'abra', 2);
-- {{0,4},{7,4}}
```

The signature is `norasearch(hay text, needle text, minmatch int) RETURNS int[][]`. Each result row is `{offset,count}`. `offset` is zero-based and only offsets from the beginning of the haystack onward are considered. A negative `minmatch` is treated as zero; matching uses the strict condition `count > minmatch`.

### Result and Performance Boundaries

The function is strict for SQL NULL input. Empty strings, no qualifying alignment, or other zero-result cases return `NULL`, not an empty array. The implementation compares raw bytes, so offsets are byte positions and multibyte UTF-8 characters can be split; do not interpret them as character indexes.

There is no index operator class or planner integration. The implementation examines each haystack position against the needle, giving worst-case work proportional to both lengths and allocating result arrays per call. Benchmark large or repetitive values before applying `norasearch` across many rows.
