## Usage

Sources:

- [Official upstream documentation](https://database.dev/mansueli/rls_helpers)

`rls_helpers` — PL/pgSQL procedures for simulating Supabase authenticated and anonymous roles while testing row-level security policies.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "rls_helpers";
SELECT extversion
FROM pg_extension
WHERE extname = 'rls_helpers';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
