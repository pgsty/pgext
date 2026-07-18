## Usage

Sources:

- [Official extension control file](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views.control)
- [Official upstream documentation](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/README.md)

`xl_global_views` — Creates global pgxl_ views over Postgres-XL nodes based on pg_ views.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "xl_global_views";
SELECT extversion
FROM pg_extension
WHERE extname = 'xl_global_views';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
