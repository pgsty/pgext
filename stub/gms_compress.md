## Usage

Sources:

- [Control 1.0](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/gms_compress.control)
- [SQL 1.0](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/gms_compress--1.0.sql)
- [Official examples](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gms_compress/sql/gms_compress.sql)

`gms_compress` 1.0 provides zlib-backed compression and decompression of `RAW` and `BLOB` values in openGauss. Its SQL uses kernel-specific types and `PACKAGE` syntax; it is not a stock PostgreSQL compatibility claim. Official openGauss builds include the module.

### Core Workflow

Install with a role permitted to create extensions in the chosen openGauss database:

```sql
CREATE EXTENSION gms_compress;
SELECT gms_compress.lz_compress('12ab56'::raw, 6);
SELECT gms_compress.lz_uncompress(
  gms_compress.lz_compress('12ab56'::raw, 6)
);
```

The extension creates the fixed `gms_compress` schema. Compression quality ranges from 1 through 9 and defaults to 6. Single-call overloads return a compressed or decompressed value; procedure overloads also expose output parameters.

### Streaming Interface

- `lz_compress_open` starts a context and returns an integer handle.
- `lz_compress_add` feeds a RAW chunk; `lz_compress_close` finishes the compressed BLOB and closes the context.
- `lz_uncompress_open` accepts a compressed BLOB; `lz_uncompress_extract` retrieves the decompressed RAW value; `lz_uncompress_close` closes the context.
- `isopen` checks whether a handle is active. These functions and procedures are in the extension's schema.

### Limits and Maintenance

At most five compression/decompression handles can be used, and each handle's data is limited to 1 GiB. Close contexts when finished and do not retain handles as persistent application identifiers. The streaming compressor's destination argument on its add procedure exists for syntax compatibility and does not carry state in the current implementation. This module handles RAW/BLOB values, not transparent table compression. Check the matching openGauss distribution when upgrading; no PostgreSQL-major support matrix is inferred from its ancestry.
