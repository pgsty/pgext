## Usage

Sources:

- [Official upstream documentation](https://database.dev/danahartweg/pg_protect_columns)

`pg_protect_columns` — A set of postgres functions that allow you to protect columns that should not be updated.

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "pg_protect_columns";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_protect_columns';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
