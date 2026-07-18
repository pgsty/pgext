## Usage

Sources:

- [Official extension control file](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify.control)

`pg_slugify` — Normalize text into URL-friendly slugs with unaccent

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `unaccent`.

```sql
CREATE EXTENSION "pg_slugify";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_slugify';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
