## Usage

Sources:

- [Official extension control file](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/cartodb.control.in)

`cartodb` — Archived PostgreSQL extension for turning a database into a CartoDB user database.

The reviewed catalog snapshot records version `0.37.1`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`, `postgis`.
The curated compatibility set is `11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cartodb";
SELECT extversion
FROM pg_extension
WHERE extname = 'cartodb';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
