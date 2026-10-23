## Usage

Sources:

- [README.md](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/README.md)
- [pg_try_cast--1.0.sql](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast--1.0.sql)
- [src/pg_try_cast.c](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/src/pg_try_cast.c)
- [.github/workflows/ci.yml](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/.github/workflows/ci.yml)
- [pg_try_cast.control](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast.control)

`pg_try_cast` provides `try_cast` and `try_convert` for converting text without aborting a query on supported input errors. Release 1.0.0 installs SQL version 1.0. PostgreSQL 16 or later is required; upstream CI covers 16–18. It is trusted and relocatable, with no preload requirement.

### Core Workflow

```sql
CREATE EXTENSION pg_try_cast;
SELECT try_cast('42', NULL::integer);
SELECT try_cast('invalid', 0);
SELECT try_convert(NULL::date, '2026-02-30');
```

### Operational Boundaries

`try_cast(value text, fallback anyelement)` takes its result type and error fallback from the second argument; `try_convert(fallback anyelement, value text)` reverses the order. NULL input returns NULL. Give the fallback an explicit type. Type modifiers and domain constraints are checked. This covers text input, not arbitrary native casts. Input functions without soft-error support can still raise an error, including malformed type names passed to regtype. A role with CREATE privilege on the database can install this trusted extension after its files are installed.
