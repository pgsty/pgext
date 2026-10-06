## Usage

Sources:

- [packages/object-store/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/README.md)
- [packages/object-store/sql/object-store--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/sql/object-store--0.47.0.sql)
- [packages/object-store/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/Makefile)
- [packages/object-store/object-store.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/object-store/object-store.control)

`object-store` stores versioned JSONB objects and child references. Objects are content-addressed after freezing, enabling structural sharing between versions.

### Core Workflow

```sql
CREATE EXTENSION "object-store" CASCADE;
SELECT * FROM object_store_public.object LIMIT 10;
```

### Operational Boundaries

`object_store_public.get_node_at_path`, `object_store_public.insert_node_at_path`, `object_store_public.update_node_at_path` and `object_store_public.freeze_objects` implement tree navigation and copy-on-write operations. Frozen rows cannot be edited through ordinary operations. Retain referenced child objects with their roots; immutability and content hashes do not replace permissions or backups.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`, `pgcrypto`, `uuid-ossp`, `pgpm-verify`. Its SQL expects the platform roles `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
