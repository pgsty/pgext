## Usage

Sources:

- [Official documentation](https://www.postgresql.org/docs/12/plpython-python23.html)
- [Official extension control file (plpythonu.control)](https://github.com/postgres/postgres/blob/REL_12_STABLE/src/pl/plpython/plpythonu.control)
- [Official extension SQL (plpythonu--1.0.sql)](https://github.com/postgres/postgres/blob/REL_12_STABLE/src/pl/plpython/plpythonu--1.0.sql)

`plpythonu` — Historical untrusted PL/Python procedural language. Use it when database code must run in or interoperate with this procedural language. The reviewed upstream project is archived or no longer maintained.

### Core Workflow

```sql
CREATE EXTENSION plpythonu;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- These Python 2 language names are historical; current PostgreSQL supports PL/Python through `plpython3u` instead.
- PL/Python is untrusted, so only a superuser can create functions in the language.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
