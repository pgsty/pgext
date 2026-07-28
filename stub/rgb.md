## Usage

Sources:

- [Official upstream README](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/README.md)
- [Official extension control file (rgb.control)](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/rgb.control)
- [Official implementation source](https://github.com/nixgeekdev/pg_rgb/blob/0674066b0bec753a4c6b3e1c903cf74400ea809b/src/rgb.c)

`rgb` — This library contains a single PostgreSQL extension, a RGB color data type, along with convenience functions for constructing, converting, and indexing RGB colors. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION rgb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
