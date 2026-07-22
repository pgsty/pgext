## Usage

Sources:

- [Official README](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/README.md)
- [Official extension SQL](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/pgcodec7--1.0.sql)
- [Official C implementation](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/pgcodec7.c)

`pgcodec7` 1.0 converts `bytea` to a set of 6-bit or 7-bit text chunks and reconstructs the original bytes from a `text[]`. It reduces hexadecimal expansion for specialized transports, but it is neither compression nor encryption.

### Core Workflow

The SQL API differs from the old README example: `pgencode7(integer, bytea)` accepts bytes, not a filename, and `pgdecode7(integer, text[])` requires an array of chunks, not one text value. Preserve chunk order during aggregation.

```sql
CREATE EXTENSION pgcodec7;

SELECT convert_from(
           pgdecode7(6, array_agg(chunk ORDER BY ord)),
           'UTF8'
       ) AS decoded
FROM pgencode7(6, convert_to('hello', 'UTF8'))
     WITH ORDINALITY AS encoded(chunk, ord);
```

`pgencode7` returns multiple rows because each output chunk is limited to roughly 1024 characters. Store an explicit ordinal whenever chunks may pass through a table, queue, or transport that does not preserve row order.

### Functions

- `pgencode7(i_encodelen integer, by_data bytea) RETURNS SETOF text` accepts only mode 6 or 7.
- `pgdecode7(i_decodelen integer, t_encodes text[]) RETURNS bytea` concatenates and decodes an ordered chunk array.

Both functions are installed in `public` and execution is granted to `PUBLIC` by the extension SQL.

### Encoding and Safety Boundaries

6-bit output uses a transport-friendlier alphabet. The 7-bit alphabet includes bytes in the 0x80–0xAF range, which are not standalone valid UTF-8 characters; many text clients, JSON encoders, collations, and gateways may reject or alter them. Use 7-bit mode only across a byte-preserving path that has been tested end to end.

The decoder’s C implementation does not safely reject every byte outside its alphabet before arithmetic. Do not pass untrusted or mutated chunks directly to `pgdecode7`; validate the mode, alphabet, chunk count, total size, and ordering first. Malformed or transport-damaged input can otherwise cause errors or backend instability.

The upstream project was built against PostgreSQL 11 and has been inactive since 2019. Test build compatibility and round trips on the exact target PostgreSQL release before relying on it.
