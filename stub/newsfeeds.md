## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds.control)
- [Official upstream documentation](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/README.md)

`newsfeeds` — PostgreSQL extension for gathering news

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "newsfeeds";
SELECT extversion
FROM pg_extension
WHERE extname = 'newsfeeds';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
