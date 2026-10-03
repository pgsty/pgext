## Usage

Sources:

- [pg_dynamic.control](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/pg_dynamic.control)
- [README.md](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/README.md)
- [pg_dynamic--0.1.0.sql](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/pg_dynamic--0.1.0.sql)
- [Makefile](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/Makefile)
- [regress/sql/integer.sql](https://github.com/JoshInnis/pg_dynamic/blob/921415dec1750ce48b8011a1cead46a8820cb7a6/regress/sql/integer.sql)

`pg_dynamic` introduces a `dynamic` type that carries values for selected PostgreSQL types and exposes casts and overloaded operations. The 0.1.0 source is experimental; its stated goal of supporting every type is broader than the implemented SQL surface.

### Core Workflow

```sql
CREATE EXTENSION pg_dynamic;
SELECT (42::bigint)::dynamic;
SELECT ((42::bigint)::dynamic)::bigint;
```

### Operational Boundaries

Superuser installation is required; the library does not require shared preload. The reviewed SQL includes bigint, inet and box conversions, arithmetic, geometric operations and selected math functions. Use only the concrete casts/functions exported by this revision; arbitrary extension types and automatic universal dispatch are not established. No PostgreSQL major-version support range is asserted by the reviewed documentation.
