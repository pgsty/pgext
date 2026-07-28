## Usage

Sources:

- [Official upstream README](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/README.md)
- [Official extension control file (pg_flatbuffers.control)](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/crates/pg_flatbuffers/pg_flatbuffers.control)
- [Official implementation source](https://github.com/youyuanwu/pg_flatbuffers/blob/a09d4f88c02d0709620abfa0106ceb0c1fd578fa/crates/pg_flatbuffers/src/lib.rs)

`pg_flatbuffers` — pg_flatbuffers: query and convert FlatBuffers payloads in bytea columns. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_flatbuffers;

-- Register a schema produced by `flatc -b --schema orders.fbs`
INSERT INTO flatbuffers_schemas (name, bfbs)
VALUES ('default', pg_read_binary_file('/tmp/orders.bfbs'));

-- Single-value extraction
SELECT flatbuffers_query('myco.orders.Order:customer.email', payload)
FROM   orders_raw
WHERE  id = 42;

-- Vector fan-out as rows (suitable for joins)
SELECT o.id, sku
FROM   orders_raw o,
       LATERAL flatbuffers_query_multi(
         'myco.orders.Order:items[*].sku', o.payload) AS sku;

-- JSON round-trip
SELECT flatbuffers_to_json('myco.orders.Order', payload) -> 'customer'
FROM   orders_raw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `flatbuffers_extension_version()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
