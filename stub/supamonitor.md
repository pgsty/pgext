## Usage

Sources:

- [Official extension control file](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/supamonitor.control)
- [Official upstream documentation](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/src/lib.rs)
- [Official Rust package manifest](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/Cargo.toml)

`supamonitor` — Preloaded pgrx query hook that logs query text, query ID, planning time, and execution time as JSON.

The reviewed catalog snapshot records version `0.0.6`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "supamonitor";
SELECT extversion
FROM pg_extension
WHERE extname = 'supamonitor';
```

The upstream project is associated with `Supabase`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
