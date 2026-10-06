## Usage

Sources:

- [packages/metaschema-modules/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/README.md)
- [packages/metaschema-modules/sql/metaschema-modules--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/sql/metaschema-modules--0.47.0.sql)
- [packages/metaschema-modules/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/Makefile)
- [packages/metaschema-modules/metaschema-modules.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/metaschema-modules/metaschema-modules.control)

`metaschema-modules` extends the metaschema model with configuration records for reusable application modules, including accounts, memberships, messaging and authentication integrations.

### Core Workflow

```sql
CREATE EXTENSION "metaschema-modules" CASCADE;
SELECT * FROM metaschema_modules_public.connected_accounts_module LIMIT 10;
```

### Operational Boundaries

These tables describe module configuration and references to the base metaschema. They do not start an authentication provider or network service. Apply changes through the matching platform tooling so that referenced schemas, fields, policies and generated DDL remain consistent; restrict writes to administrators.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`, `uuid-ossp`, `metaschema-schema`, `pgpm-verify`. Its SQL expects the platform roles `administrator`, `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
