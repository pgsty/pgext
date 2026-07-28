## Usage

Sources:

- [Official upstream README](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/README.md)
- [Official extension control file (verhoeff.control)](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/verhoeff.control)
- [Official extension SQL (verhoeff--1.0.sql)](https://github.com/pschlump/pgverhoeff/blob/f433b15dbde363eeea0a918517adef7088a61dec/verhoeff--1.0.sql)

`verhoeff` — Postgres offers powerful extensibility features that allow developers to enhance its functionality. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION verhoeff;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `generate_verhoeff(inp text)` is an extension function and returns `text`.
- `validate_verhoeff(inp text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
