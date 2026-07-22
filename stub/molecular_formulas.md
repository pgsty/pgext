## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/molecular_formulas.control)
- [Official Molecular Formula README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/molecular_formulas/extension/usr/share/postgresql/17/extension/molecular_formulas--0.1.0.sql)

`molecular_formulas` packages a Rust parser and utility library for molecular formulas as a PGRX extension artifact. In the reviewed 0.1.0 artifact, the generated extension SQL exposes no user-facing SQL parser, type, or function.

### Core Workflow

```sql
CREATE EXTENSION molecular_formulas;

SELECT extversion
FROM pg_extension
WHERE extname = 'molecular_formulas';
```

This registers the extension only. Rust objects such as `MolecularFormula`, `Ion`, `Atom`, and `Token` are library APIs and are not PostgreSQL objects in the pinned generated SQL.

### Current SQL Surface

- No SQL type, function, operator, view, or configuration parameter is declared in the reviewed generated SQL.
- The README describes the Rust parser as work in progress and lists molecular-orbital validation as future work; do not infer chemical-validity guarantees from parsing alone.

### Operational Notes

Version 0.1.0 is non-relocatable. Its control file sets `superuser = true` and `trusted = false`, so a superuser must create it; it declares no prerequisite extension or preload requirement. Reinspect generated SQL and release notes before an upgrade because the SQL boundary may change independently of the Rust crate API.
