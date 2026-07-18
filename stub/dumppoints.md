## Usage

Sources:

- [Official extension control file](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints.control)
- [Official upstream documentation](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/README)

`dumppoints` — C implementation of ST_DumpPoints for PostGIS geometry collections.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `postgis`.

```sql
CREATE EXTENSION "dumppoints";
SELECT extversion
FROM pg_extension
WHERE extname = 'dumppoints';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
