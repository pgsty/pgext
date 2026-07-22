## Usage

Sources:

- [PostgreSQL binding at the reviewed commit](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven--0.5.0.sql)
- [PostgreSQL examples](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/bindings/sql/postgresql/proven_functions.sql)
- [Project documentation](https://github.com/hyperpolymath/proven/blob/302f2a473fb7baf522aecf59eaa60f4e7c4d966e/README.adoc)

`proven` exposes a small PostgreSQL C binding to libproven's Idris 2/Zig safety routines. Version 0.5.0 provides checked `BIGINT` arithmetic, simplified email/URL/IPv4/JSON validation, HTML-entity escaping, hexadecimal conversion, and a checksum helper.

```sql
CREATE EXTENSION proven;
SELECT proven_safe_add(9223372036854775807, 1);
SELECT proven_validate_email('user@example.com');
SELECT proven_hex_encode('hello');
```

Arithmetic functions return `NULL` on overflow, underflow, division by zero, or other errors rather than raising an exception. The validators establish only the upstream parser's simplified syntax; they do not prove address ownership, deliverability, URL safety, routability, or application-policy compliance.

Despite its name, `proven_hash_sha256(text)` returns a CRC32 checksum in hexadecimal, not SHA-256. Do not use it for passwords, signatures, integrity against an attacker, or content-addressed security; use a maintained cryptographic implementation such as `pgcrypto`. HTML escaping is context-specific and is not a universal XSS sanitizer.

Pin the native libproven ABI with the extension build and test invalid UTF-8, large inputs, `NULL`, overflow, and upgrade behavior. Formal verification of a core routine does not verify the PostgreSQL C wrapper, FFI, build chain, or surrounding application.
