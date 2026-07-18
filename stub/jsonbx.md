## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/README.md)
- [Version 1.0 SQL objects](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/jsonbx--1.0.sql)
- [C implementation](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/jsonbx.c)
- [PGXN 1.0.0 documentation](https://pgxn.org/dist/jsonbx/1.0.0/README.html)

`jsonbx` was a PostgreSQL 9.4 compatibility extension for JSONB manipulation. It supplied pretty-printing, concatenation, deletion by key/index/path, and path replacement before those functions and operators entered PostgreSQL core in 9.5.

```sql
SELECT jsonb_pretty('{"a":1,"b":2}'::jsonb);
SELECT '{"a":1}'::jsonb || '{"b":2}'::jsonb;
SELECT '{"a":1,"b":2}'::jsonb - 'a';
```

Do not install `jsonbx` on supported PostgreSQL releases. The same object names already exist in core, so the extension script can conflict during `CREATE EXTENSION`; use the built-in functions and operators shown above instead.

Version 1.0 is historical code tied to the PostgreSQL 9.4 internal JSONB API. It should only be preserved for forensic work on an isolated legacy 9.4 installation. Remove the extension dependency before upgrading such a database, while retaining application SQL that now resolves to core objects.
