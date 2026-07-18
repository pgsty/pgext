## Usage

Sources:

- [Official primary documentation](https://database.dev/martindonadieu/supabase_auth_apikey)

`supabase_auth_apikey` — Supabase Auth API-key helpers for creating, deleting, validating, and authorizing user API keys.

The reviewed catalog snapshot records version `0.0.3`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "supabase_auth_apikey";
SELECT extversion
FROM pg_extension
WHERE extname = 'supabase_auth_apikey';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
