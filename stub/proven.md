## Usage

Sources:

- [Official extension control file](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven.control)

`proven` — proven: PostgreSQL C bindings for formally verified safety functions from libproven

The reviewed catalog snapshot records version `0.5.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "proven";
SELECT extversion
FROM pg_extension
WHERE extname = 'proven';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
