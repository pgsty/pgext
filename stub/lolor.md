## Usage

Sources:

- [v1.2.2 README](https://github.com/pgEdge/lolor/blob/v1.2.2/README.md)
- [Control file](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor.control)
- [Version 1.2.2 migration](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor--1.2.1--1.2.2.sql)
- [Usage reference](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/using_lolor.md)
- [Release notes](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/lolor_release_notes.md)
`lolor` 1.2.2 stores large-object chunks and metadata in ordinary tables so logical replication can include them. Enabling it replaces the database's native large-object routines; it is a database-wide behavior change, not an independent object store.

### Configure and Enable

Assign each writing node a distinct nonzero `lolor.node` before creating large objects. Upstream documents values from 1 through 2^28. Configure the server's parameter and create the extension in each participating database:

```conf
lolor.node = 1
```

```sql
CREATE EXTENSION lolor;
SET search_path = lolor, "$user", public, pg_catalog;
```

The control file fixes schema `lolor`, declares trusted installation and disallows relocation. The upstream README requires PostgreSQL 16 or newer; the current Pigsty pgEdge bundle has separately tested builds for 15–18. That packaging result does not establish support for arbitrary stock PostgreSQL 15 installations.

### Large Object Workflow

```sql
WITH created AS (
  SELECT lo_from_bytea(0, convert_to('example data', 'UTF8')) AS oid
)
SELECT oid, convert_from(lo_get(oid), 'UTF8') AS contents FROM created;
```

Standard calls such as `lo_create()`, `lo_get()`, `lo_put()` and `lo_unlink()` use the replacement routines. Descriptor-based access through `lo_open()`, `loread()`, `lowrite()` and `lo_close()` must stay within the same transaction. File import/export operates on server-side paths and remains subject to the relevant function and filesystem privileges.

The extension uses `lolor.pg_largeobject` and `lolor.pg_largeobject_metadata`. Existing catalog routines are retained with renamed originals so disabling or removing the extension can restore the native function names.

### Replication

With a configured Spock replication set, include both tables:

```sql
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

Install compatible versions and coordinate node identifiers and object ownership on every node. Creating the extension alone does not create a logical-replication topology.

### Upgrade and Boundaries

Version 1.2.2 repairs extension and major-version upgrade behavior and adds `lolor.disable()`, `lolor.enable()` and `lolor.is_enabled()`. Update lolor **before** running pg_upgrade; the fix does not retroactively repair a major upgrade already attempted with older extension files:

```sql
ALTER EXTENSION lolor UPDATE TO '1.2.2';
SELECT lolor.is_enabled();
```

Disabling changes which native function names are active; it does not migrate stored large objects. Native large-object migration is not provided. Native large-object functionality and lolor storage cannot be used interchangeably while lolor is enabled. Upstream excludes ALTER LARGE OBJECT, GRANT ON LARGE OBJECT, COMMENT ON LARGE OBJECT and REVOKE ON LARGE OBJECT. Plan backup, restore and replication for both ordinary tables.
