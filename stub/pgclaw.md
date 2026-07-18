## Usage

Sources:

- [Official extension control file](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/pgclaw.control)
- [Official upstream documentation](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/README.md)
- [Official Rust package manifest](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/Cargo.toml)

`pgclaw` — Store LLM or OpenClaw agents as PostgreSQL values and process watched rows with a background worker

The reviewed catalog snapshot records version `0.1.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgclaw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgclaw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
