## Usage

Sources:

- [Official repository README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [Extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/pgx_show_hooks.control)
- [Hook inspection implementation](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/src/lib.rs)

`pgx_show_hooks` is a diagnostic extension that reports the in-process addresses currently installed in PostgreSQL's hook variables and PL/pgSQL plugin callbacks. Use it to compare hook state in a backend while testing extension loading and hook conflicts; it identifies occupied callback slots but does not identify every callback's owning extension or change hook order.

### Core Workflow

Create the extension as a superuser, then inspect hook state in the same backend where you need the observation:

```sql
CREATE EXTENSION pgx_show_hooks;

SELECT name, addr
FROM show_hooks.all()
WHERE addr IS NOT NULL
ORDER BY name;
```

`show_hooks.all()` returns one row per hook or PL/pgSQL callback. `addr` is a hexadecimal process address when a callback is present and `NULL` when the slot is empty. Capture a baseline, load or preload the extension under test, and compare results from equivalent backend sessions.

### Object Index

- `show_hooks.all() -> table(name text, addr text)` lists PostgreSQL hook variables, the `PLpgSQL_plugin` rendezvous pointer, and individual PL/pgSQL callback pointers.

### Operational Notes

`pgx_show_hooks` installs into the fixed `show_hooks` schema and requires superuser privileges to create. Its control file does not declare a preload requirement. Results are process-local observations: PostgreSQL backends are separate processes, hook addresses can differ across sessions or restarts, and an extension loaded only in one backend may not appear in another.

Treat `addr` as sensitive diagnostic data. It is useful for equality and presence checks inside a controlled test, not as a stable identifier across processes or releases. The reviewed repository is a PostgreSQL 15-era research project; verify the exact binary and server combination before using it to draw conclusions about hook coverage.
