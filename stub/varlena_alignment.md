## Usage

Sources:

- [Official extension control file](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/varlena_alignment.control)
- [Official Rust package manifest](https://github.com/optionfactory/varlena_alignment/blob/c1bf619811831c577245ee9f390f6174b908a314/Cargo.toml)

`varlena_alignment` — pgrx test extension documenting varlena alignment behavior in pgrx set_varsize helpers.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "varlena_alignment";
SELECT extversion
FROM pg_extension
WHERE extname = 'varlena_alignment';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
