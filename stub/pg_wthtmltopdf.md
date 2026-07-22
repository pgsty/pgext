## Usage

Sources:

- [Extension SQL for version 1.0](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf--1.0.sql)
- [PostgreSQL C wrapper](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf.c)
- [Wt PDF renderer implementation](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/MyWPdfRenderer.cc)
- [Extension control file](https://github.com/RekGRpth/pg_wthtmltopdf/blob/dc68969a4821945d7743cd4ab4b93258dd365251/pg_wthtmltopdf.control)

`pg_wthtmltopdf` exposes one server-side HTML-to-PDF function backed by Wt's `WPdfRenderer` and libharu. The function is small and rigid: it renders one A4 portrait page configuration and offers no SQL options for page size, orientation, margins, fonts, headers, or footers.

### Core Workflow

After the server library has been built against compatible Wt and libharu versions, create the extension and pass HTML text to `wthtmltopdf(text)`.

```sql
CREATE EXTENSION pg_wthtmltopdf;

SELECT wthtmltopdf(
  '<html><body><h1>Invoice</h1><p>Total: 42</p></body></html>'
);
```

The renderer sets A4 portrait output, compression, UTF encodings, a margin value of 1.0, and 96 DPI. Rendering happens synchronously inside the calling PostgreSQL backend.

### Important Object

- `wthtmltopdf(text) RETURNS text`: renders the supplied HTML and returns the generated PDF bytes.

### Binary-Type Warning

Although the SQL signature says `text`, the C function writes raw PDF bytes directly into a PostgreSQL text datum. PDF is binary data and may contain bytes that are invalid in the database encoding. Normal text clients, casts, string functions, logical decoding, dumps, or drivers can therefore reject, transform, or corrupt the result. This is not equivalent to a safe `bytea` API.

Before adoption, test the exact client in binary-result mode and verify the resulting file byte-for-byte. If the client cannot retrieve the raw payload safely, do not use this function; wrap or patch the extension to return `bytea`, or render outside PostgreSQL.

### Failure and Security Notes

NULL input raises an error. Malformed HTML or Wt/libharu failures can raise a PostgreSQL error, while some renderer failures return NULL. Large or adversarial HTML consumes CPU and memory in a database backend, so restrict EXECUTE privileges, cap statement duration, and avoid untrusted input. The extension provides no URL-fetch, stylesheet, font, or sandbox policy at the SQL surface; behavior depends on the linked Wt stack. Its upstream README is empty, so the reviewed SQL and source code are the authoritative interface for version `1.0`.
