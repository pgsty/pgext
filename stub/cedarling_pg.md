## Usage

Sources:

- [Official upstream README](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/README.md)
- [Official extension control file (cedarling_pg.control)](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/cedarling_pg/cedarling_pg.control)
- [Official extension SQL (cedarling_pg--0.1.0.sql)](https://github.com/janssenproject/jans/blob/2d578f00e5408a213d70402f10205b09467d16ed/jans-cedarling/cedarling_pg/sql/cedarling_pg--0.1.0.sql)

`cedarling_pg` — PostgreSQL integration for Cedarling authorization decisions, including JWT-aware checks and RLS-oriented helpers. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION cedarling_pg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `cedarling.entity_map` is a table installed or managed by the extension.
- `cedarling.mask_rules` is a table installed or managed by the extension.
- `cedarling.policy_history` is a table installed or managed by the extension.
- `cedarling.policy_versions` is a table installed or managed by the extension.
- `cedarling` is a schema created by the extension.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
