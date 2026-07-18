## Usage

Sources:

- [Official extension control file](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/pg_dbo_timestamp.control)
- [Official upstream documentation](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/README.md)

`pg_dbo_timestamp` — PostgreSQL extension for storing time and author of database structure modification.

The reviewed catalog snapshot records version `1.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_dbo_timestamp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_dbo_timestamp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
