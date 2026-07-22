## Usage

Sources:

- [Official README](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/README.md)
- [Extension SQL](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji--1.0.sql)
- [Extension control file](https://github.com/truthly/pg-emoji/blob/aa891c6f30c1e75ccc5a135959b692cdcbc24c79/emoji.control)

`emoji` is a pure-SQL codec that represents binary or UTF-8 text as a sequence of emoji. Use it for reversible visual encodings, not for encryption, authentication, or reliable error correction: its header contains only a 9-bit checksum.

### Core Workflow

Install the extension, encode a value, and decode it with the matching function:

```sql
CREATE EXTENSION emoji;

SELECT emoji.encode('\x0123456789abcdef'::bytea);
SELECT emoji.decode(emoji.encode('\x0123456789abcdef'::bytea));

SELECT emoji.from_text('Hello, PostgreSQL!');
SELECT emoji.to_text(emoji.from_text('Hello, PostgreSQL!'));
```

`emoji.from_text(text)` converts text to UTF-8 bytes before encoding, while `emoji.to_text(text)` decodes bytes as UTF-8. Invalid emoji input, a missing mapping, or a checksum mismatch makes `emoji.decode(text)` return `NULL`; the text wrapper consequently returns `NULL` too.

### API

- `emoji.encode(bytea) → text` encodes bytes into the fixed emoji alphabet.
- `emoji.decode(text) → bytea` validates the header and reconstructs the bytes.
- `emoji.from_text(text) → text` is the UTF-8 text encoder.
- `emoji.to_text(text) → text` is the UTF-8 text decoder.

The first output emoji stores a padding flag plus a 9-bit checksum. Upstream describes the resulting accidental-error detection probability as 511/512, about 99.8%; this is not a cryptographic integrity guarantee.

### Operational Notes

Version `1.0` installs a fixed `emoji` schema and a 1,024-row mapping based on the first 1,024 single-code-point entries from Emoji 13.1. Encoded data is therefore coupled to that table: do not reorder or edit `emoji.chars`, and retain the same extension data when decoding old values.

The SQL implementation calls `sha512(bytea)` but declares no extension dependency. Confirm that the target PostgreSQL build provides this function before installation. The reviewed upstream source contains only the `1.0` release from 2021, so test compatibility and Unicode/client rendering behavior on the exact server and client versions you deploy.
