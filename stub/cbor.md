## Usage

Sources:

- [Official README](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/README.md)
- [Version 0.1 SQL objects](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/sql/cbor.sql)
- [CBOR input and output implementation](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/src/cbor_io.c)

`cbor` adds a Concise Binary Object Representation type based on RFC 7049. Use it when a database must store and exchange CBOR bytes while retaining a textual SQL representation; it is not a replacement for `jsonb` query and indexing features.

### Encode and Decode

```sql
CREATE EXTENSION cbor;

SELECT '{"sensor":"alpha","values":[1,2,3]}'::cbor;

WITH encoded AS (
  SELECT cbor_encode('{"ok":true,"count":3}'::cbor) AS payload
)
SELECT cbor_decode(payload)
FROM encoded;
```

`cbor_encode(cbor)` returns the binary representation as `bytea`; `cbor_decode(bytea)` parses binary CBOR. The type's text input and output cover integers, byte and text strings, arrays, maps, tags, floating-point values, and simple values.

### Operators and Indexes

The extension supplies `=`, `<>`, `<`, `<=`, `>`, and `>=`, plus default B-tree and hash operator classes. It also exposes `cbor_contains(cbor, cbor)` and `cbor_contained(cbor, cbor)` functions, but version `0.1.0` does not define containment operators.

```sql
CREATE TABLE messages (id bigint GENERATED ALWAYS AS IDENTITY, body cbor);
CREATE INDEX messages_body_btree ON messages (body);
```

This old release targets the original RFC 7049 rather than the later RFC 8949 specification. Before exchanging data with other implementations, test tags, floating-point values, duplicate map keys, ordering, canonical encoding, and equality semantics. Preserve the original `bytea` if byte-for-byte identity matters, because a decode/encode round trip need not preserve an external encoder's exact representation.
