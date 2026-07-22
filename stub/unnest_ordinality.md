## Usage

Sources:

- [Official documentation at the reviewed commit](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/doc/unnest_ordinality.md)
- [SQL declaration](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/sql/unnest_ordinality.sql)
- [Set-returning implementation](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/src/unnest_ordinality.c)
- [Extension control file](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/unnest_ordinality.control)
- [Official PGXN distribution](https://pgxn.org/dist/unnest_ordinality/)

`unnest_ordinality` is a legacy set-returning function that expands any PostgreSQL array into `(element_number, element)` rows. It predates PostgreSQL's built-in `WITH ORDINALITY` syntax and is deprecated for maintained server versions.

### Core Workflow

```sql
CREATE EXTENSION unnest_ordinality;

SELECT *
FROM unnest_ordinality(ARRAY['a', NULL, 'c']);

-- Preferred on PostgreSQL 9.4 and later:
SELECT ordinality::integer AS element_number, element
FROM unnest(ARRAY['a', NULL, 'c']) WITH ORDINALITY AS u(element, ordinality);
```

The extension function returns `element_number integer` starting at 1 and `element anyelement`. It preserves NULL elements and flattens multidimensional arrays in storage order. The numbers are sequential ordinals, not the source array's declared subscripts; custom lower bounds and dimensional coordinates are not preserved. An empty array returns no rows.

### Deprecation and Performance Boundaries

The C function deconstructs the entire source array and materializes every output row in a tuplestore before returning. Large or toasted arrays therefore require memory and temporary-file capacity proportional to the expanded result. It uses old backend tuple and array APIs and was mainly tested with PostgreSQL 9.2.

Use built-in `unnest(...) WITH ORDINALITY` on PostgreSQL 9.4 or later. Keep `unnest_ordinality` only for a pinned legacy compatibility requirement, and migrate SQL before upgrading or removing the old binary.
