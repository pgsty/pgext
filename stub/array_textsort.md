## Usage

Sources:

- [Official README](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/README.md)
- [Version 1.1 extension SQL](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort--1.1.sql)
- [C implementation](https://github.com/2ndQuadrant/array_textsort/blob/dd60ea1a5403e1ff938aeb6d26d0e9a901ea30f4/array_textsort.c)

`array_textsort` adds sorting and deduplication functions for one-dimensional `text[]` values. It is useful when an application needs a canonical array order or a distinct set represented as an array rather than rows.

### Core Workflow

```sql
CREATE EXTENSION array_textsort;

SELECT array_textsort(ARRAY['pear', 'apple', 'banana']);
SELECT array_distinct(ARRAY['pear', 'apple', 'pear', 'banana']);
```

Both functions sort according to the database locale. `array_distinct` then removes adjacent duplicates from that sorted result.

### Object Index

- `array_textsort(text[]) RETURNS text[]` sorts a one-dimensional text array.
- `array_distinct(text[]) RETURNS text[]` sorts the array and removes duplicate elements.

### Operational Notes

Version `1.1` is relocatable and declares no extension dependency or preload requirement. The C functions are `STRICT` and `IMMUTABLE`; SQL `NULL` input therefore returns SQL `NULL`.

The implementation rejects arrays containing null elements. Use it only with one-dimensional arrays: that is the documented interface, and multidimensional behavior is not defined as a supported contract. Locale determines comparison order, so the same strings can sort differently in databases with different locale settings. If a stable cross-database order is required, normalize values outside these functions or use a row-based query with an explicit collation.
