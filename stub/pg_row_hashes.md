## Usage

Sources:

- [Official README](https://github.com/ibotty/pg-row-hashes/blob/main/README.md)
- [Official version 0.3.2 implementation](https://github.com/ibotty/pg-row-hashes/blob/main/src/lib.rs)
- [Official UUID XOR aggregate](https://github.com/ibotty/pg-row-hashes/blob/main/src/xor_agg.rs)
- [Official control file](https://github.com/ibotty/pg-row-hashes/blob/main/pg_row_hashes.control)

`pg_row_hashes` provides deterministic, non-cryptographic fingerprints for ordered identifiers and map-like row checksums. Version `0.3.2` offers FarmHash 128-bit results as `uuid`, SeaHash 64-bit results as `bigint`, an underscore-joined MD5 identifier helper, and a UUID XOR aggregate.

### Core Workflow

Use `id_*` functions when argument order is part of the identifier, and `checksum_*` functions for key/value pairs:

```sql
CREATE EXTENSION pg_row_hashes;

SELECT id_farmhash('tenant-7', 'order-42');
SELECT id_seahash('tenant-7', 'order-42');

SELECT checksum_farmhash(
  'status', 'paid',
  'amount', '19.95'
);

SELECT checksum_farmhash_extendable(
  'status', 'paid',
  'optional_note', NULL
);
```

The extendable variants skip pairs whose value is `NULL`, so adding a new nullable field does not change old-row fingerprints until the field has a value. Non-extendable variants retain the key's presence.

### API

- `id_farmhash(VARIADIC text[])` / `id_seahash(VARIADIC text[])`: ordered multi-part fingerprints.
- `id_farmhash_bytea(bytea)` / `id_seahash_bytea(bytea)`: fingerprint raw bytes.
- `id_underscore_md5(VARIADIC text[])`: MD5 of arguments joined by `_`, returned as `uuid`.
- `checksum_farmhash` / `checksum_seahash`: order-normalized key/value checksums that retain null-valued keys.
- `checksum_farmhash_extendable` / `checksum_seahash_extendable`: skip null-valued pairs.
- `bit_xor(uuid)`: parallel-safe UUID XOR aggregate.

### Caveats

FarmHash, SeaHash, MD5, and XOR here are fingerprinting tools, not authentication, tamper evidence, or collision-resistant cryptography. Store the source/business key alongside any fingerprint and handle possible collisions. XOR also cancels duplicate equal UUIDs, so it cannot prove multiset contents by itself.

Checksum arguments are text key/value pairs and must have even arity. Callers must define stable type-to-text, locale, timezone, numeric, and JSON canonicalization rules; a formatting change changes the fingerprint. Version `0.3.2` builds for PostgreSQL 17/18, so verify the package feature matches the target server.
