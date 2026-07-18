## Usage

Sources:

- [Upstream README at the reviewed revision](https://bitbucket.org/adunstan/pg-closed-ranges/src/9a736068119bb0f4fa5d59c7cb41198a248a2b04/README.md)
- [Extension control file](https://bitbucket.org/adunstan/pg-closed-ranges/src/9a736068119bb0f4fa5d59c7cb41198a248a2b04/cranges.control)

`cranges` adds closed-form variants of PostgreSQL's three built-in discrete range types: `cdaterange`, `cint4range`, and `cint8range`. Their canonical representation includes both endpoints, unlike the built-in half-open `[)` representation.

```sql
CREATE EXTENSION cranges;
SELECT '[1,5]'::cint4range;
SELECT '[2026-01-01,2026-01-31]'::cdaterange;
```

The types otherwise reuse PostgreSQL range behavior, including operators and indexing support. The two-argument constructor remains hard-wired by PostgreSQL to assume `[)`, so `cint4range(1,5)` is canonicalized to `[1,4]`. Upstream recommends the explicit three-argument constructor or literal syntax to avoid this surprise.

The project has no current release or PostgreSQL compatibility matrix and is hosted on Bitbucket rather than GitHub. Test casts, operators, indexes, dumps, and upgrades on the target PostgreSQL major before adopting the types in persistent schemas.
