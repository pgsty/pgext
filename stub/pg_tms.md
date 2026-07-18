## Usage

Sources:

- [Official extension control file](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/pg_tms.control)
- [Official upstream documentation](https://github.com/jkeifer/pg_tms/blob/3e11a51f42095632d36bcd8b9775da41edaa8cde/README.md)

`pg_tms` — functions to tile rasters in TMS format

The reviewed catalog snapshot records version `0.0.4`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `postgis`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_tms";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_tms';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
