## Usage

Sources:

- [Version 4.0.9 SQL](https://github.com/pgroonga/pgroonga/blob/4.0.9/data/pgroonga_database.sql)
- [Version 4.0.9 control](https://github.com/pgroonga/pgroonga/blob/4.0.9/pgroonga_database.control)
- [Version 4.0.9 implementation](https://github.com/pgroonga/pgroonga/blob/4.0.9/src/pgroonga-database.c)
- [Official recovery procedure](https://pgroonga.github.io/reference/functions/pgroonga-database-remove.html)

`pgroonga_database` 4.0.9 is a recovery-only helper for a damaged internal PGroonga database. It provides one SQL function that removes PGroonga files from database and applicable tablespace directories. It does not provide the search access method.

### Recovery Workflow

Normal index corruption may be repairable with REINDEX alone. Use this module only when the internal Groonga database itself is damaged and a planned rebuild is necessary. Arrange a recovery window and disconnect every session using PGroonga before proceeding; remaining sessions may crash when their files are removed.

In a fresh administrative connection that has not opened any PGroonga index:

```sql
CREATE EXTENSION pgroonga_database;
SELECT pgroonga_database_remove();
```

Disconnect that connection immediately afterward. In another fresh connection, run REINDEX for **every** PGroonga index to recreate the internal database from PostgreSQL table data. Resume application traffic only after rebuilding and checking all affected indexes.

### Return Value and Boundaries

`pgroonga_database_remove()` returns true when it reaches the end of its cleanup loop. If a tablespace ownership check fails, the loop stops and the function can still return true; this result alone does not prove that every location was cleaned. Other failures can raise errors. It removes the internal files directly; it neither exports them nor rebuilds indexes. Do not use other PGroonga features in the cleanup connection. This is not a routine vacuum, an uninstall command, or an action that can be made safe merely by wrapping the call in a SQL transaction.

The control file does not mark the extension trusted or relocatable. The C implementation checks tablespace ownership when traversing locations. Use an administrator who owns the required locations and verify that the cleanup and full reindex completed. The module has no preload requirement; enable it only for the recovery task.
