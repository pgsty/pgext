## Usage

Sources:

- [Version 0.0.1 installation SQL](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris--0.0.1.sql)
- [Extension control file](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/pg_abris.control)
- [Official upstream build notes](https://github.com/abris-platform/pg_abris/blob/cca112ac51e23ee130be17a25cb0724a085aa58e/README.md)

`pg_abris` is a pure PL/pgSQL metadata layer for the Abris platform. It creates a fixed `meta` schema whose views expose database objects as Abris entities, properties, relations, and projections, while triggers make parts of that metadata surface writable.

### Core Workflow

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION pg_abris;

SELECT entity_id, schema_name, table_name, primarykey
FROM meta.entity
ORDER BY schema_name, table_name;
```

The extension requires `uuid-ossp`. The installation SQL creates and seeds the `meta` schema directly; use it only in a database dedicated to software that understands the Abris metadata contract.

### Installed Surface

- `meta.entity`, `meta.property`, and `meta.relation` present relations, columns, and relationships through catalog-backed views.
- Projection, menu, page, and extra-attribute tables store application presentation metadata.
- INSTEAD OF and table triggers translate selected writes against metadata views into underlying DDL or metadata-table changes.
- `meta.clean()` removes stale auxiliary rows whose referenced catalog objects no longer exist.

The control file says version `0.0.1` is relocatable, but the SQL hard-codes `meta`, so moving the extension is unsafe. Upstream provides no user guide, supported PostgreSQL matrix, release history, or clear license file. Review the full installation SQL, object ownership, trigger-generated DDL, and privileges before use; do not expose writable metadata views to untrusted roles or treat this as a general-purpose metadata framework.
