## Usage

Sources:

- [Official README](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/README.md)
- [Rust SQL API implementation](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/src/lib.rs)
- [Build and PostgreSQL feature matrix](https://github.com/sqids/sqids-postgresql/blob/709167f4c3c97f7af096a94c854e8e5071ddeaf5/Cargo.toml)

`pg_sqids` encodes one or more non-negative integers into short, URL-safe Sqids and decodes them back in PostgreSQL. It is convenient for public-facing identifiers and compact links, but it is reversible obfuscation, not encryption or access control.

### Core Workflow

```sql
CREATE EXTENSION pg_sqids;

SELECT sqids_encode(1, 2, 3);
-- 86Rf07

SELECT sqids_decode('86Rf07');
-- {1,2,3}

SELECT sqids_encode(10::smallint, 1, 2, 3);
-- 86Rf07xd4z
```

The first `smallint` overload sets a minimum output length. Other overloads accept a custom alphabet, a `text[]` blocklist, or combinations of alphabet, minimum length, and blocklist; decoding must use the same options.

```sql
SELECT sqids_encode(
    'k3G7QAe51FCsPW92uEOyq4Bg6Sp8YzVTmnU0liwDdHXLajZrfxNhobJIRcMvKt',
    ARRAY['XRKUdQ'],
    1, 2, 3
);
```

### Input and Identity Rules

- Negative numbers are rejected. The implementation maps null variadic numeric arguments to zero, so reject nulls in application input when distinct identity is required.
- Minimum length must be between 0 and 255.
- Multiple different Sqid strings can decode to the same numeric sequence. If canonical text matters, decode, re-encode with the same options, and compare the result with the supplied string.
- A blocklist changes generated output; it is not a content filter for arbitrary decoded input.

Sqids reveal the encoded numbers and their count to anyone who knows the algorithm and options. Do not encode secrets, authorization scope, or sensitive user counts, and continue to enforce database authorization on the decoded key.

### Compatibility

Release 0.1.0 uses `pgrx` 0.18.1 and the upstream Cargo features cover PostgreSQL 13 through 18, with PostgreSQL 13 as the default build feature. Install a binary built for the exact server major. The control file marks the extension trusted but also requests superuser installation; actual installation policy depends on the packaged control file and PostgreSQL version.
