## Usage

Sources:

- [Official README](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/README.md)
- [Extension control file](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc.control)
- [Version 1.0 extension SQL](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc--1.0.sql)
- [C implementation](https://github.com/RekGRpth/pg_htmldoc/blob/74259576cabbdaa2234d95a300a3c128637f665e/pg_htmldoc.c)

`pg_htmldoc` 1.0 embeds HTMLDOC in a PostgreSQL backend. It accumulates HTML, server-side files, or URLs into a backend-local document and renders that document as PDF or PostScript; the extension does not expose HTMLDOC's EPUB or HTML output modes.

### Core Workflow

The least privileged input path is an in-memory HTML string followed by a `bytea` result:

```sql
CREATE EXTENSION pg_htmldoc;

SELECT htmldoc_addhtml('<h1>Quarterly report</h1><p>Complete.</p>');
SELECT octet_length(convert2pdf());
```

Multiple `htmldoc_add*` calls append sources to the same document before conversion. To retrieve the rendered document, call `convert2pdf()` or `convert2ps()` without arguments and consume the returned `bytea` in the client.

### Functions

- `htmldoc_addhtml(html text) RETURNS boolean` adds in-memory HTML.
- `htmldoc_addfile(file text) RETURNS boolean` reads a path visible to the database server OS user.
- `htmldoc_addurl(url text) RETURNS boolean` fetches a URL from the database server.
- `convert2pdf()` and `convert2ps()` return the rendered document as `bytea`.
- `convert2pdf(file text)` and `convert2ps(file text)` write to a server-side file and return `boolean`.

### Security and State Boundaries

The file overloads do not access the SQL client's filesystem. Input and output paths are opened by the PostgreSQL backend, while URL fetching uses the backend's network reachability. Revoke execution from untrusted users to prevent server-side file disclosure, arbitrary file creation within the service account's permissions, and server-side request forgery:

```sql
REVOKE EXECUTE ON FUNCTION htmldoc_addfile(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION htmldoc_addurl(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION convert2pdf(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION convert2ps(text) FROM PUBLIC;
```

The C implementation keeps the document in process-global state for the current PostgreSQL backend and does not explicitly clear it immediately after conversion. Reusing a pooled session can therefore carry or accumulate prior document state. Use a controlled, dedicated session for each conversion job and validate isolation before relying on repeated calls.

HTML, referenced resources, URLs, and output size are untrusted native-parser inputs that can consume backend memory, CPU, network, and disk. Apply statement timeouts, endpoint allowlists, input-size limits, and role restrictions outside the extension.
