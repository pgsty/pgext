## Usage

Sources:

- [pgbson 2.1.0 README](https://api.pgxn.org/src/bson/bson-2.1.0/README.md)
- [pgbson 2.1 control file](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson.control)
- [pgbson 2.1 SQL API](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson--2.1.sql)

`pgbson` adds a BSON data type, typed dot-path accessors, JSON-style navigation, casts, comparison operators, and btree/hash indexing. The PGXN distribution release is `2.1.0`, while the SQL extension version is `2.1`. Use BSON when binary round-trip fidelity or BSON-specific scalar types matter; use `jsonb` when PostgreSQL-native JSON indexing is the primary requirement.

### Install and Store BSON

```sql
CREATE EXTENSION pgbson;
SELECT pgbson_version();

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload bson NOT NULL
);

INSERT INTO events (payload)
VALUES ('{"user":{"name":"Ada"},"attempt":3}'::jsonb::bson);
```

The native module depends on `libbson`. The implicit `bytea`-to-`bson` cast validates BSON input, while the reverse cast preserves the binary representation.

### Extract Values

Typed accessors avoid materializing each intermediate document:

```sql
SELECT bson_get_string(payload, 'user.name'),
       bson_get_int32(payload, 'attempt')
FROM events;
```

Other typed getters cover 64-bit integers, doubles, decimals, datetimes, binary values, booleans, embedded BSON documents, and JSONB arrays. A missing path or a type mismatch returns `NULL`, so validate the expected BSON schema at ingestion when those cases must be distinguished.

Version 2.1 adds a type-agnostic terminal extractor:

```sql
SELECT bson_get_value(payload, 'user.name')
FROM events;
-- { "_" : "Ada" }
```

`bson_get_value` always wraps the selected scalar, array, or document under the key `_`. Remove exactly that one wrapper in the caller. It intentionally has no chainable `->` equivalent.

### Navigate, Compare, and Index

```sql
SELECT payload->'user'->>'name'
FROM events;

CREATE INDEX events_user_name_idx
ON events (bson_get_string(payload, 'user.name'));

CREATE INDEX events_payload_btree_idx ON events (payload);
CREATE INDEX events_payload_hash_idx ON events USING hash (payload);
```

Version 2.1 provides logical comparison operators `=`, `<>`, `<`, `<=`, `>`, and `>=`; `==` and `<<>>` perform binary equality and inequality. The default btree operator class uses logical BSON comparison, while the hash operator class uses binary equality. Choose intentionally when field order or byte identity matters.

### Upgrade and Caveats

```sql
ALTER EXTENSION pgbson UPDATE TO '2.1';
```

- Installing a 2.1 shared library does not update an existing 2.0 extension's SQL objects; run the extension update after installing the files.
- The 2.1 shared library fixes a backend crash when `bson_get_bson()` or `->` resolves to a scalar endpoint. Earlier binaries should be replaced even when an application does not yet use the new 2.1 SQL function.
- BSON-to-JSON/JSONB casts use Extended JSON. BSON and JSONB have different type, equality, and ordering semantics, so conversion is not lossless for every workflow.
- In 2.1, `->>` on a BSON datetime includes the trailing `Z`; `bson_get_datetime()` is unchanged. Check clients that compare the old text form.
- BSON top-level values are documents, not bare arrays or scalars. `bson_get_value` uses its `_` wrapper to return any nested shape within that restriction.
