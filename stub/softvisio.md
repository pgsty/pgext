## Usage

Sources:

- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/README.md)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/softvisio.control)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/softvisio--2.0.4.sql)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/Makefile)
- [Official source](https://github.com/softvisio/postgresql-softvisio/blob/v2.0.4/package.json)
- [Official source](https://www.postgresql.org/docs/18/functions-string.html)

softvisio 2.0.4 provides numeric domains, Unicode helpers and a large-object size function in SQL/PL/pgSQL. It needs no native library or preload. The project marks its package as proprietary. Its SQL uses the server's casefold function, available in PostgreSQL 18; no general supported-major matrix is published.

### Core workflow

```sql
CREATE EXTENSION softvisio;
SELECT 9007199254740991::int53;
SELECT 12.34::number2;
SELECT normalize_nfc('é'), normalize_nfc_casefold('Straße');
```

int53 restricts bigint values to the exact JavaScript integer range. number1 through number16 are numeric(16, scale) domains with additional exclusive magnitude bounds. Rounding to the declared scale precedes domain checks; these types do not promise arbitrary precision or lossless binary floating-point conversion.

### Boundaries

lo_size(oid) opens a large object read-only, seeks to its end and closes the descriptor; it requires the caller's normal large-object privileges and a transaction. The normalization helpers are strict, immutable and parallel safe. Unicode normalization requires UTF8, and case folding follows the input collation.

The control is relocatable and superuser=false, so installation still requires the normal database/schema creation rights. The installer reuses visible domains with matching names without verifying their definitions; use a controlled schema and search_path and inspect existing int53/number domains before installing. The supplied Makefile includes the reviewed upgrade chain through 2.0.4. Back up dependent objects and inspect the available extension update path before upgrading.
