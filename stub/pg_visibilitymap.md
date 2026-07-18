## Usage

Sources:

- [Upstream README](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/README.md)
- [Extension control file](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/pg_visibilitymap.control)
- [SQL installation script](https://github.com/MasahikoSawada/pg_visibilitymap/blob/993d358dc9e9c3af003cee35c63780bfd24e4b5a/pg_visibilitymap--1.0.sql)
- [Supported pg_visibility documentation](https://www.postgresql.org/docs/current/pgvisibility.html)

`pg_visibilitymap` version `1.0` inspects the all-visible bits in a table or materialized view's visibility map. It returns all map entries through `pg_visibilitymap(regclass)` or checks one heap block with `pg_is_all_visible(regclass,bigint)`.

### Example

```sql
CREATE EXTENSION pg_visibilitymap;
SELECT * FROM pg_visibilitymap('public.orders'::regclass);
SELECT pg_is_all_visible('public.orders'::regclass, 0);
```

The install script revokes both functions from `PUBLIC`; grant access only to trusted diagnostic roles. A visibility-map bit is maintenance metadata, not an independent proof that every tuple is visible to a particular snapshot. This 2015-era extension is deprecated and depends on old PostgreSQL internals. Prefer PostgreSQL's maintained `pg_visibility` extension for current systems.
