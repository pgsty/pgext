## Usage

Sources:

- [Official README](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/README.md)
- [Extension control file](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/uuid47.control)
- [Extension SQL](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/sql/uuid47--1.0.sql)
- [PostgreSQL implementation](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/src/uuid47_pg.c)

`uuid47` stores time-ordered UUIDv7 bytes while presenting a deterministic UUIDv4-looking facade through text output and casts. A 128-bit key drives a reversible SipHash-derived mask of the timestamp bits; this is identifier obfuscation with key-management requirements, not encryption of row data.

### Core Workflow

Set a stable key before any text input, output, cast, dump, or client exchange. Provision it through protected session or role configuration rather than embedding it in application SQL.

```sql
CREATE EXTENSION uuid47;
SET uuid47.key = '0011223344556677:8899aabbccddeeff';

CREATE TABLE event (
  event_id uuid47 PRIMARY KEY DEFAULT uuid47_generate_monotonic(),
  payload jsonb NOT NULL
);

INSERT INTO event(payload) VALUES ('{"kind":"created"}');

SELECT event_id,
       uuid47_timestamp(event_id),
       uuid47_as_v7(event_id),
       uuid47_key_fingerprint()
FROM event;
```

`uuid47_generate()` creates random UUIDv7 suffixes; `uuid47_generate_monotonic()` adds per-backend monotonic ordering within a millisecond; `uuid47_generate_at(timestamptz)` uses a supplied time. `uuid47_timestamp`, `uuid47_as_v7`, and `uuid47_explain` inspect stored UUIDv7 state. Assignment casts between `uuid47` and `uuid` use the session key, while explicit-key variants `uuid47_to_uuid_with_key` and `uuid_to_uuid47_with_key` accept a 16-byte key.

The type supplies comparison operators, default B-tree and hash operator classes, and the nondefault `uuid47_brin_minmax_multi_ops`. Comparisons and indexes operate on internal UUIDv7 bytes, providing time-oriented ordering independent of the displayed facade.

### Key and Interchange Caveats

`uuid47.key` is `USERSET`, so every session can choose a different key. The same stored value then displays as a different UUIDv4-looking identifier, and a facade created under one key decodes to the wrong UUIDv7 under another. Standardize the key across trusted clients, keep it out of logs, record a key identifier out of band, and test backup/restore with the intended key. A plain text dump invokes keyed output; losing the key changes the external identity surface.

Key rotation requires a compatibility plan for already issued facade IDs; changing the setting alone does not preserve their decoding. The mask hides direct timestamp appearance but leaves the random bits unchanged and is not authorization or confidentiality. Use a truly independent external ID when consumers must retain identifiers across key rotation or when correlation risk is unacceptable. No preload is required, but binaries and BRIN support must match the target PostgreSQL major.
