## Usage

Sources:

- [README](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/README.md)
- [Control file](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/pg_htmldoc.control)
- [SQL](https://api.pgxn.org/src/pg_htmldoc/pg_htmldoc-1.0.11/pg_htmldoc--1.0.sql)

`pg_htmldoc` embeds HTMLDOC to turn queued HTML, files, or URLs into PDF or PostScript. Distribution 1.0.11 retains control version 1.0 but changes important permission and session-state behavior.

### Core Workflow

The following in-memory input example requires a superuser because embedded markup can reference server files or URLs. The no-argument conversion returns binary data to the client.

```sql
CREATE EXTENSION pg_htmldoc;
SELECT htmldoc_addhtml('<h1>Quarterly report</h1><p>Complete.</p>');
SELECT octet_length(convert2pdf());
```

### Objects and State

`htmldoc_addfile`, `htmldoc_addurl`, and `htmldoc_addhtml` append inputs to one backend-local document. `convert2pdf` and `convert2ps` return `bytea`, or write to a server-side path when passed a filename. Successful conversion clears the queued document; a later conversion needs new input. Null input and conversion without a queued document fail.

### Access Boundary

`pg_htmldoc.whitelist` is a superuser-controlled list of file and HTTP(S) prefixes. For ordinary roles it is the explicit access grant to matching input files or URLs; an empty list denies that access. For superusers a nonempty list narrows access. It does not permit ordinary roles to call `htmldoc_addhtml` or filename-output overloads: those always require a superuser. Returning a rendered `bytea` does not itself write a file.

The extension is relocatable and needs no preload. HTMLDOC must be available as a shared library. Bound native-parser inputs and external resource access, and keep pooled-session behavior in mind when composing multiple inputs. SQL function grants and the whitelist serve different purposes and both need appropriate configuration.
