## Usage

Sources:

- [Official extension control file](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/pg_sysdatetime.control)
- [Official upstream documentation](https://github.com/2ndQuadrant/pg_sysdatetime/blob/e1b4a25b11ef7acdb493dc7606c91d9d11d11a52/README.md)

`pg_sysdatetime` — High-precision SQL Server-style system datetime functions for PostgreSQL, mainly useful on older Windows builds.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_sysdatetime";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sysdatetime';
```

The upstream project is associated with `2ndQuadrant`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
