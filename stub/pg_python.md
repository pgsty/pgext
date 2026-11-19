## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_python/pg_python.c)

`pg_python` `1.0` provides registered python scripts executed through pl/python. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pg_python;
SELECT python_eval('return args["a"] + args["b"]', '{"a":2,"b":3}');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

python_register and python_exec version and run scripts; arguments are passed as a JSON object. Execution uses the in-process untrusted plpython3u runtime. The environment registry and pg_python.executable are metadata, not process isolation or a package installer. Restrict the runner, registry and wrapper functions to trusted users: Python can access server resources, and external effects do not roll back with SQL.

This module corresponds to the former upstream name `alohadb_python`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
