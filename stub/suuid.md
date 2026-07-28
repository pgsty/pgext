## Usage

Sources:

- [Official upstream README](https://github.com/jwdeitch/pg_suuid/blob/5cab05c4eb9d6989eb6d0372116ec65079ec9cd2/readme)
- [Official extension control file (suuid.control)](https://github.com/jwdeitch/pg_suuid/blob/5cab05c4eb9d6989eb6d0372116ec65079ec9cd2/suuid.control)

`suuid` — small UUID. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION suuid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
