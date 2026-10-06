## Usage

Sources:

- [packages/uuid/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/README.md)
- [packages/uuid/sql/pgpm-uuid--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/sql/pgpm-uuid--0.47.0.sql)
- [packages/uuid/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/Makefile)
- [packages/uuid/pgpm-uuid.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/uuid/pgpm-uuid.control)

`pgpm-uuid` installs pseudo-ordered UUID generators and trigger helpers in `uuids`. It can cluster identifiers by time or seed while retaining a random component.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-uuid" CASCADE;
SELECT uuids.pseudo_order_uuid();
SELECT uuids.pseudo_order_seed_uuid('tenant-a');
```

### Operational Boundaries

The generators are not UUID v7 or strict monotonic sequences. `uuids.trigger_set_uuid_seed` and `uuids.trigger_set_uuid_related_field` require the documented trigger arguments and columns. The schema overlaps the legacy `launchql-uuid`; do not co-install those variants.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `pgcrypto`, `plpgsql`, `uuid-ossp`, `hstore`, `pgpm-verify`. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
