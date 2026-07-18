## Usage

Sources:

- [Official extension control file](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/clerk_fdw.control)
- [Official upstream documentation](https://github.com/Jayko001/clerk_fdw/blob/ff55a536f658b2e4461407b8e10b3e47a6dfb5d1/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/clerk_fdw/)

`clerk_fdw` — Foreign data wrapper for querying the Clerk user-management API from PostgreSQL.

The reviewed catalog snapshot records version `0.3.3`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "clerk_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'clerk_fdw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
