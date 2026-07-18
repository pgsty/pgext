## Usage

Sources:

- [Official extension control file](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.control)
- [Official upstream documentation](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/README.md)

`recommend` — Experimental C extension providing scalar distance, text similarity, and one-dimensional array similarity functions derived from smlar.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "recommend";
SELECT extversion
FROM pg_extension
WHERE extname = 'recommend';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
