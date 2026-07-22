## Usage

Sources:

- [Official README](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/README.md)
- [Extension SQL for version 9.4](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc--9.4.sql)
- [Extension control file](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc.control)

`pgorafunc` supplies a subset of Oracle-compatible functions, operators, data types, and package-style APIs for PostgreSQL. It is primarily a migration aid for applications that expect familiar Oracle date behavior or packages such as `dbms_output` and `utl_file`; it is not a complete Oracle compatibility layer.

### Core Workflow

Install the extension in the target database and put its `oracle` schema early enough in the session search path when unqualified Oracle-compatible names are required:

```sql
CREATE EXTENSION pgorafunc;

SET search_path TO oracle, "$user", public, pg_catalog;

SELECT oracle.add_months(oracle.date '2005-05-31 10:12:12', 1);
SELECT oracle.months_between(
  oracle.date '1995-02-02 10:00:00',
  oracle.date '1995-01-01 10:21:11'
);
SELECT oracle.to_date('02/16/09 04:12:12', 'MM/DD/YY HH24:MI:SS');
```

The search-path choice is material: the extension installs compatibility functions both in `oracle` and alongside built-in-visible names, and the README explicitly recommends the shown ordering for Oracle-style resolution. Schema-qualify calls in shared applications when changing a global search path would be unsafe.

### Important Objects

- `oracle.date` preserves a time-of-day component and has Oracle-style arithmetic overloads.
- Date helpers include `add_months`, `last_day`, `next_day`, `months_between`, `trunc`, `round`, and `to_date`. Their accepted format models and rounding behavior should be tested against the migrated workload rather than assumed from PostgreSQL's similarly named functions.
- `dbms_output` provides a session-local line queue with operations such as `enable`, `put_line`, `get_line`, and `get_lines`.
- `utl_file` exposes server-side file operations, including opening, reading, writing, copying, renaming, and removing files.
- Compatibility objects also include the `dual` table and additional Oracle-style packages documented by the project.

For session-local output, enable the queue, add lines, and then consume them:

```sql
SELECT dbms_output.enable();
SELECT dbms_output.put_line('migration step complete');
SELECT * FROM dbms_output.get_lines(0);
```

### Operational Notes

The control file declares version 9.4, a relocatable extension, and the module library, but no preload requirement or extension dependency. This is old, abandoned compatibility code; validate every required function and PostgreSQL-version combination before adopting it for a new migration.

Treat `utl_file` as a privileged server-side capability. Its operations run against files accessible to the PostgreSQL server process, so restrict who may execute them and test paths and permissions in a non-production environment. The `dbms_output` queue is session-local and disappears with the session. Existing Oracle applications may also depend on subtle implicit casts, locale rules, or package semantics outside the implemented subset.
