## Usage

Sources:

- [Official extension control file](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/rostgis.control)
- [Official upstream documentation](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/docs/README.md)
- [Official Rust package manifest](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/Cargo.toml)

`rostgis` — Experimental PostGIS-compatible spatial extension written in Rust.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "rostgis";
SELECT extversion
FROM pg_extension
WHERE extname = 'rostgis';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
