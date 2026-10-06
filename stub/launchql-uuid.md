## Usage

Sources:

- [packages/uuid/readme.md](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/readme.md)
- [packages/uuid/sql/launchql-uuid--0.4.5.sql](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/sql/launchql-uuid--0.4.5.sql)
- [packages/uuid/Makefile](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/Makefile)
- [packages/uuid/launchql-uuid.control](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/uuid/launchql-uuid.control)

`launchql-uuid` provides pseudo-ordered and seed-prefixed UUID generators plus row-trigger helpers in `uuids`.

### Core Workflow

```sql
CREATE EXTENSION "launchql-uuid" CASCADE;
SELECT uuids.pseudo_order_uuid();
SELECT uuids.pseudo_order_seed_uuid('tenant-a');
```

### Operational Boundaries

`uuids.trigger_set_uuid_seed` and `uuids.trigger_set_uuid_related_field` require the columns and arguments in the versioned SQL. Generated UUIDs are neither UUID v7 nor a strict monotonic sequence. This historical distribution overlaps the `pgpm-uuid` schema and must not be installed alongside it.

Install its declared dependencies first: `pgcrypto`, `plpgsql`, `uuid-ossp`, `hstore`. This is SQL/PLpgSQL code with no own shared library or preload. The control permits non-superuser installation but is not marked trusted; dependency and schema privileges still apply.
