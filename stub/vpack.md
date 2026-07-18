## Usage

Sources:

- [Official extension control file](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack.control)
- [Official upstream documentation](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/README.md)

`vpack` — VelocyPack data type, path type, operators, casts, and JSON/BSON conversion support for PostgreSQL.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "vpack";
SELECT extversion
FROM pg_extension
WHERE extname = 'vpack';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
