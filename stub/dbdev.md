## Usage

Sources:

- [Official upstream documentation](https://database.dev/supabase/dbdev)
- [Official project or provider page](https://database.dev/)
- [Official upstream README](https://github.com/supabase/dbdev/blob/8b3b966d97f87027a36c052e679ebb45f7e25736/README.md)

`dbdev` — In-database client for installing packages from the database.dev registry.

The reviewed catalog snapshot records version `0.0.5`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `http`, `pg_tle`.

```sql
CREATE EXTENSION "dbdev";
SELECT extversion
FROM pg_extension
WHERE extname = 'dbdev';
```

The upstream project is associated with `Supabase`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
