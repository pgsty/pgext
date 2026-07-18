## Usage

Sources:

- [Official extension control file](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/conda_pgsql_rust_ext.control)
- [Official Rust package manifest](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_rust_ext/Cargo.toml)
- [Official upstream README](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/README.md)

`conda_pgsql_rust_ext` — Conda version data type and helpers for PostgreSQL, implemented with pgrx and rattler_conda_types.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "conda_pgsql_rust_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'conda_pgsql_rust_ext';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
