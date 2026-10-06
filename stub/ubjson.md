## Usage

Sources:

- [README.rst](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/README.rst)
- [pg_ubjson/ubjson--0.0.1.sql](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/ubjson--0.0.1.sql)
- [pg_ubjson/Makefile](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/Makefile)
- [pg_ubjson/ubjson.control](https://github.com/dvarrazzo/jsonb_parser/blob/78128f3094e976a55adb04c9fe287b7ac2045e60/pg_ubjson/ubjson.control)

`ubjson` is the PostgreSQL component of the jsonb format parser project. It uses JSONB-compatible input/output and storage with its own binary send/receive functions for client transfer.

### Core Workflow

```sql
CREATE EXTENSION ubjson;
SELECT '{"answer":42}'::jsonb::ubjson;
SELECT ('{"answer":42}'::jsonb::ubjson)::jsonb;
```

### Operational Boundaries

The type shares JSONB’s textual representation and supports assignment casts in both directions without copying. Its binary representation is a separate protocol; clients must use a matching decoder rather than treating it as ordinary PostgreSQL JSONB wire data. This does not add a general-purpose JSON query language.

Creating the C extension requires a superuser. No preload is specified. The source is experimental and supplies no PostgreSQL-major support matrix; evaluate binary compatibility before upgrading either client or server.
