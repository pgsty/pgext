## Usage

Sources:

- [Official upstream documentation](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/doc/unnest_ordinality.md)
- [Official PGXN distribution page](https://pgxn.org/dist/unnest_ordinality/)
- [Official upstream README](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/README.md)

`unnest_ordinality` — Legacy array unnest function that returns each element with its ordinal position.

The reviewed catalog snapshot records version `1.0.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "unnest_ordinality";
SELECT extversion
FROM pg_extension
WHERE extname = 'unnest_ordinality';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
