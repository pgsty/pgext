## Usage

Sources:

- [Official extension control file](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/ajbool.control)
- [Official upstream documentation](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/ajbool/)

`ajbool` — Three-valued boolean type with a pseudo-NULL unknown state for primary keys.

The reviewed catalog snapshot records version `0.0.3`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ajbool";
SELECT extversion
FROM pg_extension
WHERE extname = 'ajbool';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
