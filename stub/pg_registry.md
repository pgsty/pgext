## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [Official pg_registry README](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_registry/README.md)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_registry/pg_registry.control)
- [SQL API and storage definitions](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_registry/src/lib.rs)

`pg_registry` version `0.3.0` stores versioned JSON Schemas in PostgreSQL, validates JSONB values, generates schemas from tables, and records schema-to-topic or table-to-topic bindings for companion streaming integrations.

### Core Workflow

```sql
CREATE EXTENSION pg_registry;

SELECT pgregistry.register_schema(
    'orders-value',
    '{"type":"object","properties":{"order_id":{"type":"integer"}},"required":["order_id"]}'::jsonb,
    'Order event schema'
) AS schema_id;

SELECT pgregistry.validate(
    1,
    '{"order_id":123}'::jsonb
);

SELECT pgregistry.bind_schema_to_topic(
    'orders', 1, 'value', 'STRICT'
);
```

Important functions are `register_schema`, `register_schema_from_table`, `generate_schema_from_table`, `get_schema`, `get_latest_schema`, `drop_schema`, `validate`, and `validate_for_topic`. Topic functions manage key/value schema bindings; table functions can create a table from a schema and store stream or upsert binding metadata.

### Operational Notes

Objects live in the non-relocatable `pgregistry` schema, and the control file requires superuser installation. `STRICT` and `LOG` are stored validation modes, but enforcement during Kafka traffic depends on the companion integration consuming these bindings; `pg_registry` alone does not start a Kafka client. Generated schemas and tables cover the documented PostgreSQL-to-JSON type mapping and may not preserve every domain, constraint, generated expression, or application invariant. Review generated DDL and access to schema-management functions before production use.

### Version 0.3.0 Boundary

This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
