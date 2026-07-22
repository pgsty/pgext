## Usage

Sources:

- [Official README](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/README.md)
- [Official extension SQL](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf--1.0.sql)
- [Official C implementation](https://github.com/RekGRpth/pg_mupdf/blob/a15db6ca07c5e95b8f2d78fe92774c659a3f3c54/pg_mupdf.c)

`pg_mupdf` runs MuPDF inside a PostgreSQL backend to convert document text, especially HTML, into a binary output such as PDF. It offers a single synchronous conversion function; fetching remote HTML and saving the returned bytes are separate application concerns.

### Core Workflow

```sql
CREATE EXTENSION pg_mupdf;

WITH rendered AS (
  SELECT mupdf(
    '<html><body><h1>Quarterly report</h1></body></html>',
    input_type := 'html',
    output_type := 'pdf',
    options := 'compress',
    range := '1-N'
  ) AS pdf
)
SELECT octet_length(pdf), left(encode(pdf, 'hex'), 8) AS signature
FROM rendered;
```

Retrieve the resulting `bytea` through a binary-safe database driver and write those bytes as the output file. PostgreSQL `COPY ... FORMAT binary` adds the PostgreSQL binary-copy envelope and does not by itself create a raw PDF file.

### Function Reference

`mupdf(data text, input_type text DEFAULT 'html', output_type text DEFAULT 'pdf', options text DEFAULT '', range text DEFAULT '1-N')` opens the input through MuPDF's registered document handlers, renders the selected page range, and returns the writer output as `bytea`. The accepted type names, writer options, and page-range grammar come from the linked MuPDF build, so validate them against the packaged library.

### Operational Boundaries

Conversion runs in the calling backend and the implementation creates an unlimited MuPDF store. Apply strict input-size, page-count, output-size, timeout, and concurrency limits outside the function; otherwise complex or malicious content can exhaust a database worker's CPU or memory. Keep MuPDF patched, because the database process parses attacker-controlled document formats in-process. Pass all arguments as non-NULL, test error cleanup and the actual linked MuPDF version, and do not grant the function broadly until document trust, HTML resource behavior, fonts, output fidelity, and failure isolation have been reviewed.
