## Usage

Sources:

- [Official extension control file](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/vmagentfdw.control)
- [Official Rust package manifest](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/Cargo.toml)

`vmagentfdw` — Foreign data wrapper for reading vmagent scrape-target status through its HTTP API.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "vmagentfdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'vmagentfdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
