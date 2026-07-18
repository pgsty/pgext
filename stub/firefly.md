## Usage

Sources:

- [Pinned upstream README](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/README.md)
- [Version 0.2.0 installation SQL](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/firefly--0.2.0.sql)
- [Pinned extension control file](https://github.com/celian-garcia/firefly-database/blob/1dfe6cf49a670b4d3b2c1ace40969cdab58189fa/firefly.control)

firefly version 0.2.0 is the database schema for a specific 3D point-editing application. It creates task and fpoint3d tables, per-task operation sequences, add/delete history functions, collection functions, triggers, enums, and a destructive reset helper. It requires PL/pgSQL and PostGIS; it is not a general change-data-capture extension.

### Application example

Use a new database because the extension creates generic unqualified object names:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION firefly;

INSERT INTO task (name, description)
VALUES ('demo', 'isolated firefly test')
RETURNING id;

SELECT save_add_operation(
    ST_Force3D(ST_Point(1, 2)),
    0
);

SELECT * FROM collect_operations(0, 0);
```

Each task insert dynamically creates a sequence named from its numeric task ID. Point rows store an integer array whose odd/even history is interpreted as add or delete state.

### Application and lifecycle limits

The SQL hard-codes tables, sequences, types, and functions without schema qualification or protected search paths. Although the control file says relocatable, moving the extension or changing search_path can break name resolution or make functions use an unintended object. The control file also points module_pathname at pgtap even though the installation is pure SQL, indicating incomplete packaging cleanup.

The implementation uses distance-based point matching without a uniqueness constraint, creates one sequence per task, and scans operation arrays linearly. Concurrent edits, duplicate near-equal points, task deletion, sequence ownership, schema relocation, and reset behavior need application-specific tests. The empty_firefly_database function truncates all application state and resets sequences.

The upstream recipe targets PostgreSQL 9.6 and the repository has not changed since 2018. Install only in an empty disposable database, revoke public function execution, grant only the application role, and keep independent backups before evaluating this historical application schema.
