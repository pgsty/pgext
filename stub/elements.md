## Usage

Sources:

- [Official control file](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/elements.control)
- [Official Elements README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/README.md)
- [Generated PostgreSQL 17 extension SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/elements_rs/extension/usr/share/postgresql/17/extension/elements--0.1.0.sql)

`elements` packages the EMI Rust chemical-elements crate as a PGRX extension artifact. The Rust crate contains element, isotope, atomic-weight, oxidation-state, valence, bonding, and electron-configuration APIs, but the reviewed 0.1.0 generated SQL exposes no user-facing SQL type or function.

### Core Workflow

```sql
CREATE EXTENSION elements;

SELECT extversion
FROM pg_extension
WHERE extname = 'elements';
```

This registers the extension, but the pinned generated SQL contains only its PGRX header comments. Do not assume that Rust names such as `Element` or `Isotope` are available as PostgreSQL types in this artifact.

### Current SQL Surface

- No SQL type, function, operator, view, or configuration parameter is declared in the reviewed generated SQL.
- The extensive element and isotope API documented by the README belongs to the Rust crate and requires a consumer compiled with the appropriate Rust feature set.

### Operational Notes

Version 0.1.0 is non-relocatable. Its control file sets `superuser = true` and `trusted = false`, so a superuser must create it; it declares no prerequisite extension or preload requirement. Recheck the generated SQL when upgrading, because a later PGRX build may add a SQL surface that is absent from this version.
