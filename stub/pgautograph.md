## Usage

Sources:

- [pgautograph Rust entry point at the reviewed commit](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/lib.rs)
- [property-graph query generator at the reviewed commit](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/src/queries/mod.rs)
- [pgautograph.control at the reviewed commit](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/pgautograph.control)
- [pgautograph Cargo manifest at the reviewed commit](https://github.com/nevindev/pgautograph/blob/a8c59f407f80fb5ed4f1ee8e3ab15745aab03100/Cargo.toml)

`pgautograph` is an early-stage Rust/`pgrx` extension for deriving property-graph DDL from relational foreign keys. The current `list_foreign_keys()` function inspects foreign-key constraints in the `public` schema, returns their source and target metadata, and emits a generated `CREATE PROPERTY GRAPH graph` statement as an INFO message.

### Inspect Foreign Keys

```sql
CREATE EXTENSION pgautograph;

SELECT *
FROM list_foreign_keys();
```

Review the returned rows and the generated INFO message. The current function generates text but does not execute the property-graph statement.

### Caveats

- The reviewed crate version is 0.0.0 and the repository has no README, license declaration, tags, or releases. Treat it as development code.
- The control file sets `superuser = true`, so extension installation requires an appropriately privileged role.
- The generator only inspects the `public` schema and uses simple capitalization and singularization rules when constructing graph labels. Review and edit generated DDL before execution.
- The Cargo manifest declares PostgreSQL 13 through 18 features with `pgrx` 0.16.1; this matches source targets, not a published compatibility guarantee.
