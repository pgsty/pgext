## Usage

Sources:

- [README.md](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/README.md)
- [uuid_v1--0.1.sql](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/uuid_v1--0.1.sql)
- [uuid_v1.control](https://github.com/ancoron/pg-uuid-v1/blob/577b1db78f16351199613d88e3e847c08b45118e/uuid_v1.control)

`uuid_v1` stores version-1 UUIDs in a type whose ordering favors the timestamp component. It supplies comparison operators, casts to/from uuid, B-tree support and component extraction functions.

### Core Workflow

```sql
CREATE EXTENSION uuid_v1;
SELECT uuid_v1_get_timestamp('b647e96b-862d-11e9-ae2b-db6f0f573554'::uuid_v1);
SELECT uuid_v1_get_epoch('b647e96b-862d-11e9-ae2b-db6f0f573554'::uuid_v1);
```

### Operational Boundaries

`uuid_v1_get_timestamp` returns timestamptz, `uuid_v1_get_epoch` returns an immutable epoch value, and `uuid_v1_get_clockseq` and `uuid_v1_get_node` expose the remaining components. Timestamp comparison operators compare full precision; a second-rounded timestamp need not equal a UUID with fractional seconds.

This extension consumes version-1 UUIDs rather than generating them. Installation requires a superuser and its C library; no preload is needed. No current PostgreSQL-major matrix is declared. Choose the epoch helper for immutable expression indexes as documented upstream.
