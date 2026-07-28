## Usage

Sources:

- [Official upstream README](https://github.com/dobpilot/fasttrun/blob/f75ebbcc5a99fd2c0698b44bd316c29a5177b1ea/README.md)
- [Official extension control file (fasttrun.control)](https://github.com/dobpilot/fasttrun/blob/f75ebbcc5a99fd2c0698b44bd316c29a5177b1ea/fasttrun.control)
- [Official extension SQL (fasttrun--2.0.sql)](https://github.com/dobpilot/fasttrun/blob/f75ebbcc5a99fd2c0698b44bd316c29a5177b1ea/fasttrun--2.0.sql)

`fasttrun` — Optimezed fasttruncate extension for 1C:Enterprise 8. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION fasttrun;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `fasttruncate(text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `2.0`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
