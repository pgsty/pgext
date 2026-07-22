## Usage

Sources:

- [Official README](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/README.md)
- [Official extension control file](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/meta_triggers.control)
- [Official trigger definitions](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/000-triggers.sql)

`meta_triggers` makes selected views from Aquameta's `meta` extension writable. Inserts, updates, and deletes against those catalog views are converted into real DDL or administrative commands, allowing metadata-driven tools to create and change database objects through rows.

### Core Workflow

The extension is installed into the `meta` schema and requires `meta`:

```sql
CREATE EXTENSION meta;
CREATE EXTENSION meta_triggers;

INSERT INTO meta.schema(name) VALUES ('app_data');
UPDATE meta.schema
SET name = 'application_data'
WHERE name = 'app_data';
DELETE FROM meta.schema WHERE name = 'application_data';
```

The upstream SQL adds `INSTEAD OF` triggers and statement-building functions for many `meta` views, including schemas, types, sequences, tables, views, columns, foreign keys, functions, triggers, roles and memberships, privileges, row-security policies, constraints, extensions, foreign-data objects, casts, and event triggers. Exact writable fields differ by view; use the corresponding `meta` view definition as the object model.

### Security and Transaction Behavior

Writes execute DDL with the privileges of the calling role and normally participate in the caller's transaction. Treat write access to these views as administrative access: a row change can create or drop schemas, roles, functions, extensions, policies, or other database objects. The SQL constructs commands dynamically, so validate application input and grant privileges narrowly. `meta_triggers` is not a detached metadata store—the underlying PostgreSQL catalog is changed immediately.
