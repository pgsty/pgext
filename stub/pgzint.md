## Usage

Sources:

- [pgzint 0.2.0 README](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/README.md)
- [pgzint 0.2.0 release notes](https://github.com/davidbeauchamp/pgzint/releases/tag/v0.2.0)
- [pgzint 0.2.0 extension SQL](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/pgzint--0.2.0.sql)
- [pgzint 0.2.0 control file](https://github.com/davidbeauchamp/pgzint/blob/v0.2.0/pgzint.control)

`pgzint` generates barcode images inside PostgreSQL with the Zint library and returns PNG bytes as `bytea`. Use it when an application can consume binary images but should not integrate Zint directly.

### Core Workflow

Install the extension, inspect the symbol catalog, and call either a convenience wrapper or the general generator:

```sql
CREATE EXTENSION pgzint;

SELECT bc_symbol_zint_id, bc_symbol_zint_constant, bc_symbol_name
FROM bc_symbols
ORDER BY bc_symbol_zint_id;

SELECT bc_qrcode('SAMPLE');

SELECT bc_generate(
  'SAMPLE', 58, NULL, 2, 0, NULL, NULL,
  NULL, NULL, NULL, NULL, 14, NULL, 0
);
```

`bc_generate` accepts the payload, Zint symbology ID, height, scale, whitespace and border widths, output flags, colors, text flag, three symbology-specific options, and rotation. In 0.2.0 the height argument is `float8`.

### Important Objects

- `bc_generate` is the C-backed general generator.
- `bc_qrcode`, `bc_excode39`, `bc_pdf417`, `bc_maxicode`, and `bc_code128` are SQL wrappers with project-selected defaults.
- `bc_symbols` maps Zint numeric IDs to constants and display names.
- `getzintsymbolid(text)` and `getzintsymbolconstant(integer)` convert between those identifiers.
- `pgzint_version()` reports the installed pgzint version.

Version 0.2.0 removes the old `barcodes` view and simplifies `bc_symbols`; integrations must use the three columns shown above rather than the metadata columns removed by this release.

### Upgrade and Requirements

After installing the 0.2.0 package files, upgrade each database that already has pgzint:

```sql
ALTER EXTENSION pgzint UPDATE TO '0.2.0';
```

pgzint 0.2.0 requires PostgreSQL 9.4 or newer and Zint 2.14 or newer compiled with PNG support. It uses Zint's in-memory PNG output instead of the earlier BMP-to-PNG conversion and no longer has a direct libpng conversion layer.

Image generation consumes database CPU and can return large binary values. Bound payload sizes, avoid unreviewed bulk generation in latency-sensitive queries, and validate output against the scanners and barcode standards required by the application.
