## Usage

Sources:

- [Official README](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/README.md)
- [Extension SQL](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag--1.0.sql)
- [BLAKE2 implementation](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag.c)

`pg_etag` generates hexadecimal BLAKE2b digests for `text`, `bytea`, ordered result streams, and PostgreSQL large objects. Use those digests as application cache validators only after defining a stable serialization and row order.

### Single Values and Rows

```sql
CREATE EXTENSION pg_etag;

SELECT etag('canonical payload');
SELECT etag(convert_to('canonical payload', 'UTF8'));

SELECT id,
       etag(concat_ws(E'\x1f', id::text, updated_at::text, payload)) AS row_etag
FROM documents;

SELECT blake2('short digest', 32);
```

`etag(text)` and `etag(bytea)` return a 64-byte BLAKE2b digest as 128 lowercase hexadecimal characters. `blake2(text, integer)` and `blake2(bytea, integer)` let the caller select a digest length from 1 through 64 bytes; the implementation clamps values outside that interval.

### Result-Set ETags

```sql
SELECT etag_agg(row_etag ORDER BY id) AS result_etag
FROM (
    SELECT id,
           etag(concat_ws(E'\x1f', id::text, updated_at::text, payload)) AS row_etag
    FROM documents
    WHERE tenant_id = 42
) AS ordered_rows;
```

`etag_agg(text)` and `etag_agg(bytea)` hash the input bytes in aggregate order and ignore null inputs. Always put a deterministic `ORDER BY` inside the aggregate. The state function concatenates input bytes without adding a delimiter, so encode field and row boundaries in the hashed representation; otherwise different segmentations can produce the same byte stream.

### Large Objects and Caveats

The overload `public.etag(oid)` hashes `pg_largeobject.data` pages ordered by page number. It requires access to the large object and is hard-coded into the `public` schema even though the control file says the extension is relocatable.

An ETag is only as stable as its serialization. Text formatting of timestamps, numerics, JSON, collations, or encodings can change across settings or versions. Hash canonical bytes, include every field that defines the representation, and remember that BLAKE2 collision resistance does not provide authentication or authorization. The extension depends on `libb2`.
