## Usage

Sources:

- [Official upstream README](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/README.md)
- [Official extension control file (ednb.control)](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/ednb/ednb.control)
- [Official extension SQL (ednb--0.0.1.sql)](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/ednb/ednb--0.0.1.sql)

`ednb` — Postgres extension to support EDN as a custom datatype. Use it when application data needs this type, domain, or its operators. Upstream describes it as a work in progress.

### Core Workflow

```sql
CREATE EXTENSION ednb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ednb_in(cstring)` is an extension function and returns `ednb`.
- `ednb_out(ednb)` is an extension function and returns `cstring`.
- `ednb` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
