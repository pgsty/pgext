## Usage

Sources:

- [Official extension control file](https://github.com/mansueli/tle/blob/3825e5a6dd3b3d0e48eea1a01eba89655788b746/email_guard/email_guard.control)
- [Official upstream documentation](https://database.dev/mansueli/email_guard)
- [Official upstream README](https://github.com/mansueli/tle/blob/3825e5a6dd3b3d0e48eea1a01eba89655788b746/README.md)

`email_guard` — Supabase Auth signup guard that blocks disposable email domains and normalizes Gmail addresses.

The reviewed catalog snapshot records version `0.35.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "email_guard";
SELECT extversion
FROM pg_extension
WHERE extname = 'email_guard';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
