## Usage

Sources:

- [Official README](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/README.md)
- [Official extension SQL](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/pg_wkhtmltopdf--1.0.sql)
- [Official C implementation](https://github.com/RekGRpth/pg_wkhtmltopdf/blob/476e17cb2b4732f88b3c62b2c4de88948dff6a94/pg_wkhtmltopdf.c)

`pg_wkhtmltopdf` runs the libwkhtmltox/QtWebKit HTML-to-PDF engine inside a PostgreSQL backend and returns the PDF as `bytea`. It is useful only when rendering must happen in the database process; an external rendering service normally provides a safer resource and network boundary.

### Core Workflow

Set the input page, then render it in the same session:

```sql
CREATE EXTENSION pg_wkhtmltopdf;

SELECT wkhtmltopdf_set_object_setting(
  'page',
  'https://example.com/invoices/123'
);

SELECT wkhtmltopdf_convert();
```

The client must extract the returned `bytea` as raw bytes. The upstream `COPY ... FORMAT binary` example writes PostgreSQL's binary COPY envelope, not a plain PDF file; do not serve that file directly as a PDF.

### Functions and State

- `wkhtmltopdf_set_object_setting(text,text)` sets an object option such as the page URL.
- `wkhtmltopdf_set_global_setting(text,text)` sets a libwkhtmltox global option.
- `wkhtmltopdf_convert()` performs the conversion synchronously and returns `bytea`.

The settings objects are held in backend-static state, and the extension exposes no SQL reset function. Keep the set-and-convert sequence in one controlled session, avoid assuming transaction rollback restores settings, and test repeated calls carefully. Session-pool reuse can otherwise carry configuration into a later request.

### Security and Operations

Rendering performs network retrieval, HTML parsing, JavaScript execution, and PDF generation inside the database backend. A slow or hostile page can consume backend CPU/memory, block a worker, reach internal HTTP services, or exercise local-file and legacy browser-engine capabilities. Restrict `EXECUTE` on all three functions, allow-list destinations outside SQL, and enforce network/local-file controls at the operating-system boundary.

The extension offers no database-level timeout, cancellation policy, queue, or renderer isolation. Put strict statement timeouts and workload limits around calls, but do not treat them as equivalent to process isolation.

The build links to wkhtmltopdf's archived QtWebKit-based C library, so library ABI, security fixes, fonts, certificates, and rendering behavior are host dependencies. Version `1.0` has no documented PostgreSQL compatibility matrix. Build and run representative, adversarial pages on the exact server image before deployment.
