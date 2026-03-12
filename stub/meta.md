

## Usage

> [meta: simplified system catalog for PostgreSQL](https://github.com/aquameta/meta)

meta provides a normalized, human-friendly system catalog with common names for views and columns, sitting on top of `pg_catalog` and `information_schema`.

### System Catalog Views

The extension creates approximately 30 views in the `meta` schema:

```sql
-- List all tables
SELECT * FROM meta.table;

-- List all columns
SELECT * FROM meta.column;

-- List all views
SELECT * FROM meta.view;

-- List schemas
SELECT * FROM meta.schema;

-- List functions
SELECT * FROM meta.function;

-- List extensions
SELECT * FROM meta.extension;

-- List triggers
SELECT * FROM meta.trigger;

-- List foreign keys
SELECT * FROM meta.foreign_key;

-- List constraints
SELECT * FROM meta.constraint_check;
SELECT * FROM meta.constraint_unique;

-- List types
SELECT * FROM meta.type;

-- List roles
SELECT * FROM meta.role;

-- List sequences
SELECT * FROM meta.sequence;

-- List operators
SELECT * FROM meta.operator;

-- List policies
SELECT * FROM meta.policy;
```

### Available Views

`cast`, `column`, `connection`, `constraint_check`, `constraint_unique`, `extension`, `foreign_column`, `foreign_data_wrapper`, `foreign_key`, `foreign_server`, `foreign_table`, `function`, `function_parameter`, `operator`, `policy`, `policy_role`, `relation`, `relation_column`, `role`, `role_inheritance`, `schema`, `sequence`, `table`, `table_privilege`, `trigger`, `type`, `view`

### Meta-Identifiers

The extension provides composite types that serve as "soft" primary keys to identify PostgreSQL objects by name (tables, columns, types, etc.) rather than by OID.

### Catalog Triggers (Optional)

With the optional `meta_triggers` extension, views become updatable, enabling DDL via DML:

```sql
-- Create a table via INSERT
INSERT INTO meta.table (schema_name, name) VALUES ('public', 'foo');
```
