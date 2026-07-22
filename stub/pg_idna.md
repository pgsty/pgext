## Usage

Sources:

- [Official README](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/README.md)
- [Version 0.1.1 changelog](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/CHANGELOG.md)
- [SQL-visible Rust implementation](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/src/lib.rs)
- [Accepted conversion options](https://github.com/owenthewizard/pg_idna/blob/cbb25acf09ba478c5ca828d722519bf5292334e1/src/my_config.rs)

`pg_idna` provides PostgreSQL helpers for Unicode IDNA conversion using the Rust `idna` crate's UTS #46 implementation. It converts internationalized domain labels between Unicode and ASCII/Punycode and exposes lightweight classification helpers.

### Core Workflow

```sql
CREATE EXTENSION pg_idna;

SELECT idna_to_ascii('☕.us');
SELECT idna_to_unicode('xn--53h.us');
SELECT idna_is_ascii('example.org');
```

The default conversion policy uses ASCII deny list `url`, hyphen policy `allow`, and, for ASCII conversion, DNS-length policy `verify`.

### Functions and Parameters

- `idna_is_ascii(input text) RETURNS boolean` reports whether every byte is ASCII.
- `idna_is_punycode(input bytea) RETURNS boolean` is only a case-insensitive `xn--` prefix hint; it does not validate the rest of the label.
- `idna_to_ascii(input, adl, h, dl)` converts to ASCII. `adl` accepts `empty`, `std3`, or `url`; `h` accepts `allow`, `check_first_last`, or `check`; `dl` accepts `ignore`, `verify_allow_root_dot`, or `verify`.
- `idna_to_unicode(input, adl, h)` rejects failed conversions.
- `idna_to_unicode_lossy(input, adl, h)` returns output with U+FFFD replacement characters where invalid Punycode is encountered.

Optional arguments have defaults, so ordinary URL-oriented calls need only the input. Use explicit policies when validating DNS hostnames rather than general URL labels.

### Caveats

Upstream labels the project pre-alpha, says not to use it in production, and does not promise stable versioning. The catalog version `0.1.1-docs` corresponds to source based on the 0.1.1 release plus documentation changes.

Invalid option strings, constraint violations, malformed Punycode, or failed conversions cause the Rust functions to panic, which pgrx surfaces as a PostgreSQL error. `idna_is_punycode` must never be used as validation. IDNA conversion also does not establish that a hostname exists, is safe to visit, or is visually unambiguous; apply application-level allowlists and confusable-name policy separately.
