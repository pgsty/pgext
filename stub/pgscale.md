## Usage

Sources:

- [Official extension control file (pgscale.control)](https://github.com/kelvich/pgscale/blob/5f24b0db82f3bfb48d8f1e5d6ca1441b543050d8/pgscale.control)
- [Official extension SQL (pgscale--1.0.sql)](https://github.com/kelvich/pgscale/blob/5f24b0db82f3bfb48d8f1e5d6ca1441b543050d8/pgscale--1.0.sql)

`pgscale` — Background worker exposing PostgreSQL statistics views through a simple HTTP endpoint. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgscale;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
