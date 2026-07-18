## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/README.md)
- [Version 0.0.1 SQL objects](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby--0.0.1.sql)
- [Call-handler implementation](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby.c)

`plruby` is an extremely early procedural-language prototype embedding Ruby in a PostgreSQL backend. Version 0.0.1 installs the untrusted language `plruby` with call, inline, and validator handlers; upstream describes it only as “sort of working” on one computer.

```sql
CREATE EXTENSION plruby;
CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$ args[0] + args[1] $$;
SELECT ruby_add(2, 3);
```

The example is suitable only for an isolated build test after reviewing the handler's actual argument and return conventions. An untrusted procedural language can execute with the database server process's operating-system authority and must remain superuser-only.

Do not deploy this prototype in production or expose it to tenant-supplied code. It copies substantial early PL/V8 patterns, has no current Ruby/PostgreSQL compatibility matrix, and makes no sandbox, memory, exception, signal, threading, garbage-collection, or transaction-safety guarantees. Prefer a maintained procedural language or execute Ruby outside the database. Any archival test belongs in a disposable host with no secrets or network trust.
