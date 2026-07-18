## Usage

Sources:

- [Official extension control file](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby.control)
- [Official upstream documentation](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/README.md)

`plmruby` — Run scalar, set-returning, inline, and trigger functions in the embedded mruby runtime

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "plmruby";
SELECT extversion
FROM pg_extension
WHERE extname = 'plmruby';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
