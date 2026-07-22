## Usage

Sources:

- [Official README for version 0.1.4](https://github.com/davidbeauchamp/pgzint/blob/47cdb014cd81752c9d42b1d06c85871872a1793f/README.md)
- [Extension SQL for version 0.1.4](https://github.com/davidbeauchamp/pgzint/blob/47cdb014cd81752c9d42b1d06c85871872a1793f/pgzint--0.1.4.sql)
- [PGXN distribution page](https://pgxn.org/dist/pgzint/)

`pgzint` generates barcode images inside PostgreSQL through the Zint barcode library. Its SQL functions return PNG bytes as `bytea`, which is useful when an application can display binary images but should not integrate Zint itself.

### Core Workflow

Install the extension, choose a symbology, and return or store the generated binary result:

```sql
CREATE EXTENSION pgzint;

-- Convenience wrapper with project-provided QR defaults.
SELECT bc_qrcode('SAMPLE');

-- Inspect the installed symbol catalog before choosing an integer ID.
SELECT symbol_id, symbol_name
FROM barcodes
ORDER BY symbol_id;
```

The actual columns exposed by `barcodes` come from the installed 0.1.4 SQL, so use `\d+ barcodes` when integrating with a client. The primary API accepts the payload followed by the Zint symbol number and rendering options:

```sql
SELECT bc_generate(
  'SAMPLE', 58, NULL, 2, 0, NULL, NULL,
  NULL, NULL, NULL, NULL, 14, NULL, 0
);
```

The example uses the same QR symbology and defaults described by the upstream README. Rendering parameters are passed through to Zint; confirm their meaning against the Zint version installed with the extension.

### Important Objects

- `bc_generate` is the general-purpose C-backed generator. Its parameters control symbology, height, scale, whitespace, border, output flags, foreground/background colors, text display, three symbol-specific options, and rotation.
- `bc_qrcode`, `bc_excode39`, `bc_pdf417`, `bc_maxicode`, and `bc_code128` are SQL convenience wrappers with project-selected defaults.
- `bc_symbols` stores the extension's symbology metadata, while `barcodes` presents it as a view.
- `getzintsymbolconstant` and `getzintsymbolid` convert between numeric constants and text identifiers used by the generator.

`bc_generate` is the only generator implemented directly in C in this release; the convenience helpers delegate to it. Call it directly when wrapper defaults do not match the required barcode size or encoding options.

### Prerequisites and Caveats

Version 0.1.4 requires PostgreSQL 9.1 or later plus the Zint and PNG libraries at build and runtime. The extension control file does not declare preload or another PostgreSQL extension dependency. Existing databases upgraded from an earlier release must run the extension update after installing the 0.1.4 files:

```sql
ALTER EXTENSION pgzint UPDATE TO '0.1.4';
```

Image generation consumes database CPU and memory and returns potentially large binary values, so bound payload sizes and avoid unreviewed bulk generation in latency-sensitive queries. The project is old and the bundled symbology metadata reflects its Zint-era definitions; verify output with the scanners and standards required by the application.
