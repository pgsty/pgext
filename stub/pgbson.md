## Usage

Sources:

- [postgresbson README at the 2.0.4 revision](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/README.md)
- [META.json version 2.0.4](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/META.json)
- [pgbson control file](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson.control)
- [Version 2.0 SQL API](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson--2.0.sql)

pgbson adds a BSON data type, typed path accessors, JSON-style operators, casts, and expression-index support. Use it when binary BSON must be stored without first converting every value to JSONB, especially when BSON type fidelity or byte-level round trips matter.

The distribution release is 2.0.4 while the extension control and SQL API version remain 2.0.

### Create the Extension

    CREATE EXTENSION pgbson;

The native module uses libbson. Install a package built against compatible PostgreSQL and libbson versions.

### Store and Validate BSON

The bytea-to-bson cast validates input when a value is written. Version 2.0.4 documents that reads can then assume the stored BSON is valid. Do not bypass the type's input or cast path with unsafe low-level writes.

### Extract Values

Typed dot-path accessors avoid materializing every intermediate object:

    SELECT bson_get_datetime(payload, 'msg.header.event.ts'),
           bson_get_string(payload, 'data.customer.name')
    FROM events;

Use bson_get_bson for a subdocument:

    SELECT bson_get_bson(payload, 'msg.header.event')
    FROM events;

JSON-style navigation is also available:

    SELECT payload->'msg'->'header'->'event'->>'ts'
    FROM events;

### Function and Operator Index

- bson_get_string, bson_get_int32, bson_get_int64, bson_get_double, bson_get_decimal: typed scalar accessors.
- bson_get_datetime, bson_get_binary, bson_get_boolean: accessors for additional BSON types.
- bson_get_bson: return an embedded BSON document.
- bson_get_jsonb_array: convert an array endpoint to a PostgreSQL jsonb array.
- -> and ->>: navigate values with JSON-like syntax.
- bson casts to json and jsonb: expose Extended JSON for PostgreSQL JSON processing.
- bson and bytea casts: preserve the BSON binary representation.

### Index and Interoperate

Create expression indexes on frequently queried paths:

    CREATE INDEX events_customer_id_idx
    ON events (bson_get_string(payload, 'data.customer.id'));

Cast a subdocument to jsonb when PostgreSQL's JSON operators are more convenient:

    SELECT bson_get_bson(payload, 'msg.header')::jsonb ? 'event'
    FROM events;

### Caveats

- A typed getter returns useful data only when the endpoint has the expected BSON type. Make type expectations explicit in ingestion code.
- bson_get_bson returns NULL for scalar endpoints because a scalar is not a BSON document.
- Dot-path accessors are generally preferable to long operator chains for repeated extraction because they avoid intermediate BSON values.
- BSON and JSONB have different type and ordering semantics. A cast can be useful but is not a lossless replacement for every BSON workflow.
