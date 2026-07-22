## Usage

Sources:

- [Pinned official README](https://github.com/aquameta/pg_catalog_get_defs/blob/93c358ae1851913960861bb010923a245e538b6e/README.md)
- [Pinned extension SQL](https://github.com/aquameta/pg_catalog_get_defs/blob/93c358ae1851913960861bb010923a245e538b6e/pg_get_typedef.sql)

`pg_catalog_get_defs` adds SQL functions that reconstruct a `CREATE TYPE` statement from PostgreSQL catalogs. It fills a type-definition introspection gap for tooling and inspection; it does not execute the returned DDL or reconstruct dependent objects.

### Core Workflow

Create the extension and pass a type OID to `get_typedef(oid)`:

```sql
CREATE EXTENSION pg_catalog_get_defs;

CREATE TYPE deployment_state AS ENUM ('pending', 'ready', 'failed');

SELECT get_typedef('deployment_state'::regtype::oid);
```

The result is text similar to `CREATE TYPE deployment_state AS ENUM (...)`. Treat it as generated output to review, qualify, and order with dependency DDL before replaying elsewhere.

### Type-Specific Functions

- `get_typedef_enum(oid)` reconstructs enum labels in sort order.
- `get_typedef_composite(oid)` reconstructs live composite attributes and non-default collations.
- `get_typedef_domain(oid)` reconstructs the base type and default expression.
- `get_typedef_range(oid)` reconstructs subtype, optional operator class/collation, canonical function, and subtype-difference function.
- `get_typedef_base(oid)` reconstructs base-type storage and I/O properties.
- `get_typedef(oid)` dispatches by `pg_type.typtype`; an undefined pseudo-type becomes shell-type DDL, while a defined pseudo-type raises an error.

### Boundaries

The generated statement is based on catalog state visible to the caller. It is not a complete dump: comments, privileges, owners, extension membership, domain constraints, array types, casts, operators, dependent functions, and surrounding schema setup are not reconstructed as a migration bundle.

Names are rendered through `regtype` or catalog formatting and may depend on `search_path`. Schema-qualify and validate generated DDL before applying it to another database. The project is small and its SQL follows older catalog layouts, so test every supported type kind on the target PostgreSQL release. No preload or restart is required.
