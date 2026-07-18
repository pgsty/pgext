## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/README.md)
- [Version 1.0 SQL objects](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/pg_elephant--1.0.sql)
- [C implementation](https://github.com/gerdansantos/pg_elephant/blob/ba00d779ccc306f36db14ec86a89208d279c2328/pg_elephant.c)

`pg_elephant` is explicitly a toy PostgreSQL C extension and build example. It installs one demonstration function.

```sql
CREATE EXTENSION pg_elephant;
SELECT pg_elephant();
```

Use it only to verify PGXS compilation, extension packaging, and function loading in a disposable development cluster. It does not provide persistence, caching, audit history, or any production feature despite the “never forgets” slogan.

The README targets old PostgreSQL 9.5 development packages and supplies no current compatibility, maintenance, or security contract. Do not add an example-only native library to production attack surface. If adapting the code as a template, replace names and metadata, add strict input validation, correct volatility/parallel/security markings, regression tests, supported-version CI, an upgrade path, and a real license review before distributing the result.
