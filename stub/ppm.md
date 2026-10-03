## Usage

Sources:

- [Control 1.0](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm.control)
- [SQL 1.0](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm--1.0.sql)
- [Type implementation](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/ppm_type.c)
- [Official test workflow](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/test_ppm.sql)
- [PostgreSQL 15 build recipe](https://github.com/TimeB30/Diploma/blob/ee6649138162c2a9db2e48b4f5f1a5651741daba/Implementation/Dockerfile)

`ppm` is an academic prototype that stores text in a PPM-compressed custom type. Its control file and install SQL define extension version 1.0; the startup log's 4.0 text is not the extension version. The repository includes a PostgreSQL 15 container build recipe but no explicit license or production support guarantee.

### Core Workflow

A superuser installs the extension, which provides the `ppm_text` type:

```sql
CREATE EXTENSION ppm;
SET ppm.max_order = 4;
CREATE TABLE compressed_notes (body ppm_text);
INSERT INTO compressed_notes VALUES ('Repeated text, repeated text.');
SELECT body, pg_column_size(body) FROM compressed_notes;
```

The type's input function compresses a C string, and its output function decompresses it for display. Its variable-length representation uses external storage, avoiding an additional built-in TOAST compression pass. It is a distinct type, not a server-wide compression replacement.

### Configuration and Objects

`ppm.max_order` sets the PPM context order for newly encoded values: default 2, range 0-8, session-settable. The encoder records the chosen order with the stored value, so decoding reads it from the value rather than the current setting.

`ppm_in` and `ppm_out` are the type's I/O functions. The install SQL does not define text operators, indexes or explicit text casts for this type; its SQL surface is limited to storage and I/O.

### Boundaries

The upstream workflow is a compression experiment comparing stored sizes and timings. Preserve original text and independently verify round trips before using the type for data that matters. No stable on-disk format, upgrade path, replication test matrix or production durability qualification is documented. Compatibility beyond the supplied PostgreSQL 15 recipe and redistribution permissions remain unverified.
