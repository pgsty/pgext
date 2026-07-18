## Usage

Sources:

- [Official upstream documentation](https://database.dev/garyaustin/custom_roles)

`custom_roles` — Role-specific Supabase RLS tables and functions for managing user roles.

The reviewed catalog snapshot records version `0.0.3`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "custom_roles";
SELECT extversion
FROM pg_extension
WHERE extname = 'custom_roles';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
