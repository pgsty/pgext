

## Usage

> [pg_idkit: multi-tool for generating new/niche universally unique identifiers](https://github.com/VADOSWARE/pg_idkit)

```sql
CREATE EXTENSION pg_idkit;
SELECT idkit_uuidv7_generate();
```

### Available Functions

| Methodology | Function | Description |
|---|---|---|
| UUID v6 | `idkit_uuidv6_generate()` | UUID v6 (RFC 4122) |
| | `idkit_uuidv6_generate_uuid()` | UUID v6 as native UUID type |
| | `idkit_uuidv6_extract_timestamptz(TEXT)` | Extract timestamp from UUID v6 |
| UUID v7 | `idkit_uuidv7_generate()` | UUID v7 (RFC 4122) |
| | `idkit_uuidv7_generate_uuid()` | UUID v7 as native UUID type |
| | `idkit_uuidv7_extract_timestamptz(TEXT)` | Extract timestamp from UUID v7 |
| NanoID | `idkit_nanoid_generate()` | NanoID |
| | `idkit_nanoid_custom_generate_text()` | NanoID with custom length and alphabet |
| KSUID | `idkit_ksuid_generate()` | K-Sortable UID |
| | `idkit_ksuid_extract_timestamptz(TEXT)` | Extract timestamp from KSUID |
| | `idkit_ksuidms_generate()` | KSUID with millisecond precision |
| | `idkit_ksuidms_extract_timestamptz(TEXT)` | Extract timestamp from KSUID-ms |
| ULID | `idkit_ulid_generate()` | Universally Unique Lexicographically Sortable ID |
| | `idkit_ulid_extract_timestamptz(TEXT)` | Extract timestamp from ULID |
| Timeflake | `idkit_timeflake_generate()` | Snowflake + Instagram ID + Firebase PushID |
| | `idkit_timeflake_extract_timestamptz(TEXT)` | Extract timestamp from Timeflake |
| PushID | `idkit_pushid_generate()` | Google Firebase PushID |
| XID | `idkit_xid_generate()` | XID |
| | `idkit_xid_extract_timestamptz(TEXT)` | Extract timestamp from XID |
| CUID | `idkit_cuid_generate()` | CUID (deprecated) |
| | `idkit_cuid_extract_timestamptz(TEXT)` | Extract timestamp from CUID |
| CUID2 | `idkit_cuid2_generate()` | CUID2 |
| | `idkit_cuid2_generate_with_len(length)` | CUID2 with custom length |
| TypeID | `idkit_typeid_generate(TEXT)` | TypeID with prefix and UUIDv7 |
| | `idkit_typeid_generate_text(TEXT)` | TypeID as text |
| | `idkit_typeid_from_uuid_v7(TEXT, TEXT)` | TypeID from a given UUID v7 |
| | `idkit_typeid_extract_timestamptz(TEXT)` | Extract timestamp from TypeID |

### Examples

```sql
-- Generate different ID types
SELECT idkit_uuidv7_generate();           -- 018c106f-9304-79bb-b5be-4483b92b036c
SELECT idkit_nanoid_generate();            -- A8jFA0r3NC6FdalR4LEJ0
SELECT idkit_ksuid_generate();             -- 2HMQIBkTJmEN11JI7tvSTMwfYI3
SELECT idkit_ulid_generate();              -- 01HPYV2X17GM5SQP22M3DVFZY3
SELECT idkit_cuid2_generate();             -- clrjx3bwh0000fj3x4c2y1z0s

-- Extract timestamp
SELECT idkit_uuidv7_extract_timestamptz('018c106f-9304-79bb-b5be-4483b92b036c');
```
