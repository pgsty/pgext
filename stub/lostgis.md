## Usage

Sources:

- [Official extension control file](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/lostgis.control)
- [Official upstream documentation](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/doc/lostgis.md)
- [Official PGXN distribution page](https://pgxn.org/dist/lostgis/)

`lostgis` — LostGIS spatial helper functions and types

The reviewed catalog snapshot records version `1.0.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `postgis`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "lostgis";
SELECT extversion
FROM pg_extension
WHERE extname = 'lostgis';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
