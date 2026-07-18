## Usage

Sources:

- [Official extension control file](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/variant.control)
- [Official PGXN distribution page](https://pgxn.org/dist/variant/)
- [Official upstream README](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/README.md)

`variant` — Variant data type that can store values from other PostgreSQL types.

The reviewed catalog snapshot records version `1.0.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "variant";
SELECT extversion
FROM pg_extension
WHERE extname = 'variant';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
