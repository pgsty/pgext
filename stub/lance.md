## Usage

Sources:

- [Official README at the reviewed revision](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/README.md)
- [Extension control file](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/lance.control)
- [SQL-visible Rust entry points](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/src/lib.rs)
- [Arrow/Lance type mapping](https://github.com/lance-format/pglance/blob/6dece2f725624f3496816135d08b9b4eec63b219/src/fdw/type_mapping.rs)

`lance` is the PostgreSQL extension produced by the `pglance` Rust crate. Version `0.0.0` exposes local or remote Lance datasets as read-only foreign tables through `lance_fdw`, with helpers for schema discovery and namespace reconciliation.

### Core Workflow

Create the extension and an FDW server, then import a dataset by URI:

```sql
CREATE EXTENSION lance;
CREATE SERVER lance_srv FOREIGN DATA WRAPPER lance_fdw;

SELECT lance_import(
  'lance_srv',
  'public',
  'events_lance',
  '/srv/lance/events.lance',
  batch_size => 2048
);

SELECT * FROM public.events_lance LIMIT 10;
```

`lance_import` inspects the Lance schema, creates the foreign table, and creates PostgreSQL composite types for nested `struct` fields when needed.

For a Lance namespace, inspect the plan before applying DDL:

```sql
SELECT *
FROM lance_attach_namespace(
  'lance_srv',
  schema_prefix => 'lance',
  dry_run => true
);

SELECT *
FROM lance_sync_namespace(
  'lance_srv',
  schema_prefix => 'lance',
  drop_missing => false,
  recreate_changed => false,
  dry_run => true
);
```

### Objects and Options

- `lance_fdw` implements foreign scans; the reviewed handler has no insert, update, or delete callbacks.
- `lance_import(server_name, local_schema, table_name, uri, batch_size)` creates one foreign table from a dataset URI.
- `lance_attach_namespace(...)` creates local schemas, foreign tables, and nested composite types for a namespace tree.
- `lance_sync_namespace(...)` reports or reconciles drift; `drop_missing` and `recreate_changed` enable destructive DDL.
- `uri` can be set on the server or foreign table. `batch_size` defaults to 1024, with a table option overriding the server option.
- Namespace tables use `ns.table_id` as a JSON array of strings; server options under `ns.*` configure the namespace implementation.

Native mappings include Boolean to `boolean`, signed/unsigned integers to the nearest PostgreSQL integer type, strings to `text`, binary values to `bytea`, dates and timestamps to PostgreSQL temporal types, lists to arrays, structs to generated composite types, and maps or unsupported types to `jsonb`.

### Operational Boundaries

The control file marks `lance` as an untrusted superuser-only extension. The reviewed crate advertises PostgreSQL 13–17 build features, but version `0.0.0` is an early interface and should be tested against the exact server and Lance files in use.

Foreign scans read files or remote object storage with the database server's filesystem, network, and credential context. Restrict URIs and storage options accordingly. Always inspect `dry_run` output before namespace attachment or synchronization; enabling `drop_missing` or `recreate_changed` can drop or replace local foreign tables and their dependent objects.
