## Usage

Sources:

- [Official extension control file](https://github.com/asotolongo/jx_io/blob/4098d0db4d71596deed1e32a71c587bb1285c01a/jx_io.control)
- [Official PGXN distribution page](https://pgxn.org/dist/jx_io/)
- [Official project or provider page](http://anthonysotolongo.wordpress.com)

`jx_io` — Extension to export and import json and xml from/to postgres

The reviewed catalog snapshot records version `0.1.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "jx_io";
SELECT extversion
FROM pg_extension
WHERE extname = 'jx_io';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
