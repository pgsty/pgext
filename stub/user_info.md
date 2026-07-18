## Usage

Sources:

- [Official upstream documentation](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/doc/user_info.md)
- [Official PGXN distribution page](https://pgxn.org/dist/user_info/)

`user_info` — Provides user-related information including owned objects and granted roles.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "user_info";
SELECT extversion
FROM pg_extension
WHERE extname = 'user_info';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
