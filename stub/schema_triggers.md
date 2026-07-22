## Usage

Sources:

- [Archived official README](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/README.md)
- [Extension SQL and event-info types](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers--0.1.sql)
- [Extension control file](https://github.com/CartoDB/pg_schema_triggers/blob/07469164cd50460ff4f1d76b36c0a03f26d91e92/schema_triggers.control)

`schema_triggers` version `0.1` extends PostgreSQL's event-trigger machinery with events for relation, column, and trigger changes. The project is archived, was last updated in 2015, and hooks PostgreSQL internals; use it only with an explicitly validated legacy server build.

### Enablement

The extension must both install its database objects and load its hook library. Add the library to configuration and restart PostgreSQL, then install it in each database that needs the helper types and functions:

```conf
shared_preload_libraries = 'schema_triggers.so'
```

```sql
CREATE EXTENSION schema_triggers;
```

For isolated testing only, the upstream README permits `LOAD 'schema_triggers.so'` in one session instead of server preload.

### Core Workflow

```sql
CREATE FUNCTION audit_relation_create()
RETURNS event_trigger
LANGUAGE plpgsql
AS $$
DECLARE
  event_info schema_triggers.relation_create_eventinfo;
BEGIN
  event_info := schema_triggers.get_relation_create_eventinfo();
  RAISE NOTICE 'created relation %', event_info.relation;
END;
$$;

CREATE EVENT TRIGGER audit_relation_create_event
ON relation_create
EXECUTE PROCEDURE audit_relation_create();
```

### Events and Helpers

The new events are `relation_create`, `relation_alter`, `relation_drop`, `column_add`, `column_alter`, `column_drop`, `trigger_create`, and `trigger_drop`. Each event has a matching `get_*_eventinfo()` function returning a composite record with relevant catalog rows and identifiers.

### Safety and Compatibility

The control file requires superuser installation. Trigger functions run synchronously during DDL and can reject or disrupt schema changes; keep them small and failure-aware. Event-info records expose version-sensitive `pg_class`, `pg_attribute`, and `pg_trigger` row layouts. Because the repository is explicitly unmaintained and predates modern PostgreSQL releases, do not assume binary or catalog compatibility from the historical “PostgreSQL 9.3 and higher” statement.
