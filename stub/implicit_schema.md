## Usage

Sources:

- [Official README and warning](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/README.md)
- [Version 0.1 installation SQL](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema--0.1.sql)
- [Event-trigger implementation](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.c)
- [Extension control file](https://github.com/dbaston/pg_implicit_schema/blob/008de73400ab2aa46d4b89c0494df28a21ad4859/implicit_schema.control)

`implicit_schema` installs an event trigger that creates a missing schema when a schema-qualified table or view is created. It is a narrow compatibility hack for software that generates schema names but cannot provision the schemas itself.

### Core Workflow

```sql
CREATE EXTENSION implicit_schema;

CREATE TABLE fizz.fuzz (id integer);

SELECT nspname
FROM pg_namespace
WHERE nspname = 'fizz';
```

The extension defines `auto_create_schema()` and an event trigger named `auto_create_schema` on `ddl_command_start`. Version `0.1` handles only `CREATE TABLE` and `CREATE VIEW`; other object types do not acquire this behavior.

### Operational and Security Boundaries

Upstream explicitly calls the extension a bad idea. It silently turns one object-creation attempt into an additional `CREATE SCHEMA IF NOT EXISTS` operation, changing privilege, audit, locking, and error expectations cluster-wide in the database where it is installed.

The reviewed C implementation interpolates the requested schema name into dynamic SQL without escaping an embedded double quote. Do not expose schema-qualified object creation to untrusted names or users. Prefer explicit schema provisioning; if compatibility requires this extension, isolate it, test adversarial identifiers and concurrent DDL, and audit every role that can create tables or views.
