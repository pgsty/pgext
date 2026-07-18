## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/marcmunro/pgbitmap/blob/58c425172ef564ab032c4f3e73da19a01759be8b/README.md)
- [Version 0.9.5 source distribution](https://api.pgxn.org/src/pgbitmap/pgbitmap-0.9.5/)
- [Generated API documentation](https://marcmunro.github.io/pgbitmap/docs/html/index.html)

`pgbitmap` defines a compact, non-sparse `bitmap` type for sets of nonnegative integer positions. It supports bit add/remove/test, minimum and maximum, union, intersection, difference, array conversion, iteration, comparisons, and aggregates.

```sql
CREATE EXTENSION pgbitmap;
SELECT to_bitmap(ARRAY[1, 2, 1000000]) AS members;
SELECT bitmap(1) + 2 + 3 AS members;
SELECT * FROM bits(to_bitmap(ARRAY[2, 5, 8]));
```

The type is designed to avoid storing all zero positions before a high bit. Its textual representation is compact and intended for round trips rather than human interpretation. Validate negative and extreme indexes, empty values, aggregates, ordering, and cross-architecture dump/restore.

Upstream calls 0.9.5 beta, while the catalog marks the project abandoned. It was developed partly for privilege sets, but a bitmap is not itself an authorization system: enforce row access and mutation privileges separately. Pin binary and text formats, test TOAST and memory behavior for wide or fragmented sets, and rebuild dependent data if representation changes. Prefer maintained core or application set representations when portability matters.
