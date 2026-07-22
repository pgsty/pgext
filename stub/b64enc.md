## Usage

Sources:

- [Official type definition](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/sql/b64enc.sql)
- [Official Rust implementation](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/src/lib.rs)
- [Official extension control file](https://github.com/adunstan/pg_b64enc_rust/blob/8b530363ecdd1b0985c29d63d76b192db563a38b/b64enc.control)

`b64enc` is an experimental fixed-width base type that decodes a URL-safe Base64 string into at most eight bytes and stores the result as one pass-by-value 64-bit datum. It is not a general replacement for PostgreSQL's `bytea` plus `encode` and `decode` functions.

### Core Workflow

Create the extension and cast a URL-safe Base64 literal to the type:

```sql
CREATE EXTENSION b64enc;

SELECT 'AQ=='::b64enc;
SELECT 'AQIDBAUGBwg='::b64enc;
```

The output function always serializes all eight stored bytes, so a short input is emitted in a canonical eight-byte form with leading zero bytes. The type defines only text input/output; the binary send/receive functions are commented out in the reviewed SQL.

### Limits and Failure Behavior

- Decoded input longer than eight bytes is rejected. There are no documented comparison operators, casts, arithmetic, or operator classes, so the type has very little useful SQL behavior beyond storage and display.
- Distinct textual encodings that decode to byte sequences differing only by leading zero bytes collapse to the same 64-bit value and output. Do not use original Base64 spelling as an identity property.
- The input implementation calls Rust `unwrap()` on Base64 decoding. Malformed ASCII can therefore take a Rust panic path instead of a controlled PostgreSQL input error; fuzz malformed and oversized values on a disposable server before accepting untrusted input.
- Version `0.0.1` uses old Rust crates and a hand-written C/Rust FFI boundary. Test backend survival, dump/restore, architecture/endian behavior, upgrades, and the exact target PostgreSQL ABI.
- For ordinary binary data, prefer `bytea` and PostgreSQL's maintained Base64 functions; they preserve arbitrary length and have established error semantics.
