


## Usage

> [pg_permissions: view object permissions and compare them with the desired state](https://github.com/cybertec-postgresql/pg_permissions)

pg_permissions lets you review actual permissions on database objects and compare them against a desired permission state.

### Define Desired Permissions

Insert entries into `permission_target` to describe what permissions should exist:

```sql
INSERT INTO permission_target (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{SELECT,INSERT,UPDATE,DELETE}', 'TABLE', 'appschema');

INSERT INTO permission_target (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{USAGE}', 'SCHEMA', 'appschema');

INSERT INTO permission_target (role_name, permissions, object_type, schema_name, object_name)
VALUES ('appuser', '{USAGE}', 'SEQUENCE', 'appschema', 'appseq');
```

Set `object_name` or `column_name` to NULL to apply to all objects of that type in the schema.

### Find Permission Differences

```sql
SELECT * FROM permission_diffs();
```

Returns rows where `missing = TRUE` (permission should exist but doesn't) or `missing = FALSE` (extra permission that shouldn't exist).

### Review Actual Permissions

Available views (all with the same column structure):

- `database_permissions` -- permissions on the current database
- `schema_permissions` -- permissions on schemas
- `table_permissions` -- permissions on tables
- `view_permissions` -- permissions on views
- `column_permissions` -- permissions on table/view columns
- `function_permissions` -- permissions on functions
- `sequence_permissions` -- permissions on sequences
- `all_permissions` -- UNION of all above

```sql
SELECT * FROM table_permissions WHERE role_name = 'appuser' AND schema_name = 'appschema';
```

### Grant/Revoke via Views

The `granted` column of the permission views is updatable -- updating it executes the appropriate `GRANT` or `REVOKE` command.

Note: superusers are not shown in the views (they automatically have all permissions).
