## Usage

Sources:

- [packages/metaschema-schema/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/README.md)
- [packages/metaschema-schema/sql/metaschema-schema--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/sql/metaschema-schema--0.47.0.sql)
- [packages/metaschema-schema/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/Makefile)
- [packages/metaschema-schema/metaschema-schema.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-schema/metaschema-schema.control)

`metaschema-schema` stores a declarative database model in metadata tables. It describes databases, schemas, tables, fields, constraints, indexes and grants for the surrounding platform.

### Core Workflow

```sql
CREATE EXTENSION "metaschema-schema" CASCADE;
SELECT * FROM metaschema_public.database LIMIT 10;
SELECT * FROM metaschema_public.schema LIMIT 10;
```

### Operational Boundaries

The fixed `metaschema_public` and `metaschema_private` schemas include validation helpers and change/job integration. A metadata row is not proof that physical DDL has been applied: the platform’s reconciler must process the model. Review generated identifiers, privileges, version alignment and migration effects before changing these records.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `citext`, `hstore`, `pgpm-inflection`, `pgpm-database-jobs`, `pgpm-types`, `pgcrypto`, `plpgsql`, `postgis`, `uuid-ossp`, `pgpm-verify`. Its SQL expects the platform roles `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
