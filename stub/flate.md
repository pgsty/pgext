## Usage

Sources:

- [Official README](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/README.markdown)
- [Official extension SQL](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/flate--1.0.1.sql)
- [Official C implementation](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/flate.c)

`flate` exposes zlib raw-DEFLATE compression and decompression for PostgreSQL `bytea` values. Version `1.0.1` is suitable when an application specifically needs raw streams; its output does not contain a zlib or gzip wrapper.

### Core Workflow

`deflate(bytea, bytea, integer)` accepts optional preset-dictionary and compression-level arguments. `inflate(bytea, bytea)` must receive the same dictionary when one was used for compression:

```sql
CREATE EXTENSION flate;

SELECT deflate(
  convert_to('hello hello hello hello', 'UTF8'),
  convert_to('hello', 'UTF8'),
  9
);

SELECT convert_from(
  inflate(E'\\xcb00110a182400'::bytea, convert_to('hello', 'UTF8')),
  'UTF8'
);
```

A NULL level selects zlib's default; explicit levels follow zlib's 0 through 9 range. Both functions are declared immutable.

### Operational Caveats

The dictionary is not embedded in the compressed value. A missing or different dictionary, damaged input, or a stream with the wrong wrapper format causes decompression to fail or produce unusable data.

Compressed input can expand to a much larger result. Limit untrusted input sizes and account for backend memory and statement timeout before calling `inflate(bytea, bytea)`. Store the encoding, wrapper choice, and dictionary version alongside data so another client can reproduce the operation.
