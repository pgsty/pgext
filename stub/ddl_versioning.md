## Usage

Sources:

- [Official ddl_versioning README](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/README.ddl_versioning)
- [Version 1.0 extension SQL](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning--1.0.sql)
- [Extension control file](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning.control)

`ddl_versioning` records successive textual definitions of selected database objects whenever a supported DDL command finishes. It is useful as a lightweight audit aid for tables, indexes, functions, and views, but it is not a complete schema-backup or rollback system.

### Core workflow

The extension is trusted, fixed to the `public` schema, and not relocatable:

```sql
CREATE EXTENSION ddl_versioning;

CREATE TABLE public.accounts (
    id bigint PRIMARY KEY,
    balance numeric NOT NULL
);

ALTER TABLE public.accounts ADD COLUMN updated_at timestamptz;
```

The database-wide `ddl_versioning_trigger` event trigger runs on `ddl_command_end`. Inspect the object catalog and its saved versions as follows:

```sql
SELECT object_id, object_type, object_identity
FROM public.ddl_versioning_object
ORDER BY object_id;

SELECT o.object_identity, v.version_id, v.object_definition,
       v.created_at, v.created_by
FROM public.ddl_versioning_object AS o
JOIN public.ddl_versioning_version AS v USING (object_id)
ORDER BY o.object_identity, v.version_id;
```

### Recorded objects

- Tables are reconstructed by `ddl_versioning_get_tabledef`; indexes, including separately created constraints and indexes, are recorded independently.
- Indexes use `pg_get_indexdef`, functions use `pg_get_functiondef`, and views use `pg_get_viewdef`.
- Other DDL tags can reach the event trigger, but unsupported object types are logged rather than converted into restorable definitions.

Each supported DDL command appends another row. The extension provides no retention policy, comparison UI, replay procedure, or automatic rollback, so prune history and build restore tooling separately if required. Table reconstruction does not capture every schema property as a self-contained migration; test any generated definition before using it for recovery. The global event trigger also adds writes and logging to DDL operations, so evaluate its overhead and permissions on the target PostgreSQL release.
