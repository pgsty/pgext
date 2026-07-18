## Usage

Sources:

- [Official extension control file](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/wrapper_deltalake.control)
- [Official upstream documentation](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/README.md)

`wrapper_deltalake` — Unfinished pgrx hello-world prototype originally planned as a Delta Lake foreign data wrapper.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "wrapper_deltalake";
SELECT extversion
FROM pg_extension
WHERE extname = 'wrapper_deltalake';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
