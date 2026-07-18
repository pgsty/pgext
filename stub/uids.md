## Usage

Sources:

- [Pinned upstream README](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/README.md)
- [Pinned Cargo manifest](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/Cargo.toml)
- [Pinned Rust function sources](https://github.com/spa5k/uids-postgres/tree/4bfbf05a6a9a0771dc66000fb3bb19974e439207/src)
- [Pinned extension control file](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/uids.control)

uids version 0.0.1 is a pgrx extension that generates UUIDv6, UUIDv7, NanoID, KSUID, ULID, Timeflake, PushID, CUID2, and TypeID values. Most functions return text; selected functions return bytea or PostgreSQL uuid. Each family has different length, ordering, timestamp exposure, entropy, and interoperability semantics.

### Generate representative identifiers

```sql
CREATE EXTENSION uids;

SELECT generate_uuidv6_uuid();
SELECT generate_uuidv7();
SELECT generate_ulid();
SELECT generate_ksuid();
SELECT generate_nanoid_length_c(10, '0123456789abcdef');
SELECT generate_typeid('user');
SELECT check_typeid('user', generate_typeid('user'));
SELECT generate_cuid2();
SELECT generate_timeflake_uuid();
SELECT generate_pushid();
```

generate_uuidv6_uuid returns the native uuid type, but generate_uuidv7 returns text. Choose and enforce one SQL type and format per column; changing generator families later is a data migration, not a transparent implementation detail.

### Validation, ordering, and build limits

generate_uuidv7_from_string and parse_uuidv7 parse generic UUID syntax but do not verify version 7; the upstream test itself uses a version-4 UUID. check_typeid parses with expect, so malformed input raises an error instead of returning false. Similar parsing and random-generation paths use unwrap or panic. Treat validation functions as exception-producing and wrap them when processing untrusted input.

NanoID length is accepted as signed integer and cast to an unsigned machine size. A negative or unreasonable length can become a huge allocation request. Custom alphabets can also be invalid. Enforce small positive length limits, a nonempty unique alphabet, and input byte limits before calling any configurable function.

Time-sortable identifiers can reveal creation time and their text order depends on collation. If order is a requirement, test concurrent generation and store sortable text under a bytewise collation or use a native type with documented comparison behavior. Uniqueness is probabilistic for most families; none of these values should automatically be treated as an authentication secret or proof of authorization.

The control file requires a superuser and is non-relocatable. The pinned Cargo manifest uses pgrx 0.11.4, defaults to PostgreSQL 16, and has features for PostgreSQL 11 through 16; native artifacts must match the exact PostgreSQL major and platform. The GitHub v1 tag still builds extension version 0.0.1, and the repository declares no license. Verify the installed extversion, function return types, randomness failure behavior, collation, dump/restore, and collision handling before adoption.
