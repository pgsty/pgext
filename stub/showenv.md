## Usage

Sources:

- [Official upstream README](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/README.md)
- [Official extension control file (showenv.control)](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/showenv.control)
- [Official extension SQL (showenv--1.0.sql)](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/showenv--1.0.sql)

`showenv` — A PostgreSQL extension that shows the server process's environment variables as a SQL result set. Use it when administering or automating the database behavior described above. The reviewed upstream project is archived or no longer maintained.

### Core Workflow

```sql
CREATE EXTENSION showenv;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `environment_variables()` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as trusted.
- Upstream material indicates that the project is abandoned or unmaintained.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
